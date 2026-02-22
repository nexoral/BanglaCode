package builtins

import (
	"BanglaCode/src/object"

	"BanglaCode/src/evaluator/builtins/buffer"
	"BanglaCode/src/evaluator/builtins/collections"
	"BanglaCode/src/evaluator/builtins/crypto"
	"BanglaCode/src/evaluator/builtins/database"
	"BanglaCode/src/evaluator/builtins/errors"
	"BanglaCode/src/evaluator/builtins/events"
	mathpkg "BanglaCode/src/evaluator/builtins/math"
	"BanglaCode/src/evaluator/builtins/number"
	"BanglaCode/src/evaluator/builtins/streams"
	"BanglaCode/src/evaluator/builtins/system"
	"BanglaCode/src/evaluator/builtins/url"
	"BanglaCode/src/evaluator/builtins/worker"
)

// EvalFunc is a function pointer for evaluating AST nodes (set by evaluator.go to avoid circular dependency)
var EvalFunc func(handler *object.Function, args []object.Object) object.Object

// Builtins is the global map of built-in functions
// Individual built-in functions are registered in their respective files using init()
var Builtins = map[string]*object.Builtin{}

// InitializeEnvironmentWithConstants adds math, path, and number constants to the environment
func InitializeEnvironmentWithConstants(env *object.Environment) {
	// Add math constants
	for name, value := range mathpkg.Constants {
		numObj := &object.Number{Value: value}
		env.Set(name, numObj)
		env.SetConstant(name, numObj) // Mark as constant
	}

	// Add path constants
	for name, value := range system.PathConstants {
		strObj := &object.String{Value: value}
		env.Set(name, strObj)
		env.SetConstant(name, strObj) // Mark as constant
	}

	// Add number constants
	for name, value := range number.Constants {
		numObj := &object.Number{Value: value}
		env.Set(name, numObj)
		env.SetConstant(name, numObj) // Mark as constant
	}
}

func init() {
	// Register system built-in functions
	for name, fn := range system.Builtins {
		Builtins[name] = fn
	}

	// Register database built-in functions
	for name, fn := range database.Builtins {
		Builtins[name] = fn
	}

	// Register event built-in functions
	for name, fn := range events.Builtins {
		Builtins[name] = fn
	}

	// Register buffer built-in functions
	for name, fn := range buffer.Builtins {
		Builtins[name] = fn
	}

	// Register worker built-in functions
	for name, fn := range worker.Builtins {
		Builtins[name] = fn
	}

	// Register streams built-in functions
	for name, fn := range streams.Builtins {
		Builtins[name] = fn
	}

	// Register URL built-in functions
	for name, fn := range url.Builtins {
		Builtins[name] = fn
	}

	// Register Set built-in functions
	for name, fn := range collections.SetBuiltins {
		Builtins[name] = fn
	}

	// Register ES6 Map built-in functions
	for name, fn := range collections.MapBuiltins {
		Builtins[name] = fn
	}

	// Register Math constants and functions
	for name, fn := range mathpkg.Builtins {
		Builtins[name] = fn
	}

	// Number built-in functions
	for name, fn := range number.Builtins {
		Builtins[name] = fn
	}

	// Crypto built-in functions
	for name, fn := range crypto.Builtins {
		Builtins[name] = fn
	}

	// Error constructors and utilities
	for name, fn := range errors.Builtins {
		Builtins[name] = fn
	}
}
