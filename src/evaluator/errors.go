package evaluator

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/object"
)

// evalTryCatchStatement evaluates try-catch statements
func evalTryCatchStatement(tcs *ast.TryCatchStatement, env *object.Environment) object.Object {
	// Execute try block
	result := Eval(tcs.TryBlock, env)

	// Check if an exception was thrown
	if exception, ok := result.(*object.Exception); ok {
		// Create catch block environment with error variable
		catchEnv := object.NewEnclosedEnvironment(env)
		if tcs.CatchParam != nil {
			// If exception has a Value (like an error Map), use that
			// Otherwise use the exception message as a string
			if exception.Value != nil {
				catchEnv.Set(tcs.CatchParam.Value, exception.Value)
			} else {
				catchEnv.Set(tcs.CatchParam.Value, &object.String{Value: exception.Message})
			}
		}

		// Execute catch block
		result = Eval(tcs.CatchBlock, catchEnv)
	}

	// Execute finally block if present
	if tcs.FinallyBlock != nil {
		finallyResult := Eval(tcs.FinallyBlock, env)
		// Finally block errors take precedence
		if isError(finallyResult) || isException(finallyResult) {
			return finallyResult
		}
	}

	return result
}

// evalThrowStatement evaluates throw statements
func evalThrowStatement(ts *ast.ThrowStatement, env *object.Environment) object.Object {
	value := Eval(ts.Value, env)
	if isError(value) {
		return value
	}

	// If it's an error Map (created by Error(), TypeError(), etc.), add stack trace
	if errorMap, ok := value.(*object.Map); ok {
		if name, exists := errorMap.Pairs["name"]; exists {
			if nameStr, ok := name.(*object.String); ok {
				// Check if it's an error type
				errorTypes := []string{"Error", "TypeError", "ReferenceError", "RangeError", "SyntaxError"}
				for _, errorType := range errorTypes {
					if nameStr.Value == errorType {
						// Add stack trace information
						stackTrace := "Stack trace:\n  at <throw statement>"
						if ts.Token.Line > 0 {
							stackTrace += " (line " + string(rune(ts.Token.Line)) + ")"
						}
						errorMap.Pairs["stack"] = &object.String{Value: stackTrace}

						// Convert to exception for throwing
						var message string
						if msg, exists := errorMap.Pairs["message"]; exists {
							if msgStr, ok := msg.(*object.String); ok {
								message = nameStr.Value + ": " + msgStr.Value
							}
						}
						return &object.Exception{Message: message, Value: errorMap}
					}
				}
			}
		}
	}

	// Convert to exception
	var message string
	switch v := value.(type) {
	case *object.String:
		message = v.Value
	default:
		message = v.Inspect()
	}

	return &object.Exception{Message: message, Value: value}
}
