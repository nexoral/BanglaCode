package math

import (
	"BanglaCode/src/object"
	"math"
)

// Builtins contains all Math-related built-in functions and constants
var Builtins = map[string]*object.Builtin{
	// Trigonometric functions
	"math_sin":   {Fn: mathSin},
	"math_cos":   {Fn: mathCos},
	"math_tan":   {Fn: mathTan},
	"math_asin":  {Fn: mathAsin},
	"math_acos":  {Fn: mathAcos},
	"math_atan":  {Fn: mathAtan},
	"math_atan2": {Fn: mathAtan2},

	// Hyperbolic functions
	"math_sinh":  {Fn: mathSinh},
	"math_cosh":  {Fn: mathCosh},
	"math_tanh":  {Fn: mathTanh},
	"math_asinh": {Fn: mathAsinh},
	"math_acosh": {Fn: mathAcosh},
	"math_atanh": {Fn: mathAtanh},

	// Logarithmic & exponential functions
	"math_log":   {Fn: mathLog},
	"math_log10": {Fn: mathLog10},
	"math_log2":  {Fn: mathLog2},
	"math_log1p": {Fn: mathLog1p},
	"math_exp":   {Fn: mathExp},
	"math_expm1": {Fn: mathExpm1},

	// Utility functions
	"math_imul":   {Fn: mathImul},
	"math_clz32":  {Fn: mathClz32},
	"math_fround": {Fn: mathFround},
	"math_hypot":  {Fn: mathHypot},
}

// Constants (to be added to global environment)
var Constants = map[string]float64{
	"MATH_PI":      math.Pi,
	"MATH_E":       math.E,
	"MATH_LN2":     math.Ln2,
	"MATH_LN10":    math.Ln10,
	"MATH_LOG2E":   math.Log2E,
	"MATH_LOG10E":  math.Log10E,
	"MATH_SQRT1_2": math.Sqrt2 / 2,
	"MATH_SQRT2":   math.Sqrt2,
}

// Helper to get number from object
func getNumber(obj object.Object) (float64, bool) {
	if num, ok := obj.(*object.Number); ok {
		return num.Value, true
	}
	return 0, false
}

// Trigonometric Functions

// mathSin returns the sine of x (x in radians)
func mathSin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_sin() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_sin() argument must be a number"}
	}

	return &object.Number{Value: math.Sin(num)}
}

// mathCos returns the cosine of x (x in radians)
func mathCos(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_cos() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_cos() argument must be a number"}
	}

	return &object.Number{Value: math.Cos(num)}
}

// mathTan returns the tangent of x (x in radians)
func mathTan(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_tan() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_tan() argument must be a number"}
	}

	return &object.Number{Value: math.Tan(num)}
}

// mathAsin returns the arcsine of x (result in radians)
func mathAsin(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_asin() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_asin() argument must be a number"}
	}

	return &object.Number{Value: math.Asin(num)}
}

// mathAcos returns the arccosine of x (result in radians)
func mathAcos(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_acos() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_acos() argument must be a number"}
	}

	return &object.Number{Value: math.Acos(num)}
}

// mathAtan returns the arctangent of x (result in radians)
func mathAtan(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_atan() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_atan() argument must be a number"}
	}

	return &object.Number{Value: math.Atan(num)}
}

// mathAtan2 returns the arctangent of y/x (result in radians)
func mathAtan2(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "math_atan2() expects 2 arguments: y, x"}
	}

	y, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_atan2() first argument must be a number"}
	}

	x, ok := getNumber(args[1])
	if !ok {
		return &object.Error{Message: "math_atan2() second argument must be a number"}
	}

	return &object.Number{Value: math.Atan2(y, x)}
}

// Hyperbolic Functions

// mathSinh returns the hyperbolic sine of x
func mathSinh(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_sinh() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_sinh() argument must be a number"}
	}

	return &object.Number{Value: math.Sinh(num)}
}

