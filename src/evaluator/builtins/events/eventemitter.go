package events

import (
	"BanglaCode/src/object"
	"fmt"
)

// Builtins exports all event-related built-in functions
var Builtins = map[string]*object.Builtin{
	"ghotona_srishti":     {Fn: createEventEmitter},
	"ghotona_shuno":       {Fn: addEventListener},
	"ghotona_ekbar":       {Fn: addEventListenerOnce},
	"ghotona_prokash":     {Fn: emitEvent},
	"ghotona_bondho":      {Fn: removeEventListener},
	"ghotona_sob_bondho":  {Fn: removeAllListeners},
	"ghotona_shrotara":    {Fn: getListeners},
	"ghotona_naam_sob":    {Fn: getEventNames},
}

// createEventEmitter creates a new EventEmitter
// Usage: dhoro emitter = ghotona_srishti();
func createEventEmitter(args ...object.Object) object.Object {
	if len(args) != 0 {
		return &object.Error{Message: "ghotona_srishti() expects 0 arguments"}
	}

	return object.CreateEventEmitter()
}

// addEventListener adds an event listener
// Usage: ghotona_shuno(emitter, "event_name", callback);
func addEventListener(args ...object.Object) object.Object {
	if len(args) != 3 {
		return &object.Error{Message: "ghotona_shuno() expects 3 arguments (emitter, event_name, callback)"}
	}

	// Get emitter
	emitter, ok := args[0].(*object.EventEmitter)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("first argument must be EventEmitter, got %s", args[0].Type())}
	}

	// Get event name
	eventName, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("second argument must be string, got %s", args[1].Type())}
	}

	// Get callback function
	callback := args[2]
	if callback.Type() != object.FUNCTION_OBJ && callback.Type() != object.BUILTIN_OBJ {
		return &object.Error{Message: fmt.Sprintf("third argument must be function, got %s", callback.Type())}
	}

	// Add listener (thread-safe)
	emitter.Mu.Lock()
	defer emitter.Mu.Unlock()

	listener := &object.EventListener{
		Callback: callback,
		Once:     false,
	}

	emitter.Events[eventName.Value] = append(emitter.Events[eventName.Value], listener)

	return emitter
}

// addEventListenerOnce adds an event listener that runs only once
// Usage: ghotona_ekbar(emitter, "event_name", callback);
func addEventListenerOnce(args ...object.Object) object.Object {
	if len(args) != 3 {
		return &object.Error{Message: "ghotona_ekbar() expects 3 arguments (emitter, event_name, callback)"}
	}

	// Get emitter
	emitter, ok := args[0].(*object.EventEmitter)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("first argument must be EventEmitter, got %s", args[0].Type())}
	}

	// Get event name
	eventName, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("second argument must be string, got %s", args[1].Type())}
	}

	// Get callback function
	callback := args[2]
	if callback.Type() != object.FUNCTION_OBJ && callback.Type() != object.BUILTIN_OBJ {
		return &object.Error{Message: fmt.Sprintf("third argument must be function, got %s", callback.Type())}
	}

	// Add listener (thread-safe)
	emitter.Mu.Lock()
	defer emitter.Mu.Unlock()

	listener := &object.EventListener{
		Callback: callback,
		Once:     true, // Mark as once
	}

	emitter.Events[eventName.Value] = append(emitter.Events[eventName.Value], listener)

	return emitter
}

// emitEvent emits an event with optional data
// Usage: ghotona_prokash(emitter, "event_name", data1, data2, ...);
func emitEvent(args ...object.Object) object.Object {
	if len(args) < 2 {
		return &object.Error{Message: "ghotona_prokash() expects at least 2 arguments (emitter, event_name, ...data)"}
	}

	// Get emitter
	emitter, ok := args[0].(*object.EventEmitter)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("first argument must be EventEmitter, got %s", args[0].Type())}
	}

	// Get event name
	eventName, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("second argument must be string, got %s", args[1].Type())}
	}

	// Get event data (remaining arguments)
	eventData := args[2:]

	// Get listeners (thread-safe read)
	emitter.Mu.RLock()
	listeners := emitter.Events[eventName.Value]
	if listeners == nil {
		emitter.Mu.RUnlock()
		// No listeners - return true (event emitted successfully, no listeners to call)
		return object.TRUE
	}

	// Copy listeners to avoid holding lock during callback execution
	listenersCopy := make([]*object.EventListener, len(listeners))
	copy(listenersCopy, listeners)
	emitter.Mu.RUnlock()

	// Track indices of "once" listeners to remove
	var onceIndices []int

	// Call each listener
	for i, listener := range listenersCopy {
		// Call callback with event data
		switch callback := listener.Callback.(type) {
		case *object.Function:
			// Call user-defined function
			// Note: This requires access to evaluator, which we'll handle via callback
			if evalFunc != nil {
				evalFunc(callback, eventData)
			}
		case *object.Builtin:
			// Call built-in function
			callback.Fn(eventData...)
		}

		// Mark for removal if once
		if listener.Once {
			onceIndices = append(onceIndices, i)
		}
	}

	// Remove "once" listeners (thread-safe)
	if len(onceIndices) > 0 {
		emitter.Mu.Lock()
		// Get current listeners again (they may have changed)
		currentListeners := emitter.Events[eventName.Value]
		// Remove in reverse order to maintain indices
		for i := len(onceIndices) - 1; i >= 0; i-- {
			idx := onceIndices[i]
			if idx < len(currentListeners) {
				currentListeners = append(currentListeners[:idx], currentListeners[idx+1:]...)
			}
		}
		emitter.Events[eventName.Value] = currentListeners
		emitter.Mu.Unlock()
	}

	return object.TRUE
}

