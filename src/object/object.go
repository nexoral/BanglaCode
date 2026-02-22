package object

import (
	"BanglaCode/src/ast"
	"bytes"
	"fmt"
	"strings"
	"sync"
)

// ObjectType represents the type of an object
type ObjectType string

const (
	NUMBER_OBJ          = "NUMBER"
	STRING_OBJ          = "STRING"
	BOOLEAN_OBJ         = "BOOLEAN"
	NULL_OBJ            = "NULL"
	RETURN_OBJ          = "RETURN"
	ERROR_OBJ           = "ERROR"
	TYPE_ERROR_OBJ      = "TYPE_ERROR"
	REFERENCE_ERROR_OBJ = "REFERENCE_ERROR"
	RANGE_ERROR_OBJ     = "RANGE_ERROR"
	SYNTAX_ERROR_OBJ    = "SYNTAX_ERROR"
	FUNCTION_OBJ        = "FUNCTION"
	BUILTIN_OBJ         = "BUILTIN"
	ARRAY_OBJ           = "ARRAY"
	MAP_OBJ             = "MAP"
	CLASS_OBJ           = "CLASS"
	INSTANCE_OBJ        = "INSTANCE"
	BREAK_OBJ           = "BREAK"
	CONTINUE_OBJ        = "CONTINUE"
	EXCEPTION_OBJ       = "EXCEPTION"
	MODULE_OBJ          = "MODULE"
	PROMISE_OBJ         = "PROMISE"
	DB_CONNECTION_OBJ   = "DB_CONNECTION"
	DB_RESULT_OBJ       = "DB_RESULT"
	DB_POOL_OBJ         = "DB_POOL"
	EVENT_EMITTER_OBJ   = "EVENT_EMITTER"
	BUFFER_OBJ          = "BUFFER"
	WORKER_OBJ          = "WORKER"
	STREAM_OBJ          = "STREAM"
	URL_OBJ             = "URL"
	URL_PARAMS_OBJ      = "URL_PARAMS"
	SET_OBJ             = "SET"
	ES6MAP_OBJ          = "ES6MAP"
	GENERATOR_OBJ       = "GENERATOR"
)

// Object represents any runtime value
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Number represents a numeric value (int or float)
type Number struct {
	Value float64
}

func (n *Number) Type() ObjectType { return NUMBER_OBJ }
func (n *Number) Inspect() string  { return fmt.Sprintf("%g", n.Value) }

// String represents a string value
type String struct {
	Value string
}

func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) Inspect() string  { return s.Value }

// Boolean represents a boolean value
type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

// Null represents a null value (khali)
type Null struct{}

func (n *Null) Type() ObjectType { return NULL_OBJ }
func (n *Null) Inspect() string  { return "khali" }

// ReturnValue wraps a value for return statements
type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Type() ObjectType { return RETURN_OBJ }
func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }

// StackFrame represents a single frame in the stack trace
type StackFrame struct {
	Function string
	File     string
	Line     int
	Column   int
}

// Error represents a runtime error
type Error struct {
	Message   string
	ErrorType ObjectType // ERROR_OBJ, TYPE_ERROR_OBJ, REFERENCE_ERROR_OBJ, etc.
	Line      int
	Column    int
	Stack     []StackFrame
}

func (e *Error) Type() ObjectType {
	if e.ErrorType != "" {
		return e.ErrorType
	}
	return ERROR_OBJ
}

func (e *Error) Inspect() string {
	var errorTypeName string
	switch e.ErrorType {
	case TYPE_ERROR_OBJ:
		errorTypeName = "TypeError"
	case REFERENCE_ERROR_OBJ:
		errorTypeName = "ReferenceError"
	case RANGE_ERROR_OBJ:
		errorTypeName = "RangeError"
	case SYNTAX_ERROR_OBJ:
		errorTypeName = "SyntaxError"
	default:
		errorTypeName = "Error"
	}

	if e.Line > 0 {
		return fmt.Sprintf("%s [line %d, col %d]: %s", errorTypeName, e.Line, e.Column, e.Message)
	}
	return errorTypeName + ": " + e.Message
}

