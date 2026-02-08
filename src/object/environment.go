package object

// Environment represents a scope for variable bindings
type Environment struct {
	store map[string]Object
	outer *Environment // parent scope
}

// NewEnvironment creates a new environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

// NewEnclosedEnvironment creates a new environment with an outer scope
func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

// Get retrieves a variable from the environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

// Set assigns a variable in the environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}

// Update updates a variable in the environment (searches outer scopes)
func (e *Environment) Update(name string, val Object) Object {
	_, ok := e.store[name]
	if ok {
		e.store[name] = val
		return val
	}
	if e.outer != nil {
		return e.outer.Update(name, val)
	}
	// If variable doesn't exist, create it in current scope
	e.store[name] = val
	return val
}

// All returns all variables in the current scope (not including outer scopes)
func (e *Environment) All() map[string]Object {
	return e.store
}
