package test

import (
	"math"
	"testing"

	"BanglaCode/src/evaluator"
	"BanglaCode/src/evaluator/builtins"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
)

// Helper function to evaluate input code
func evalMathInput(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	builtins.InitializeEnvironmentWithConstants(env)
	return evaluator.Eval(program, env)
}

// Helper to check if float values are approximately equal
func floatEquals(a, b, epsilon float64) bool {
	return math.Abs(a-b) < epsilon
}

// Test Math Constants

func TestMathConstants(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"MATH_PI", math.Pi},
		{"MATH_E", math.E},
		{"MATH_LN2", math.Ln2},
		{"MATH_LN10", math.Ln10},
		{"MATH_LOG2E", math.Log2E},
		{"MATH_LOG10E", math.Log10E},
		{"MATH_SQRT2", math.Sqrt2},
		{"MATH_SQRT1_2", math.Sqrt2 / 2},
	}

	for _, tt := range tests {
		result := evalMathInput(tt.input)
		num, ok := result.(*object.Number)
		if !ok {
			t.Errorf("Expected Number for %s, got %T", tt.input, result)
			continue
		}

		if !floatEquals(num.Value, tt.expected, 1e-10) {
			t.Errorf("Expected %s = %f, got %f", tt.input, tt.expected, num.Value)
		}
	}
}

// Test Trigonometric Functions

func TestMathSin(t *testing.T) {
	input := `math_sin(MATH_PI / 2)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if !floatEquals(num.Value, 1.0, 1e-10) {
		t.Errorf("Expected sin(PI/2) = 1.0, got %f", num.Value)
	}
}

func TestMathCos(t *testing.T) {
	input := `math_cos(0)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if !floatEquals(num.Value, 1.0, 1e-10) {
		t.Errorf("Expected cos(0) = 1.0, got %f", num.Value)
	}
}

func TestMathTan(t *testing.T) {
	input := `math_tan(MATH_PI / 4)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if !floatEquals(num.Value, 1.0, 1e-10) {
		t.Errorf("Expected tan(PI/4) = 1.0, got %f", num.Value)
	}
}

func TestMathAsin(t *testing.T) {
	input := `math_asin(0.5)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	expected := math.Asin(0.5)
	if !floatEquals(num.Value, expected, 1e-10) {
		t.Errorf("Expected asin(0.5) = %f, got %f", expected, num.Value)
	}
}

func TestMathAcos(t *testing.T) {
	input := `math_acos(0.5)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	expected := math.Acos(0.5)
	if !floatEquals(num.Value, expected, 1e-10) {
		t.Errorf("Expected acos(0.5) = %f, got %f", expected, num.Value)
	}
}

func TestMathAtan(t *testing.T) {
	input := `math_atan(1)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	expected := math.Pi / 4
	if !floatEquals(num.Value, expected, 1e-10) {
		t.Errorf("Expected atan(1) = PI/4, got %f", num.Value)
	}
}

func TestMathAtan2(t *testing.T) {
	input := `math_atan2(1, 1)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	expected := math.Pi / 4
	if !floatEquals(num.Value, expected, 1e-10) {
		t.Errorf("Expected atan2(1,1) = PI/4, got %f", num.Value)
	}
}

// Test Hyperbolic Functions

func TestMathSinh(t *testing.T) {
	input := `math_sinh(1)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	expected := math.Sinh(1)
	if !floatEquals(num.Value, expected, 1e-10) {
		t.Errorf("Expected sinh(1) = %f, got %f", expected, num.Value)
	}
}

func TestMathCosh(t *testing.T) {
	input := `math_cosh(0)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if !floatEquals(num.Value, 1.0, 1e-10) {
		t.Errorf("Expected cosh(0) = 1.0, got %f", num.Value)
	}
}

func TestMathTanh(t *testing.T) {
	input := `math_tanh(0)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if !floatEquals(num.Value, 0.0, 1e-10) {
		t.Errorf("Expected tanh(0) = 0.0, got %f", num.Value)
	}
}

func TestMathAsinh(t *testing.T) {
	input := `math_asinh(1)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	expected := math.Asinh(1)
	if !floatEquals(num.Value, expected, 1e-10) {
		t.Errorf("Expected asinh(1) = %f, got %f", expected, num.Value)
	}
}

func TestMathAcosh(t *testing.T) {
	input := `math_acosh(2)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	expected := math.Acosh(2)
	if !floatEquals(num.Value, expected, 1e-10) {
		t.Errorf("Expected acosh(2) = %f, got %f", expected, num.Value)
	}
}

func TestMathAtanh(t *testing.T) {
	input := `math_atanh(0.5)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	expected := math.Atanh(0.5)
	if !floatEquals(num.Value, expected, 1e-10) {
		t.Errorf("Expected atanh(0.5) = %f, got %f", expected, num.Value)
	}
}

// Test Logarithmic & Exponential Functions

func TestMathLog(t *testing.T) {
	input := `math_log(MATH_E)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if !floatEquals(num.Value, 1.0, 1e-10) {
		t.Errorf("Expected log(e) = 1.0, got %f", num.Value)
	}
}