// GetStack returns formatted stack trace
func (e *Error) GetStack() string {
	if len(e.Stack) == 0 {
		return e.Inspect()
	}

	var buf bytes.Buffer
	buf.WriteString(e.Inspect())
	buf.WriteString("\nStack trace:\n")

	for i, frame := range e.Stack {
		if frame.Function != "" {
			buf.WriteString(fmt.Sprintf("  at %s", frame.Function))
		} else {
			buf.WriteString("  at <anonymous>")
		}

		if frame.File != "" {
			buf.WriteString(fmt.Sprintf(" (%s:%d:%d)", frame.File, frame.Line, frame.Column))
		} else if frame.Line > 0 {
			buf.WriteString(fmt.Sprintf(" (line %d, col %d)", frame.Line, frame.Column))
		}

		if i < len(e.Stack)-1 {
			buf.WriteString("\n")
		}
	}

	return buf.String()
}

// NewError creates a generic error
func NewError(message string) *Error {
	return &Error{
		Message:   message,
		ErrorType: ERROR_OBJ,
	}
}

// NewTypeError creates a TypeError
func NewTypeError(message string) *Error {
	return &Error{
		Message:   message,
		ErrorType: TYPE_ERROR_OBJ,
	}
}

// NewReferenceError creates a ReferenceError
func NewReferenceError(message string) *Error {
	return &Error{
		Message:   message,
		ErrorType: REFERENCE_ERROR_OBJ,
	}
}

// NewRangeError creates a RangeError
func NewRangeError(message string) *Error {
	return &Error{
		Message:   message,
		ErrorType: RANGE_ERROR_OBJ,
	}
}

// NewSyntaxError creates a SyntaxError
func NewSyntaxError(message string) *Error {
	return &Error{
		Message:   message,
		ErrorType: SYNTAX_ERROR_OBJ,
	}
}

// AddStackFrame adds a frame to the error's stack trace
func (e *Error) AddStackFrame(function, file string, line, column int) {
	e.Stack = append(e.Stack, StackFrame{
		Function: function,
		File:     file,
		Line:     line,
		Column:   column,
	})
}

// Function represents a user-defined function
type Function struct {
	Parameters    []*ast.Identifier
	RestParameter *ast.Identifier // optional rest parameter (...args)
	Body          *ast.BlockStatement
	Env           *Environment
	Name          string
	IsAsync       bool // true for async functions (proyash kaj)
	IsGenerator   bool // true for generator functions (kaj*)
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	if f.RestParameter != nil {
		params = append(params, "..."+f.RestParameter.String())
	}
	if f.IsGenerator {
		out.WriteString("kaj*")
	} else {
		out.WriteString("kaj")
	}
	if f.Name != "" {
		out.WriteString(" " + f.Name)
	}
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

// BuiltinFunction represents a built-in function
type BuiltinFunction func(args ...Object) Object

// Builtin wraps a built-in function
type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }
func (b *Builtin) Inspect() string  { return "builtin function" }

// Array represents an array/list
type Array struct {
	Elements []Object
}

