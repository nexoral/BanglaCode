package events

import "BanglaCode/src/object"

// SetEvalFunc sets the function evaluator callback
// This is called by the main evaluator to allow event callbacks to execute user functions
func SetEvalFunc(fn func(*object.Function, []object.Object) object.Object) {
	evalFunc = fn
}
