package errors

import (
	"BanglaCode/src/object"
	"fmt"
)

// Builtins holds error-related built-in functions
var Builtins = map[string]*object.Builtin{}

func init() {
	// Register error constructor built-in functions

	// Error() - Generic error constructor
	Builtins["Error"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return object.NewTypeError("Error() requires at least 1 argument (message)")
			}

			message := ""
			if str, ok := args[0].(*object.String); ok {
				message = str.Value
			} else {
				message = args[0].Inspect()
			}

			// Create error object as a Map to make it accessible in BanglaCode
			errorMap := &object.Map{
				Pairs: map[string]object.Object{
					"message": &object.String{Value: message},
					"name":    &object.String{Value: "Error"},
					"stack":   &object.String{Value: ""}, // Will be populated when thrown
				},
			}

			return errorMap
		},
	}

	// TypeError() - Type error constructor
	Builtins["TypeError"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return object.NewTypeError("TypeError() requires at least 1 argument (message)")
			}

			message := ""
			if str, ok := args[0].(*object.String); ok {
				message = str.Value
			} else {
				message = args[0].Inspect()
			}

			errorMap := &object.Map{
				Pairs: map[string]object.Object{
					"message": &object.String{Value: message},
					"name":    &object.String{Value: "TypeError"},
					"stack":   &object.String{Value: ""},
				},
			}

			return errorMap
		},
	}

	// ReferenceError() - Reference error constructor
	Builtins["ReferenceError"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return object.NewTypeError("ReferenceError() requires at least 1 argument (message)")
			}

			message := ""
			if str, ok := args[0].(*object.String); ok {
				message = str.Value
			} else {
				message = args[0].Inspect()
			}

			errorMap := &object.Map{
				Pairs: map[string]object.Object{
					"message": &object.String{Value: message},
					"name":    &object.String{Value: "ReferenceError"},
					"stack":   &object.String{Value: ""},
				},
			}

			return errorMap
		},
	}

	// RangeError() - Range error constructor
	Builtins["RangeError"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return object.NewTypeError("RangeError() requires at least 1 argument (message)")
			}

			message := ""
			if str, ok := args[0].(*object.String); ok {
				message = str.Value
			} else {
				message = args[0].Inspect()
			}

			errorMap := &object.Map{
				Pairs: map[string]object.Object{
					"message": &object.String{Value: message},
					"name":    &object.String{Value: "RangeError"},
					"stack":   &object.String{Value: ""},
				},
			}

			return errorMap
		},
	}

	// SyntaxError() - Syntax error constructor
	Builtins["SyntaxError"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return object.NewTypeError("SyntaxError() requires at least 1 argument (message)")
			}

			message := ""
			if str, ok := args[0].(*object.String); ok {
				message = str.Value
			} else {
				message = args[0].Inspect()
			}

			errorMap := &object.Map{
				Pairs: map[string]object.Object{
					"message": &object.String{Value: message},
					"name":    &object.String{Value: "SyntaxError"},
					"stack":   &object.String{Value: ""},
				},
			}

			return errorMap
		},
	}

	// bhul_message() - Get error message (বুল = error, message = message)
	Builtins["bhul_message"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return object.NewTypeError("bhul_message() requires exactly 1 argument (error object)")
			}

			// If it's a Map (error object), get the message
			if errorMap, ok := args[0].(*object.Map); ok {
				if msg, exists := errorMap.Pairs["message"]; exists {
					return msg
				}
			}

			// If it's an Error object, get the message
			if err, ok := args[0].(*object.Error); ok {
				return &object.String{Value: err.Message}
			}

			// If it's an Exception, get the message
			if exc, ok := args[0].(*object.Exception); ok {
				return &object.String{Value: exc.Message}
			}

			return &object.String{Value: args[0].Inspect()}
		},
	}

	// bhul_stack() - Get error stack trace (স্ট্যাক = stack)
	Builtins["bhul_stack"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return object.NewTypeError("bhul_stack() requires exactly 1 argument (error object)")
			}

			// If it's a Map (error object), get the stack
			if errorMap, ok := args[0].(*object.Map); ok {
				if stack, exists := errorMap.Pairs["stack"]; exists {
					return stack
				}
			}

			// If it's an Error object with stack trace
			if err, ok := args[0].(*object.Error); ok {
				return &object.String{Value: err.GetStack()}
			}

			// If it's an Exception, return message
			if exc, ok := args[0].(*object.Exception); ok {
				return &object.String{Value: exc.Inspect()}
			}

			return &object.String{Value: args[0].Inspect()}
		},
	}

	// bhul_naam() - Get error name/type (নাম = name)
	Builtins["bhul_naam"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return object.NewTypeError("bhul_naam() requires exactly 1 argument (error object)")
			}

			// If it's a Map (error object), get the name
			if errorMap, ok := args[0].(*object.Map); ok {
				if name, exists := errorMap.Pairs["name"]; exists {
					return name
				}
			}

			// If it's an Error object, get the type name
			if err, ok := args[0].(*object.Error); ok {
				switch err.ErrorType {
				case object.TYPE_ERROR_OBJ:
					return &object.String{Value: "TypeError"}
				case object.REFERENCE_ERROR_OBJ:
					return &object.String{Value: "ReferenceError"}
				case object.RANGE_ERROR_OBJ:
					return &object.String{Value: "RangeError"}
				case object.SYNTAX_ERROR_OBJ:
					return &object.String{Value: "SyntaxError"}
				default:
					return &object.String{Value: "Error"}
				}
			}

			return &object.String{Value: "Error"}
		},
	}

	// is_error() - Check if object is an error
	Builtins["is_error"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return object.FALSE
			}

			// Check if it's an error Map
			if errorMap, ok := args[0].(*object.Map); ok {
				if name, exists := errorMap.Pairs["name"]; exists {
					if nameStr, ok := name.(*object.String); ok {
						errorTypes := []string{"Error", "TypeError", "ReferenceError", "RangeError", "SyntaxError"}
						for _, errorType := range errorTypes {
							if nameStr.Value == errorType {
								return object.TRUE
							}
						}
					}
				}
			}

			// Check if it's an Error object
			if _, ok := args[0].(*object.Error); ok {
				return object.TRUE
			}

			// Check if it's an Exception
			if _, ok := args[0].(*object.Exception); ok {
				return object.TRUE
			}

			return object.FALSE
		},
	}
}