func (a *Array) Type() ObjectType { return ARRAY_OBJ }
func (a *Array) Inspect() string {
	var out bytes.Buffer
	elements := []string{}
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// Map represents a hash map/object
type Map struct {
	Pairs map[string]Object
}

func (m *Map) Type() ObjectType { return MAP_OBJ }
func (m *Map) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for key, value := range m.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", key, value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

// Class represents a class definition
type Class struct {
	Name             string
	Methods          map[string]*Function
	Getters          map[string]*Function // getter methods
	Setters          map[string]*Function // setter methods
	StaticProperties map[string]Object    // static properties
}

func (c *Class) Type() ObjectType { return CLASS_OBJ }
func (c *Class) Inspect() string  { return "sreni " + c.Name }

// Instance represents an instance of a class
type Instance struct {
	Class         *Class
	Properties    map[string]Object
	PrivateFields map[string]Object // private fields (underscore prefix)
}

func (i *Instance) Type() ObjectType { return INSTANCE_OBJ }
func (i *Instance) Inspect() string  { return i.Class.Name + " er udahoron" }

// Break represents a break statement
type Break struct{}

func (b *Break) Type() ObjectType { return BREAK_OBJ }
func (b *Break) Inspect() string  { return "thamo" }

// Continue represents a continue statement
type Continue struct{}

func (c *Continue) Type() ObjectType { return CONTINUE_OBJ }
func (c *Continue) Inspect() string  { return "chharo" }

// Exception represents a thrown exception
type Exception struct {
	Message string
	Value   Object
}

func (e *Exception) Type() ObjectType { return EXCEPTION_OBJ }
func (e *Exception) Inspect() string {
	if e.Value != nil {
		return "Exception: " + e.Value.Inspect()
	}
	return "Exception: " + e.Message
}

// Module represents an imported module
type Module struct {
	Name    string
	Exports map[string]Object
}

func (m *Module) Type() ObjectType { return MODULE_OBJ }
func (m *Module) Inspect() string  { return "module " + m.Name }

// PromiseState represents the state of a promise
type PromiseState string

const (
	PROMISE_PENDING  PromiseState = "PENDING"
	PROMISE_RESOLVED PromiseState = "RESOLVED"
	PROMISE_REJECTED PromiseState = "REJECTED"
)

// Promise represents an async operation that will complete in the future
type Promise struct {
	State      PromiseState
	Value      Object      // resolved value
	Error      Object      // rejection error
	ResultChan chan Object // for goroutine communication
	ErrorChan  chan Object // for error communication
	Mu         sync.RWMutex
}

func (p *Promise) Type() ObjectType { return PROMISE_OBJ }
func (p *Promise) Inspect() string {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	switch p.State {
	case PROMISE_RESOLVED:
		if p.Value != nil {
			return fmt.Sprintf("Promise(RESOLVED: %s)", p.Value.Inspect())
		}
		return "Promise(RESOLVED)"
	case PROMISE_REJECTED:
		if p.Error != nil {
			return fmt.Sprintf("Promise(REJECTED: %s)", p.Error.Inspect())
		}
		return "Promise(REJECTED)"
	default:
		return "Promise(PENDING)"
	}
}

// CreatePromise creates a new pending promise with channels
func CreatePromise() *Promise {
	return &Promise{
		State:      PROMISE_PENDING,
		ResultChan: make(chan Object, 1),
		ErrorChan:  make(chan Object, 1),
	}
}

// ResolvePromise resolves a promise with a value
func ResolvePromise(promise *Promise, value Object) {
	promise.Mu.Lock()
	promise.State = PROMISE_RESOLVED
	promise.Value = value
	promise.Mu.Unlock()
	promise.ResultChan <- value
}

// RejectPromise rejects a promise with an error
func RejectPromise(promise *Promise, err Object) {
	promise.Mu.Lock()
	promise.State = PROMISE_REJECTED
	promise.Error = err
	promise.Mu.Unlock()
	promise.ErrorChan <- err
}

// DBConnection represents a database connection
type DBConnection struct {
	ID       string            // Unique connection identifier
	DBType   string            // Database type: "postgres", "mysql", "mongodb", "redis"
	Native   interface{}       // Underlying driver connection (*sql.DB, *mongo.Client, *redis.Client)
	PoolID   string            // Reference to connection pool (if pooled)
	Metadata map[string]Object // Connection metadata (host, port, database, etc.)
	Mu       sync.RWMutex      // Thread-safe access
}

func (d *DBConnection) Type() ObjectType { return DB_CONNECTION_OBJ }
func (d *DBConnection) Inspect() string {
	d.Mu.RLock()
	defer d.Mu.RUnlock()
	return fmt.Sprintf("DB_CONNECTION(%s, id=%s)", d.DBType, d.ID)
}

// DBResult represents a database query result
type DBResult struct {
	Rows         []map[string]Object // Result rows as maps (column name -> value)
	RowsAffected int64               // Rows affected by INSERT/UPDATE/DELETE
	LastInsertID int64               // Last inserted ID (for SQL databases)
	Error        *Error              // Query error (if any)
}

func (d *DBResult) Type() ObjectType { return DB_RESULT_OBJ }
func (d *DBResult) Inspect() string {
	if d.Error != nil {
		return fmt.Sprintf("DB_RESULT(error: %s)", d.Error.Message)
	}
	return fmt.Sprintf("DB_RESULT(rows=%d, affected=%d)", len(d.Rows), d.RowsAffected)
}

// DBPool represents a connection pool
type DBPool struct {
	ID          string            // Unique pool identifier
	DBType      string            // Database type
	MaxConns    int               // Maximum connections
	ActiveConns int               // Currently active connections
	Config      map[string]Object // Pool configuration
	Conns       []*DBConnection   // Pool of reusable connections
	Mu          sync.RWMutex      // Thread-safe pool management
}

func (d *DBPool) Type() ObjectType { return DB_POOL_OBJ }
func (d *DBPool) Inspect() string {
	d.Mu.RLock()
	defer d.Mu.RUnlock()
	return fmt.Sprintf("DB_POOL(%s, active=%d/%d)", d.DBType, d.ActiveConns, d.MaxConns)
}

// EventListener represents a single event listener
type EventListener struct {
	Callback Object // Function to call
	Once     bool   // If true, remove after first call
}

// EventEmitter represents an event emitter for event-driven architecture
type EventEmitter struct {
	Events map[string][]*EventListener // Event name -> list of listeners
	Mu     sync.RWMutex                // Thread-safe access
}

func (e *EventEmitter) Type() ObjectType { return EVENT_EMITTER_OBJ }
func (e *EventEmitter) Inspect() string {
	e.Mu.RLock()
	defer e.Mu.RUnlock()
	eventCount := len(e.Events)
	listenerCount := 0
	for _, listeners := range e.Events {
		listenerCount += len(listeners)
	}
	return fmt.Sprintf("EventEmitter(events=%d, listeners=%d)", eventCount, listenerCount)
}

// CreateEventEmitter creates a new EventEmitter
func CreateEventEmitter() *EventEmitter {
	return &EventEmitter{
		Events: make(map[string][]*EventListener),
	}
}

// Buffer represents a binary data buffer
type Buffer struct {
	Data []byte       // Raw binary data
	Mu   sync.RWMutex // Thread-safe access
}

func (b *Buffer) Type() ObjectType { return BUFFER_OBJ }
func (b *Buffer) Inspect() string {
	b.Mu.RLock()
	defer b.Mu.RUnlock()
	return fmt.Sprintf("Buffer(length=%d)", len(b.Data))
}

// CreateBuffer creates a new Buffer with the given size
func CreateBuffer(size int) *Buffer {
	return &Buffer{
		Data: make([]byte, size),
	}
}

// CreateBufferFrom creates a new Buffer from existing data
func CreateBufferFrom(data []byte) *Buffer {
	// Make a copy to avoid external modifications
	bufData := make([]byte, len(data))
	copy(bufData, data)
	return &Buffer{
		Data: bufData,
	}
}

// Worker represents a worker thread
type Worker struct {
	ID           int           // Unique worker ID
	MessageChan  chan Object   // Channel for sending messages to worker
	ResponseChan chan Object   // Channel for receiving messages from worker
	StopChan     chan struct{} // Channel to signal worker termination
	OnMessage    *Function     // Message handler function
	IsRunning    bool          // Worker running state
	WorkerData   Object        // Initial data passed to worker
	Mu           sync.RWMutex  // Thread-safe access
}

func (w *Worker) Type() ObjectType { return WORKER_OBJ }
func (w *Worker) Inspect() string {
	w.Mu.RLock()
	defer w.Mu.RUnlock()
	status := "terminated"
	if w.IsRunning {
		status = "running"
	}
	return fmt.Sprintf("Worker(id=%d, status=%s)", w.ID, status)
}

// Stream represents a data stream (Readable, Writable, or Transform)
type Stream struct {
	StreamType    string       // "readable", "writable", "transform"
	Buffer        []byte       // Internal buffer
	IsClosed      bool         // Stream closed state
	IsEnded       bool         // Stream ended state (readable)
	HighWaterMark int          // Buffer size threshold
	OnData        *Function    // Data event handler
	OnEnd         *Function    // End event handler
	OnError       *Function    // Error event handler
	Mu            sync.RWMutex // Thread-safe access
}

func (s *Stream) Type() ObjectType { return STREAM_OBJ }
func (s *Stream) Inspect() string {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	status := "open"
	if s.IsClosed {
		status = "closed"
	} else if s.IsEnded {
		status = "ended"
	}
	return fmt.Sprintf("Stream(type=%s, status=%s, buffered=%d)", s.StreamType, status, len(s.Buffer))
}

// URL represents a parsed URL with all components
type URL struct {
	Href     string // Full URL
	Protocol string // e.g., "http:", "https:"
	Username string // Username in URL
	Password string // Password in URL
	Hostname string // Hostname without port
	Port     string // Port number as string
	Host     string // Hostname:port
	Pathname string // Path component
	Search   string // Query string including "?"
	Hash     string // Fragment including "#"
	Origin   string // Protocol + hostname + port
}

func (u *URL) Type() ObjectType { return URL_OBJ }
func (u *URL) Inspect() string {
	return fmt.Sprintf("URL(%s)", u.Href)
}

// URLSearchParams represents URL query parameters
type URLSearchParams struct {
	Params map[string][]string // Key to multiple values
}

func (u *URLSearchParams) Type() ObjectType { return URL_PARAMS_OBJ }
func (u *URLSearchParams) Inspect() string {
	var pairs []string
	for key, values := range u.Params {
		for _, val := range values {
			pairs = append(pairs, fmt.Sprintf("%s=%s", key, val))
		}
	}
	return fmt.Sprintf("URLSearchParams(%s)", strings.Join(pairs, "&"))
}

// Set is a collection of unique values
type Set struct {
	Elements map[string]bool // Using string hash of values for uniqueness
	Order    []Object        // Maintain insertion order
}

func (s *Set) Type() ObjectType { return SET_OBJ }
func (s *Set) Inspect() string {
	var out bytes.Buffer
	elements := []string{}
	for _, elem := range s.Order {
		elements = append(elements, elem.Inspect())
	}
	out.WriteString("Set(")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString(")")
	return out.String()
}

// ES6Map is a Map that supports any type as key (not just strings)
type ES6Map struct {
	Pairs map[string]Object // key hash -> value
	Keys  map[string]Object // key hash -> original key
	Order []string          // Maintain insertion order (key hashes)
}

func (m *ES6Map) Type() ObjectType { return ES6MAP_OBJ }
func (m *ES6Map) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for _, keyHash := range m.Order {
		key := m.Keys[keyHash]
		value := m.Pairs[keyHash]
		pairs = append(pairs, fmt.Sprintf("%s => %s", key.Inspect(), value.Inspect()))
	}
	out.WriteString("Map(")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString(")")
	return out.String()
}

// Generator represents a generator object
type Generator struct {
	Function *Function    // The generator function
	Env      *Environment // Execution environment
	State    string       // "suspended", "executing", "completed"
	Value    Object       // Last yielded/returned value
	Index    int          // Current execution position (statement index)
	Done     bool         // Whether generator is exhausted
}

func (g *Generator) Type() ObjectType { return GENERATOR_OBJ }
func (g *Generator) Inspect() string {
	return fmt.Sprintf("Generator(state=%s, done=%t)", g.State, g.Done)
}

// Singleton objects for common values
var (
	NULL     = &Null{}
	TRUE     = &Boolean{Value: true}
	FALSE    = &Boolean{Value: false}
	BREAK    = &Break{}
	CONTINUE = &Continue{}
)

// NativeBoolToBooleanObject converts a native bool to a Boolean object
func NativeBoolToBooleanObject(input bool) *Boolean {
	if input {
		return TRUE
	}
	return FALSE
}
