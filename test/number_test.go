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
func evalNumberInput(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	builtins.InitializeEnvironmentWithConstants(env)
	return evaluator.Eval(program, env)
}

// TestNumberConstants tests all Number constants
func TestNumberConstants(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		{"NUMBER_MAX_SAFE_INTEGER", 9007199254740991.0},
		{"NUMBER_MIN_SAFE_INTEGER", -9007199254740991.0},
		{"NUMBER_MAX_VALUE", math.MaxFloat64},
		{"NUMBER_MIN_VALUE", math.SmallestNonzeroFloat64},
		{"NUMBER_POSITIVE_INFINITY", math.Inf(1)},
		{"NUMBER_NEGATIVE_INFINITY", math.Inf(-1)},
		{"NUMBER_EPSILON", 2.220446049250313e-16},
	}

	for _, tt := range tests {
		result := evalNumberInput(tt.input)
		num, ok := result.(*object.Number)
		if !ok {
			t.Errorf("Expected Number for %s, got %T", tt.input, result)
			continue
		}

		if math.IsInf(tt.expected, 1) && !math.IsInf(num.Value, 1) {
			t.Errorf("%s: expected +Infinity, got %f", tt.input, num.Value)
		} else if math.IsInf(tt.expected, -1) && !math.IsInf(num.Value, -1) {
			t.Errorf("%s: expected -Infinity, got %f", tt.input, num.Value)
		} else if !math.IsInf(tt.expected, 0) && num.Value != tt.expected {
			t.Errorf("%s: expected %e, got %e", tt.input, tt.expected, num.Value)
		}
	}
}

// TestNumberNaNConstant tests NUMBER_NAN constant separately
func TestNumberNaNConstant(t *testing.T) {
	result := evalNumberInput("NUMBER_NAN")
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number for NUMBER_NAN, got %T", result)
	}

	if !math.IsNaN(num.Value) {
		t.Errorf("NUMBER_NAN: expected NaN, got %f", num.Value)
	}
}

// TestNumberIsFinite tests sonkhya_sesh (Number.isFinite)
func TestNumberIsFinite(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"sonkhya_sesh(42)", true},
		{"sonkhya_sesh(3.14)", true},
		{"sonkhya_sesh(0)", true},
		{"sonkhya_sesh(-100)", true},
		{"sonkhya_sesh(NUMBER_POSITIVE_INFINITY)", false},
		{"sonkhya_sesh(NUMBER_NEGATIVE_INFINITY)", false},
		{"sonkhya_sesh(NUMBER_NAN)", false},
		{"sonkhya_sesh(NUMBER_MAX_VALUE)", true},
		{"sonkhya_sesh(NUMBER_MIN_VALUE)", true},
	}

	for _, tt := range tests {
		result := evalNumberInput(tt.input)
		boolean, ok := result.(*object.Boolean)
		if !ok {
			t.Errorf("Expected Boolean for %s, got %T", tt.input, result)
			continue
		}

		if boolean.Value != tt.expected {
			t.Errorf("%s: expected %v, got %v", tt.input, tt.expected, boolean.Value)
		}
	}
}

// TestNumberIsFiniteNonNumber tests sonkhya_sesh with non-number arguments
func TestNumberIsFiniteNonNumber(t *testing.T) {
	tests := []string{
		`sonkhya_sesh("hello")`,
		`sonkhya_sesh(sotti)`,
		`sonkhya_sesh(khali)`,
		`sonkhya_sesh([1, 2, 3])`,
	}

	for _, input := range tests {
		result := evalNumberInput(input)
		boolean, ok := result.(*object.Boolean)
		if !ok {
			t.Errorf("Expected Boolean for %s, got %T", input, result)
			continue
		}

		if boolean.Value != false {
			t.Errorf("%s: expected false, got true", input)
		}
	}
}

// TestNumberIsInteger tests sonkhya_purno (Number.isInteger)
func TestNumberIsInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"sonkhya_purno(42)", true},
		{"sonkhya_purno(0)", true},
		{"sonkhya_purno(-100)", true},
		{"sonkhya_purno(3.14)", false},
		{"sonkhya_purno(3.0)", true}, // 3.0 is an integer
		{"sonkhya_purno(NUMBER_POSITIVE_INFINITY)", false},
		{"sonkhya_purno(NUMBER_NEGATIVE_INFINITY)", false},
		{"sonkhya_purno(NUMBER_NAN)", false},
		{"sonkhya_purno(NUMBER_MAX_SAFE_INTEGER)", true},
		{"sonkhya_purno(NUMBER_MIN_SAFE_INTEGER)", true},
	}

	for _, tt := range tests {
		result := evalNumberInput(tt.input)
		boolean, ok := result.(*object.Boolean)
		if !ok {
			t.Errorf("Expected Boolean for %s, got %T", tt.input, result)
			continue
		}

		if boolean.Value != tt.expected {
			t.Errorf("%s: expected %v, got %v", tt.input, tt.expected, boolean.Value)
		}
	}
}

// TestNumberIsNaN tests sonkhya_na_check (Number.isNaN)
func TestNumberIsNaN(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"sonkhya_na_check(NUMBER_NAN)", true},
		{"sonkhya_na_check(42)", false},
		{"sonkhya_na_check(3.14)", false},
		{"sonkhya_na_check(NUMBER_POSITIVE_INFINITY)", false},
		{"sonkhya_na_check(NUMBER_NEGATIVE_INFINITY)", false},
	}

	for _, tt := range tests {
		result := evalNumberInput(tt.input)
		boolean, ok := result.(*object.Boolean)
		if !ok {
			t.Errorf("Expected Boolean for %s, got %T", tt.input, result)
			continue
		}

		if boolean.Value != tt.expected {
			t.Errorf("%s: expected %v, got %v", tt.input, tt.expected, boolean.Value)
		}
	}
}