// removeEventListener removes a specific event listener
// Usage: ghotona_bondho(emitter, "event_name", callback);
func removeEventListener(args ...object.Object) object.Object {
	if len(args) != 3 {
		return &object.Error{Message: "ghotona_bondho() expects 3 arguments (emitter, event_name, callback)"}
	}

	// Get emitter
	emitter, ok := args[0].(*object.EventEmitter)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("first argument must be EventEmitter, got %s", args[0].Type())}
	}

	// Get event name
	eventName, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("second argument must be string, got %s", args[1].Type())}
	}

	// Get callback to remove
	callbackToRemove := args[2]

	// Remove listener (thread-safe)
	emitter.Mu.Lock()
	defer emitter.Mu.Unlock()

	listeners := emitter.Events[eventName.Value]
	if listeners == nil {
		return emitter
	}

	// Find and remove matching callback
	newListeners := make([]*object.EventListener, 0, len(listeners))
	for _, listener := range listeners {
		// Simple pointer comparison (works for same function reference)
		if listener.Callback != callbackToRemove {
			newListeners = append(newListeners, listener)
		}
	}

	if len(newListeners) == 0 {
		delete(emitter.Events, eventName.Value)
	} else {
		emitter.Events[eventName.Value] = newListeners
	}

	return emitter
}

// removeAllListeners removes all listeners for an event or all events
// Usage: ghotona_sob_bondho(emitter);        // Remove all listeners
//        ghotona_sob_bondho(emitter, "event"); // Remove listeners for specific event
func removeAllListeners(args ...object.Object) object.Object {
	if len(args) < 1 || len(args) > 2 {
		return &object.Error{Message: "ghotona_sob_bondho() expects 1 or 2 arguments (emitter, [event_name])"}
	}

	// Get emitter
	emitter, ok := args[0].(*object.EventEmitter)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("first argument must be EventEmitter, got %s", args[0].Type())}
	}

	emitter.Mu.Lock()
	defer emitter.Mu.Unlock()

	// If event name provided, remove only that event's listeners
	if len(args) == 2 {
		eventName, ok := args[1].(*object.String)
		if !ok {
			return &object.Error{Message: fmt.Sprintf("second argument must be string, got %s", args[1].Type())}
		}
		delete(emitter.Events, eventName.Value)
	} else {
		// Remove all listeners
		emitter.Events = make(map[string][]*object.EventListener)
	}

	return emitter
}

// getListeners returns all listeners for a specific event
// Usage: dhoro listeners = ghotona_shrotara(emitter, "event_name");
func getListeners(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "ghotona_shrotara() expects 2 arguments (emitter, event_name)"}
	}

	// Get emitter
	emitter, ok := args[0].(*object.EventEmitter)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("first argument must be EventEmitter, got %s", args[0].Type())}
	}

	// Get event name
	eventName, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("second argument must be string, got %s", args[1].Type())}
	}

	emitter.Mu.RLock()
	defer emitter.Mu.RUnlock()

	listeners := emitter.Events[eventName.Value]
	if listeners == nil {
		return &object.Array{Elements: []object.Object{}}
	}

	// Return array of callbacks
	callbacks := make([]object.Object, len(listeners))
	for i, listener := range listeners {
		callbacks[i] = listener.Callback
	}

	return &object.Array{Elements: callbacks}
}

// getEventNames returns all event names that have listeners
// Usage: dhoro events = ghotona_naam_sob(emitter);
func getEventNames(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "ghotona_naam_sob() expects 1 argument (emitter)"}
	}

	// Get emitter
	emitter, ok := args[0].(*object.EventEmitter)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("first argument must be EventEmitter, got %s", args[0].Type())}
	}

	emitter.Mu.RLock()
	defer emitter.Mu.RUnlock()

	// Collect event names
	eventNames := make([]object.Object, 0, len(emitter.Events))
	for eventName := range emitter.Events {
		eventNames = append(eventNames, &object.String{Value: eventName})
	}

	return &object.Array{Elements: eventNames}
}

// evalFunc is a callback to evaluate user-defined functions
// This will be set by the evaluator to allow calling user functions
var evalFunc func(fn *object.Function, args []object.Object) object.Object