func TestMathLog10(t *testing.T) {
	input := `math_log10(100)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if !floatEquals(num.Value, 2.0, 1e-10) {
		t.Errorf("Expected log10(100) = 2.0, got %f", num.Value)
	}
}

func TestMathLog2(t *testing.T) {
	input := `math_log2(8)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if !floatEquals(num.Value, 3.0, 1e-10) {
		t.Errorf("Expected log2(8) = 3.0, got %f", num.Value)
	}
}

func TestMathLog1p(t *testing.T) {
	input := `math_log1p(0)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if !floatEquals(num.Value, 0.0, 1e-10) {
		t.Errorf("Expected log1p(0) = 0.0, got %f", num.Value)
	}
}

func TestMathExp(t *testing.T) {
	input := `math_exp(0)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if !floatEquals(num.Value, 1.0, 1e-10) {
		t.Errorf("Expected exp(0) = 1.0, got %f", num.Value)
	}
}

func TestMathExpm1(t *testing.T) {
	input := `math_expm1(0)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if !floatEquals(num.Value, 0.0, 1e-10) {
		t.Errorf("Expected expm1(0) = 0.0, got %f", num.Value)
	}
}

// Test Utility Functions

func TestMathImul(t *testing.T) {
	input := `math_imul(2, 4)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if num.Value != 8 {
		t.Errorf("Expected imul(2, 4) = 8, got %f", num.Value)
	}
}

func TestMathImulOverflow(t *testing.T) {
	// Test 32-bit multiplication with large numbers
	input := `math_imul(100000, 50000)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	// imul does 32-bit integer multiplication
	a := int32(100000)
	b := int32(50000)
	expected := float64(a * b)
	if num.Value != expected {
		t.Errorf("Expected imul result %f, got %f", expected, num.Value)
	}
}

func TestMathClz32(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"math_clz32(0)", 32},
		{"math_clz32(1)", 31},
		{"math_clz32(2)", 30},
		{"math_clz32(4)", 29},
		{"math_clz32(8)", 28},
	}

	for _, tt := range tests {
		result := evalMathInput(tt.input)
		num, ok := result.(*object.Number)
		if !ok {
			t.Errorf("Expected Number for %s, got %T", tt.input, result)
			continue
		}

		if num.Value != tt.expected {
			t.Errorf("Expected %s = %f, got %f", tt.input, tt.expected, num.Value)
		}
	}
}

func TestMathFround(t *testing.T) {
	input := `math_fround(1.337)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	expected := float64(float32(1.337))
	if !floatEquals(num.Value, expected, 1e-10) {
		t.Errorf("Expected fround(1.337) = %f, got %f", expected, num.Value)
	}
}

func TestMathHypot(t *testing.T) {
	input := `math_hypot(3, 4)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if !floatEquals(num.Value, 5.0, 1e-10) {
		t.Errorf("Expected hypot(3, 4) = 5.0, got %f", num.Value)
	}
}

func TestMathHypotMultiple(t *testing.T) {
	input := `math_hypot(1, 2, 2)`
	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	expected := 3.0 // sqrt(1 + 4 + 4) = sqrt(9) = 3
	if !floatEquals(num.Value, expected, 1e-10) {
		t.Errorf("Expected hypot(1, 2, 2) = 3.0, got %f", num.Value)
	}
}

// Test Real-World Use Cases

func TestCircleCalculations(t *testing.T) {
	input := `
		// Calculate circle circumference and area
		dhoro radius = 5;
		dhoro circumference = 2 * MATH_PI * radius;
		dhoro area = MATH_PI * radius * radius;
		[circumference, area]
	`

	result := evalMathInput(input)
	arr, ok := result.(*object.Array)
	if !ok {
		t.Fatalf("Expected Array, got %T", result)
	}

	if len(arr.Elements) != 2 {
		t.Fatalf("Expected 2 elements, got %d", len(arr.Elements))
	}

	circumference := arr.Elements[0].(*object.Number).Value
	area := arr.Elements[1].(*object.Number).Value

	expectedCirc := 2 * math.Pi * 5
	expectedArea := math.Pi * 25

	if !floatEquals(circumference, expectedCirc, 1e-10) {
		t.Errorf("Expected circumference %f, got %f", expectedCirc, circumference)
	}

	if !floatEquals(area, expectedArea, 1e-10) {
		t.Errorf("Expected area %f, got %f", expectedArea, area)
	}
}

func TestDistanceCalculation(t *testing.T) {
	input := `
		// Calculate distance between two points using Pythagorean theorem
		dhoro x1 = 0;
		dhoro y1 = 0;
		dhoro x2 = 3;
		dhoro y2 = 4;
		dhoro distance = math_hypot(x2 - x1, y2 - y1);
		distance
	`

	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	if !floatEquals(num.Value, 5.0, 1e-10) {
		t.Errorf("Expected distance 5.0, got %f", num.Value)
	}
}

func TestExponentialGrowth(t *testing.T) {
	input := `
		// Calculate exponential growth
		dhoro initial = 100;
		dhoro rate = 0.05;
		dhoro time = 10;
		dhoro final = initial * math_exp(rate * time);
		final
	`

	result := evalMathInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T", result)
	}

	expected := 100 * math.Exp(0.05*10)
	if !floatEquals(num.Value, expected, 1e-8) {
		t.Errorf("Expected exponential growth %f, got %f", expected, num.Value)
	}
}
