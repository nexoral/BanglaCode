package evaluator

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/object"
	"fmt"
	"time"
)

// evalAsyncFunctionLiteral evaluates an async function literal and creates an async function object
func evalAsyncFunctionLiteral(node *ast.AsyncFunctionLiteral, env *object.Environment) object.Object {
	params := node.Parameters
	body := node.Body
	name := ""
	if node.Name != nil {
		name = node.Name.Value
	}

	fn := &object.Function{
		Parameters:    params,
		RestParameter: node.RestParameter,
		Env:           env,
		Body:          body,
		Name:          name,
		IsAsync:       true, // Mark as async
	}

	// If function has a name, bind it in the environment
	if name != "" {
		env.Set(name, fn)
	}

	return fn
}

// evalAsyncFunctionCall executes an async function in a goroutine and returns a promise
func evalAsyncFunctionCall(fn *object.Function, args []object.Object, env *object.Environment) object.Object {
	promise := object.CreatePromise()

	// Spawn goroutine to execute async function
	go func() {
		// Recover from panics in async functions
		defer func() {
			if r := recover(); r != nil {
				errorMsg := fmt.Sprintf("panic in async function: %v", r)
				object.RejectPromise(promise, &object.Error{Message: errorMsg})
			}
		}()

		// Create new environment for function execution
		extendedEnv := extendFunctionEnv(fn, args)

		// Execute function body
		result := Eval(fn.Body, extendedEnv)

		// Unwrap return value
		if returnValue, ok := result.(*object.ReturnValue); ok {
			result = returnValue.Value
		}

		// Check for errors or exceptions
		if err, ok := result.(*object.Error); ok {
			object.RejectPromise(promise, err)
			return
		}

		if exc, ok := result.(*object.Exception); ok {
			object.RejectPromise(promise, exc)
			return
		}

		// Resolve promise with result
		object.ResolvePromise(promise, result)
	}()

	return promise
}

// evalAwaitExpression waits for a promise to resolve or reject
// FIXED: Removed race condition by relying on channel read directly
// ADDED: 30-second timeout to prevent infinite blocking
func evalAwaitExpression(node *ast.AwaitExpression, env *object.Environment) object.Object {
	// Evaluate the expression that should produce a promise
	value := Eval(node.Expression, env)
	if isError(value) {
		return value
	}

	// Value must be a promise
	promise, ok := value.(*object.Promise)
	if !ok {
		return newError("opekha (await) can only be used with promises, got %s", value.Type())
	}

	// Wait for promise to complete with timeout
	// Channels are thread-safe, no need to check state first
	select {
	case result := <-promise.ResultChan:
		return result
	case err := <-promise.ErrorChan:
		return err
	case <-time.After(30 * time.Second):
		return newError("await timeout: promise did not resolve within 30 seconds")
	}
}