// mathCosh returns the hyperbolic cosine of x
func mathCosh(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_cosh() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_cosh() argument must be a number"}
	}

	return &object.Number{Value: math.Cosh(num)}
}

// mathTanh returns the hyperbolic tangent of x
func mathTanh(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_tanh() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_tanh() argument must be a number"}
	}

	return &object.Number{Value: math.Tanh(num)}
}

// mathAsinh returns the inverse hyperbolic sine of x
func mathAsinh(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_asinh() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_asinh() argument must be a number"}
	}

	return &object.Number{Value: math.Asinh(num)}
}

// mathAcosh returns the inverse hyperbolic cosine of x
func mathAcosh(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_acosh() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_acosh() argument must be a number"}
	}

	return &object.Number{Value: math.Acosh(num)}
}

// mathAtanh returns the inverse hyperbolic tangent of x
func mathAtanh(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_atanh() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_atanh() argument must be a number"}
	}

	return &object.Number{Value: math.Atanh(num)}
}

// Logarithmic & Exponential Functions

// mathLog returns the natural logarithm of x
func mathLog(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_log() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_log() argument must be a number"}
	}

	return &object.Number{Value: math.Log(num)}
}

// mathLog10 returns the base-10 logarithm of x
func mathLog10(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_log10() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_log10() argument must be a number"}
	}

	return &object.Number{Value: math.Log10(num)}
}

// mathLog2 returns the base-2 logarithm of x
func mathLog2(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_log2() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_log2() argument must be a number"}
	}

	return &object.Number{Value: math.Log2(num)}
}

// mathLog1p returns the natural logarithm of 1 + x
func mathLog1p(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_log1p() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_log1p() argument must be a number"}
	}

	return &object.Number{Value: math.Log1p(num)}
}

// mathExp returns e raised to the power of x
func mathExp(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_exp() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_exp() argument must be a number"}
	}

	return &object.Number{Value: math.Exp(num)}
}

// mathExpm1 returns e^x - 1
func mathExpm1(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_expm1() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_expm1() argument must be a number"}
	}

	return &object.Number{Value: math.Expm1(num)}
}

// Utility Functions

// mathImul returns the 32-bit integer multiplication of a and b
func mathImul(args ...object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "math_imul() expects 2 arguments"}
	}

	a, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_imul() first argument must be a number"}
	}

	b, ok := getNumber(args[1])
	if !ok {
		return &object.Error{Message: "math_imul() second argument must be a number"}
	}

	// Convert to 32-bit integers and multiply
	result := int32(a) * int32(b)
	return &object.Number{Value: float64(result)}
}

// mathClz32 counts leading zero bits in 32-bit binary representation
func mathClz32(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_clz32() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_clz32() argument must be a number"}
	}

	// Convert to 32-bit unsigned integer
	val := uint32(num)

	// Count leading zeros
	if val == 0 {
		return &object.Number{Value: 32}
	}

	count := 0
	for i := 31; i >= 0; i-- {
		if (val & (1 << uint(i))) != 0 {
			break
		}
		count++
	}

	return &object.Number{Value: float64(count)}
}

// mathFround returns the nearest 32-bit single precision float representation
func mathFround(args ...object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "math_fround() expects 1 argument"}
	}

	num, ok := getNumber(args[0])
	if !ok {
		return &object.Error{Message: "math_fround() argument must be a number"}
	}

	// Convert to float32 and back to float64
	return &object.Number{Value: float64(float32(num))}
}

// mathHypot returns the square root of the sum of squares of its arguments
func mathHypot(args ...object.Object) object.Object {
	if len(args) == 0 {
		return &object.Number{Value: 0}
	}

	var sum float64
	for i, arg := range args {
		num, ok := getNumber(arg)
		if !ok {
			return &object.Error{Message: "math_hypot() argument " + string(rune(i+1)) + " must be a number"}
		}
		sum += num * num
	}

	return &object.Number{Value: math.Sqrt(sum)}
}
