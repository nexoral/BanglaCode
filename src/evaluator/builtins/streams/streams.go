package streams

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/object"
	"fmt"
)

var (
	evalFunc func(node ast.Node, env *object.Environment) object.Object
)

// SetEvalFunc sets the evaluation function for executing stream callbacks
func SetEvalFunc(fn func(ast.Node, *object.Environment) object.Object) {
	evalFunc = fn
}

// Builtins contains all stream-related built-in functions
var Builtins = map[string]*object.Builtin{
	"stream_readable_srishti": {
		Fn: streamReadableSrishti,
	},
	"stream_writable_srishti": {
		Fn: streamWritableSrishti,
	},
	"stream_poro": {
		Fn: streamPoro,
	},
	"stream_lekho": {
		Fn: streamLekho,
	},
	"stream_bondho": {
		Fn: streamBondho,
	},
	"stream_shesh": {
		Fn: streamShesh,
	},
	"stream_pipe": {
		Fn: streamPipe,
	},
	"stream_on": {
		Fn: streamOn,
	},
}

// streamReadableSrishti creates a new readable stream
// Usage: dhoro stream = stream_readable_srishti();
func streamReadableSrishti(args ...object.Object) object.Object {
	// Optional: high water mark (buffer size threshold)
	highWaterMark := 16384 // Default 16KB
	if len(args) > 0 {
		if num, ok := args[0].(*object.Number); ok {
			highWaterMark = int(num.Value)
		}
	}

	stream := &object.Stream{
		StreamType:    "readable",
		Buffer:        make([]byte, 0),
		IsClosed:      false,
		IsEnded:       false,
		HighWaterMark: highWaterMark,
	}

	return stream
}

// streamWritableSrishti creates a new writable stream
// Usage: dhoro stream = stream_writable_srishti();
func streamWritableSrishti(args ...object.Object) object.Object {
	// Optional: high water mark
	highWaterMark := 16384
	if len(args) > 0 {
		if num, ok := args[0].(*object.Number); ok {
			highWaterMark = int(num.Value)
		}
	}

	stream := &object.Stream{
		StreamType:    "writable",
		Buffer:        make([]byte, 0),
		IsClosed:      false,
		HighWaterMark: highWaterMark,
	}

	return stream
}

// streamPoro reads data from a readable stream
// Usage: dhoro data = stream_poro(stream, size?);
func streamPoro(args ...object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "stream_poro() requires at least 1 argument (stream)"}
	}

	stream, ok := args[0].(*object.Stream)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("stream_poro() first argument must be a Stream, got %s", args[0].Type())}
	}

	stream.Mu.Lock()
	defer stream.Mu.Unlock()

	if stream.StreamType != "readable" {
		return &object.Error{Message: "stream_poro() can only read from readable streams"}
	}

	if stream.IsClosed {
		return object.NULL
	}

	// Optional: read size
	readSize := len(stream.Buffer) // Read all by default
	if len(args) > 1 {
		if num, ok := args[1].(*object.Number); ok {
			readSize = int(num.Value)
			if readSize > len(stream.Buffer) {
				readSize = len(stream.Buffer)
			}
		}
	}

	if readSize == 0 {
		return object.NULL
	}

	// Read data from buffer
	data := make([]byte, readSize)
	copy(data, stream.Buffer[:readSize])

	// Remove read data from buffer
	stream.Buffer = stream.Buffer[readSize:]

	// Check if stream ended and buffer is empty
	if stream.IsEnded && len(stream.Buffer) == 0 {
		stream.IsClosed = true
	}

	return &object.String{Value: string(data)}
}

