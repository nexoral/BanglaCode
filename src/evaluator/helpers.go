package evaluator

import (
	"BanglaCode/src/object"
	"fmt"
)

// newError creates a new error object without position info
func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

// newErrorAt creates a new error object with line and column info
func newErrorAt(line, column int, format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...), Line: line, Column: column}
}

// isError checks if an object is an error
func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}
	return false
}

// isException checks if an object is an exception
func isException(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.EXCEPTION_OBJ
	}
	return false
}

// isTruthy determines if an object is truthy
func isTruthy(obj object.Object) bool {
	switch obj {
	case object.NULL:
		return false
	case object.TRUE:
		return true
	case object.FALSE:
		return false
	default:
		return true
	}
}

// objectsEqual checks if two objects are equal
func objectsEqual(left, right object.Object) bool {
	if left.Type() != right.Type() {
		return false
	}

	switch l := left.(type) {
	case *object.Number:
		return l.Value == right.(*object.Number).Value
	case *object.String:
		return l.Value == right.(*object.String).Value
	case *object.Boolean:
		return l.Value == right.(*object.Boolean).Value
	case *object.Null:
		return true
	default:
		return left == right
	}
}

// unwrapReturnValue extracts the value from a return object
func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}
	return obj
}

// getMapKey converts an object to a string key for map access
func getMapKey(key object.Object) string {
	switch k := key.(type) {
	case *object.String:
		return k.Value
	case *object.Number:
		return k.Inspect()
	default:
		return ""
	}
}