// TestNumberIsSafeInteger tests sonkhya_nirapod (Number.isSafeInteger)
func TestNumberIsSafeInteger(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"sonkhya_nirapod(42)", true},
		{"sonkhya_nirapod(0)", true},
		{"sonkhya_nirapod(-100)", true},
		{"sonkhya_nirapod(3.14)", false},
		{"sonkhya_nirapod(NUMBER_MAX_SAFE_INTEGER)", true},
		{"sonkhya_nirapod(NUMBER_MIN_SAFE_INTEGER)", true},
		{"sonkhya_nirapod(NUMBER_MAX_SAFE_INTEGER + 1)", false},
		{"sonkhya_nirapod(NUMBER_MIN_SAFE_INTEGER - 1)", false},
		{"sonkhya_nirapod(NUMBER_POSITIVE_INFINITY)", false},
		{"sonkhya_nirapod(NUMBER_NAN)", false},
	}

	for _, tt := range tests {
		result := evalNumberInput(tt.input)
		boolean, ok := result.(*object.Boolean)
		if !ok {
			t.Errorf("Expected Boolean for %s, got %T", tt.input, result)
			continue
		}

		if boolean.Value != tt.expected {
			t.Errorf("%s: expected %v, got %v", tt.input, tt.expected, boolean.Value)
		}
	}
}

// TestNumberSafeIntegerBoundary tests boundary values for safe integers
func TestNumberSafeIntegerBoundary(t *testing.T) {
	input := `
		dhoro max = NUMBER_MAX_SAFE_INTEGER;
		dhoro min = NUMBER_MIN_SAFE_INTEGER;
		dhoro max_safe = sonkhya_nirapod(max);
		dhoro min_safe = sonkhya_nirapod(min);
		dhoro above_max = sonkhya_nirapod(max + 1);
		dhoro below_min = sonkhya_nirapod(min - 1);
		[max_safe, min_safe, above_max, below_min]
	`

	result := evalNumberInput(input)
	arr, ok := result.(*object.Array)
	if !ok {
		t.Fatalf("Expected Array, got %T", result)
	}

	// Check max_safe (index 0)
	if arr.Elements[0].(*object.Boolean).Value != true {
		t.Error("Expected max_safe to be true")
	}

	// Check min_safe (index 1)
	if arr.Elements[1].(*object.Boolean).Value != true {
		t.Error("Expected min_safe to be true")
	}

	// Check above_max (index 2)
	if arr.Elements[2].(*object.Boolean).Value != false {
		t.Error("Expected above_max to be false")
	}

	// Check below_min (index 3)
	if arr.Elements[3].(*object.Boolean).Value != false {
		t.Error("Expected below_min to be false")
	}
}

// TestNumberValidation tests number validation in real code
func TestNumberValidation(t *testing.T) {
	input := `
		kaj validateNumber(value) {
			jodi (sonkhya_na_check(value)) {
				ferao "NaN is not a valid number";
			}
			jodi (!sonkhya_sesh(value)) {
				ferao "Infinity is not allowed";
			}
			jodi (!sonkhya_purno(value)) {
				ferao "Only integers allowed";
			}
			jodi (!sonkhya_nirapod(value)) {
				ferao "Number exceeds safe integer range";
			}
			ferao "Valid number";
		}

		dhoro valid = validateNumber(42);
		dhoro nan = validateNumber(NUMBER_NAN);
		dhoro infinity = validateNumber(NUMBER_POSITIVE_INFINITY);
		dhoro decimal = validateNumber(3.14);
		dhoro unsafe = validateNumber(NUMBER_MAX_SAFE_INTEGER + 2);
		[valid, nan, infinity, decimal, unsafe]
	`

	result := evalNumberInput(input)
	arr, ok := result.(*object.Array)
	if !ok {
		t.Fatalf("Expected Array, got %T", result)
	}

	// Check valid (index 0)
	if arr.Elements[0].(*object.String).Value != "Valid number" {
		t.Errorf("Expected 'Valid number', got %s", arr.Elements[0].(*object.String).Value)
	}

	// Check nan (index 1)
	if arr.Elements[1].(*object.String).Value != "NaN is not a valid number" {
		t.Errorf("Expected 'NaN is not a valid number', got %s", arr.Elements[1].(*object.String).Value)
	}

	// Check infinity (index 2)
	if arr.Elements[2].(*object.String).Value != "Infinity is not allowed" {
		t.Errorf("Expected 'Infinity is not allowed', got %s", arr.Elements[2].(*object.String).Value)
	}

	// Check decimal (index 3)
	if arr.Elements[3].(*object.String).Value != "Only integers allowed" {
		t.Errorf("Expected 'Only integers allowed', got %s", arr.Elements[3].(*object.String).Value)
	}

	// Check unsafe (index 4)
	if arr.Elements[4].(*object.String).Value != "Number exceeds safe integer range" {
		t.Errorf("Expected 'Number exceeds safe integer range', got %s", arr.Elements[4].(*object.String).Value)
	}
}