// CreateErrorWithStack creates an error with stack trace
func CreateErrorWithStack(message string, errorType object.ObjectType, functionName string, line, column int) *object.Error {
	err := &object.Error{
		Message:   message,
		ErrorType: errorType,
		Line:      line,
		Column:    column,
	}

	if functionName != "" || line > 0 {
		err.AddStackFrame(functionName, "", line, column)
	}

	return err
}

// ConvertMapToError converts error Map to Error object for internal use
func ConvertMapToError(errorMap *object.Map) *object.Error {
	message := ""
	var errorType object.ObjectType = object.ERROR_OBJ
	stack := ""

	if msg, exists := errorMap.Pairs["message"]; exists {
		if msgStr, ok := msg.(*object.String); ok {
			message = msgStr.Value
		}
	}

	if name, exists := errorMap.Pairs["name"]; exists {
		if nameStr, ok := name.(*object.String); ok {
			switch nameStr.Value {
			case "TypeError":
				errorType = object.TYPE_ERROR_OBJ
			case "ReferenceError":
				errorType = object.REFERENCE_ERROR_OBJ
			case "RangeError":
				errorType = object.RANGE_ERROR_OBJ
			case "SyntaxError":
				errorType = object.SYNTAX_ERROR_OBJ
			}
		}
	}

	if stackObj, exists := errorMap.Pairs["stack"]; exists {
		if stackStr, ok := stackObj.(*object.String); ok {
			stack = stackStr.Value
		}
	}

	err := &object.Error{
		Message:   message,
		ErrorType: errorType,
	}

	// Parse stack if available (simplified - in production would parse properly)
	if stack != "" {
		err.Stack = []object.StackFrame{
			{Function: "<from stack>", File: "", Line: 0, Column: 0},
		}
	}

	return err
}

// Helper function to create stack trace string
func CreateStackTrace(frames []object.StackFrame) string {
	if len(frames) == 0 {
		return ""
	}

	result := "Stack trace:\n"
	for _, frame := range frames {
		if frame.Function != "" {
			result += fmt.Sprintf("  at %s", frame.Function)
		} else {
			result += "  at <anonymous>"
		}

		if frame.File != "" {
			result += fmt.Sprintf(" (%s:%d:%d)\n", frame.File, frame.Line, frame.Column)
		} else if frame.Line > 0 {
			result += fmt.Sprintf(" (line %d, col %d)\n", frame.Line, frame.Column)
		} else {
			result += "\n"
		}
	}

	return result
}