// streamLekho writes data to a writable stream
// Usage: stream_lekho(stream, data);
func streamLekho(args ...object.Object) object.Object {
	if len(args) < 2 {
		return &object.Error{Message: "stream_lekho() requires 2 arguments (stream, data)"}
	}

	stream, ok := args[0].(*object.Stream)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("stream_lekho() first argument must be a Stream, got %s", args[0].Type())}
	}

	stream.Mu.Lock()
	defer stream.Mu.Unlock()

	if stream.StreamType != "writable" {
		return &object.Error{Message: "stream_lekho() can only write to writable streams"}
	}

	if stream.IsClosed {
		return &object.Error{Message: "Cannot write to closed stream"}
	}

	// Convert data to bytes
	var data []byte
	switch arg := args[1].(type) {
	case *object.String:
		data = []byte(arg.Value)
	case *object.Buffer:
		arg.Mu.RLock()
		data = make([]byte, len(arg.Data))
		copy(data, arg.Data)
		arg.Mu.RUnlock()
	default:
		data = []byte(args[1].Inspect())
	}

	// Append to buffer
	stream.Buffer = append(stream.Buffer, data...)

	// Trigger data event if handler exists
	if stream.OnData != nil && evalFunc != nil {
		env := object.NewEnvironment()

		callExpr := &ast.CallExpression{
			Function:  &ast.Identifier{Value: "dataHandler"},
			Arguments: []ast.Expression{&ast.Identifier{Value: "chunk"}},
		}

		env.Set("dataHandler", stream.OnData)
		env.Set("chunk", &object.String{Value: string(data)})

		evalFunc(callExpr, env)
	}

	// Return true if below high water mark, false if backpressure needed
	return object.NativeBoolToBooleanObject(len(stream.Buffer) < stream.HighWaterMark)
}

// streamBondho closes a stream
// Usage: stream_bondho(stream);
func streamBondho(args ...object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "stream_bondho() requires 1 argument (stream)"}
	}

	stream, ok := args[0].(*object.Stream)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("stream_bondho() argument must be a Stream, got %s", args[0].Type())}
	}

	stream.Mu.Lock()
	stream.IsClosed = true
	stream.Mu.Unlock()

	return object.NULL
}

// streamShesh signals end of readable stream
// Usage: stream_shesh(stream);
func streamShesh(args ...object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "stream_shesh() requires 1 argument (stream)"}
	}

	stream, ok := args[0].(*object.Stream)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("stream_shesh() argument must be a Stream, got %s", args[0].Type())}
	}

	stream.Mu.Lock()
	stream.IsEnded = true

	// If buffer is empty, close immediately
	if len(stream.Buffer) == 0 {
		stream.IsClosed = true
	}
	stream.Mu.Unlock()

	// Trigger end event if handler exists
	if stream.OnEnd != nil && evalFunc != nil {
		env := object.NewEnvironment()

		callExpr := &ast.CallExpression{
			Function:  &ast.Identifier{Value: "endHandler"},
			Arguments: []ast.Expression{},
		}

		env.Set("endHandler", stream.OnEnd)

		evalFunc(callExpr, env)
	}

	return object.NULL
}

// streamPipe pipes data from readable to writable stream
// Usage: stream_pipe(readable, writable);
func streamPipe(args ...object.Object) object.Object {
	if len(args) < 2 {
		return &object.Error{Message: "stream_pipe() requires 2 arguments (readable, writable)"}
	}

	readable, ok := args[0].(*object.Stream)
	if !ok || readable.StreamType != "readable" {
		return &object.Error{Message: "stream_pipe() first argument must be a readable Stream"}
	}

	writable, ok := args[1].(*object.Stream)
	if !ok || writable.StreamType != "writable" {
		return &object.Error{Message: "stream_pipe() second argument must be a writable Stream"}
	}

	// Transfer all data from readable to writable
	readable.Mu.Lock()
	data := make([]byte, len(readable.Buffer))
	copy(data, readable.Buffer)
	readable.Buffer = readable.Buffer[:0] // Clear buffer
	readable.Mu.Unlock()

	writable.Mu.Lock()
	writable.Buffer = append(writable.Buffer, data...)
	writable.Mu.Unlock()

	return writable
}

// streamOn registers event handlers for streams
// Usage: stream_on(stream, "data", kaj(chunk) { ... });
func streamOn(args ...object.Object) object.Object {
	if len(args) < 3 {
		return &object.Error{Message: "stream_on() requires 3 arguments (stream, event, callback)"}
	}

	stream, ok := args[0].(*object.Stream)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("stream_on() first argument must be a Stream, got %s", args[0].Type())}
	}

	eventName, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("stream_on() second argument must be a string, got %s", args[1].Type())}
	}

	callback, ok := args[2].(*object.Function)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("stream_on() third argument must be a function, got %s", args[2].Type())}
	}

	stream.Mu.Lock()
	defer stream.Mu.Unlock()

	switch eventName.Value {
	case "data":
		stream.OnData = callback
	case "end":
		stream.OnEnd = callback
	case "error":
		stream.OnError = callback
	default:
		return &object.Error{Message: fmt.Sprintf("Unknown stream event: %s", eventName.Value)}
	}

	return object.NULL
}
