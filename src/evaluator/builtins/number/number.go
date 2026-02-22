package number

import (
	"BanglaCode/src/object"
	"fmt"
	"math"
)

// Builtins contains all Number-related built-in functions
var Builtins = map[string]*object.Builtin{
	// Number.isFinite() - Check if number is finite
	"sonkhya_sesh": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("sonkhya_sesh requires 1 argument (number)")
			}

			num, ok := args[0].(*object.Number)
			if !ok {
				// Non-number arguments return false
				return object.FALSE
			}

			// Check if finite (not infinity and not NaN)
			if math.IsInf(num.Value, 0) || math.IsNaN(num.Value) {
				return object.FALSE
			}
			return object.TRUE
		},
	},

	// Number.isInteger() - Check if number is an integer
	"sonkhya_purno": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("sonkhya_purno requires 1 argument (number)")
			}

			num, ok := args[0].(*object.Number)
			if !ok {
				// Non-number arguments return false
				return object.FALSE
			}

			// Check if finite first (NaN and Infinity are not integers)
			if math.IsInf(num.Value, 0) || math.IsNaN(num.Value) {
				return object.FALSE
			}

			// Check if integer (no fractional part)
			if num.Value == math.Trunc(num.Value) {
				return object.TRUE
			}
			return object.FALSE
		},
	},

	// Number.isNaN() - Check if value is NaN
	"sonkhya_na_check": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("sonkhya_na_check requires 1 argument")
			}

			num, ok := args[0].(*object.Number)
			if !ok {
				// Non-number arguments return false (strict check)
				return object.FALSE
			}

			if math.IsNaN(num.Value) {
				return object.TRUE
			}
			return object.FALSE
		},
	},

	// Number.isSafeInteger() - Check if number is a safe integer
	"sonkhya_nirapod": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("sonkhya_nirapod requires 1 argument (number)")
			}

			num, ok := args[0].(*object.Number)
			if !ok {
				// Non-number arguments return false
				return object.FALSE
			}

			// Check if finite and integer first
			if math.IsInf(num.Value, 0) || math.IsNaN(num.Value) {
				return object.FALSE
			}
			if num.Value != math.Trunc(num.Value) {
				return object.FALSE
			}

			// Check if within safe integer range
			// Safe integer range: -(2^53 - 1) to (2^53 - 1)
			maxSafe := 9007199254740991.0  // 2^53 - 1
			minSafe := -9007199254740991.0 // -(2^53 - 1)

			if num.Value >= minSafe && num.Value <= maxSafe {
				return object.TRUE
			}
			return object.FALSE
		},
	},
}

// Constants contains all Number-related constants
var Constants = map[string]float64{
	"NUMBER_MAX_SAFE_INTEGER":  9007199254740991.0,          // 2^53 - 1
	"NUMBER_MIN_SAFE_INTEGER":  -9007199254740991.0,         // -(2^53 - 1)
	"NUMBER_MAX_VALUE":         math.MaxFloat64,             // ~1.8e308
	"NUMBER_MIN_VALUE":         math.SmallestNonzeroFloat64, // ~5e-324
	"NUMBER_POSITIVE_INFINITY": math.Inf(1),                 // +Infinity
	"NUMBER_NEGATIVE_INFINITY": math.Inf(-1),                // -Infinity
	"NUMBER_EPSILON":           2.220446049250313e-16,       // Smallest difference
	"NUMBER_NAN":               math.NaN(),                  // NaN value
}

// Helper function to create error objects
func newError(format string, args ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, args...)}
}
