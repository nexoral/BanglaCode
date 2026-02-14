package builtins

import (
	"BanglaCode/src/object"
	"fmt"
)

// newError creates a new error object without position info
func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
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
