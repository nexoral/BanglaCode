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
			catchEnv.Set(tcs.CatchParam.Value, &object.String{Value: exception.Message})
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

	// Convert to exception
	var message string
	switch v := value.(type) {
	case *object.String:
		message = v.Value
	default:
		message = v.Inspect()
	}

	return &object.Exception{Message: message}
}
