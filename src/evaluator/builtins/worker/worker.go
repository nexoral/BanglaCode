package worker

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/object"
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	workerIDCounter int32
	evalFunc        func(node ast.Node, env *object.Environment) object.Object
	workerMu        sync.RWMutex
	workers         = make(map[int]*object.Worker)
)

// SetEvalFunc sets the evaluation function for executing worker code
func SetEvalFunc(fn func(ast.Node, *object.Environment) object.Object) {
	evalFunc = fn
}

// Builtins contains all worker-related built-in functions
var Builtins = map[string]*object.Builtin{
	"kaj_kormi_srishti": {
		Fn: kajKormiSrishti,
	},
	"kaj_kormi_pathao": {
		Fn: kajKormiPathao,
	},
	"kaj_kormi_bondho": {
		Fn: kajKormiBondho,
	},
	"kaj_kormi_shuno": {
		Fn: kajKormiShuno,
	},
}

// kajKormiSrishti creates a new worker thread
// Usage: dhoro worker = kaj_kormi_srishti(kaj() { ... }, initialData);
func kajKormiSrishti(args ...object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "kaj_kormi_srishti() requires at least 1 argument (function)"}
	}

	// First argument must be a function
	workerFn, ok := args[0].(*object.Function)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("kaj_kormi_srishti() first argument must be a function, got %s", args[0].Type())}
	}

	// Optional: initial worker data
	var workerData object.Object = object.NULL
	if len(args) > 1 {
		workerData = args[1]
	}

	// Generate unique worker ID
	id := int(atomic.AddInt32(&workerIDCounter, 1))

	// Create worker object
	worker := &object.Worker{
		ID:           id,
		MessageChan:  make(chan object.Object, 10), // Buffered channel
		ResponseChan: make(chan object.Object, 10),
		StopChan:     make(chan struct{}),
		IsRunning:    true,
		WorkerData:   workerData,
	}

	// Store worker in registry
	workerMu.Lock()
	workers[id] = worker
	workerMu.Unlock()

	// Start worker goroutine
	go runWorker(worker, workerFn)

	return worker
}

// runWorker runs the worker function in a separate goroutine
func runWorker(worker *object.Worker, workerFn *object.Function) {
	defer func() {
		// Cleanup on worker exit
		worker.Mu.Lock()
		worker.IsRunning = false
		worker.Mu.Unlock()

		workerMu.Lock()
		delete(workers, worker.ID)
		workerMu.Unlock()

		// Close channels
		close(worker.MessageChan)
		close(worker.ResponseChan)

		// Recover from panics
		if r := recover(); r != nil {
			fmt.Printf("Worker %d panicked: %v\n", worker.ID, r)
		}
	}()

	// Create worker environment with workerData
	workerEnv := object.NewEnvironment()
	workerEnv.Set("kaj_kormi_tothya", worker.WorkerData) // workerData accessible in worker

	// Set up self-reference for postMessage from within worker
	workerEnv.Set("pathao_message", &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return &object.Error{Message: "pathao_message() requires 1 argument"}
			}
			// Send message back to parent
			select {
			case worker.ResponseChan <- args[0]:
				return object.NULL
			case <-worker.StopChan:
				return &object.Error{Message: "Worker terminated"}
			}
		},
	})

	// Execute worker function with initial data
	if evalFunc != nil {
		// Create a call expression to execute the worker function
		callExpr := &ast.CallExpression{
			Function:  &ast.Identifier{Value: "workerFn"},
			Arguments: []ast.Expression{},
		}

		// Set the function in environment
		workerEnv.Set("workerFn", workerFn)

		// Execute the function
		evalFunc(callExpr, workerEnv)
	}

	// Message loop: listen for messages from parent
	for {
		select {
		case msg, ok := <-worker.MessageChan:
			if !ok {
				return // Channel closed
			}

			// If worker has onMessage handler, call it
			if worker.OnMessage != nil && evalFunc != nil {
				// Create call expression for handler
				handlerCall := &ast.CallExpression{
					Function:  &ast.Identifier{Value: "messageHandler"},
					Arguments: []ast.Expression{&ast.Identifier{Value: "messageData"}},
				}

				// Set handler and message in environment
				workerEnv.Set("messageHandler", worker.OnMessage)
				workerEnv.Set("messageData", msg)

				// Call handler
				evalFunc(handlerCall, workerEnv)
			}

		case <-worker.StopChan:
			return // Worker terminated
		}
	}
}

// kajKormiPathao sends a message to a worker
// Usage: kaj_kormi_pathao(worker, data);
func kajKormiPathao(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "kaj_kormi_pathao() requires 2 arguments (worker, message)"}
	}

	worker, ok := args[0].(*object.Worker)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("kaj_kormi_pathao() first argument must be a Worker, got %s", args[0].Type())}
	}

	worker.Mu.RLock()
	if !worker.IsRunning {
		worker.Mu.RUnlock()
		return &object.Error{Message: "Cannot send message to terminated worker"}
	}
	worker.Mu.RUnlock()

	// Send message to worker
	select {
	case worker.MessageChan <- args[1]:
		return object.NULL
	default:
		return &object.Error{Message: "Worker message channel is full"}
	}
}

// kajKormiBondho terminates a worker
// Usage: kaj_kormi_bondho(worker);
func kajKormiBondho(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "kaj_kormi_bondho() requires 1 argument (worker)"}
	}

	worker, ok := args[0].(*object.Worker)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("kaj_kormi_bondho() argument must be a Worker, got %s", args[0].Type())}
	}

	worker.Mu.Lock()
	if !worker.IsRunning {
		worker.Mu.Unlock()
		return object.NULL // Already terminated
	}
	worker.IsRunning = false
	worker.Mu.Unlock()

	// Signal worker to stop
	close(worker.StopChan)

	return object.NULL
}

// kajKormiShuno sets up a message handler for worker responses
// Usage: kaj_kormi_shuno(worker, kaj(data) { dekho(data); });
func kajKormiShuno(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "kaj_kormi_shuno() requires 2 arguments (worker, callback)"}
	}

	worker, ok := args[0].(*object.Worker)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("kaj_kormi_shuno() first argument must be a Worker, got %s", args[0].Type())}
	}

	callback, ok := args[1].(*object.Function)
	if !ok {
		return &object.Error{Message: fmt.Sprintf("kaj_kormi_shuno() second argument must be a function, got %s", args[1].Type())}
	}

	worker.Mu.Lock()
	worker.OnMessage = callback
	worker.Mu.Unlock()

	// Start goroutine to listen for worker responses
	go func() {
		for {
			select {
			case msg, ok := <-worker.ResponseChan:
				if !ok {
					return // Channel closed
				}

				// Call callback with message
				if evalFunc != nil && callback != nil {
					// Create temporary environment for callback
					callbackEnv := object.NewEnvironment()

					// Create call expression
					callExpr := &ast.CallExpression{
						Function:  &ast.Identifier{Value: "responseCallback"},
						Arguments: []ast.Expression{&ast.Identifier{Value: "responseData"}},
					}

					// Set callback and message in environment
					callbackEnv.Set("responseCallback", callback)
					callbackEnv.Set("responseData", msg)

					// Execute callback
					evalFunc(callExpr, callbackEnv)
				}

			case <-worker.StopChan:
				return // Worker terminated
			}
		}
	}()

	return object.NULL
}
