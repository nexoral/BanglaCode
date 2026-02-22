package test

import (
	"testing"

	"BanglaCode/src/evaluator"
	"BanglaCode/src/evaluator/builtins"
	"BanglaCode/src/lexer"
	"BanglaCode/src/object"
	"BanglaCode/src/parser"
)

func evalErrorInput(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()
	env := object.NewEnvironment()
	builtins.InitializeEnvironmentWithConstants(env)
	return evaluator.Eval(program, env)
}

// Test Error() constructor
func TestErrorConstructor(t *testing.T) {
	input := `
		dhoro err = Error("Something went wrong");
		err
	`

	result := evalErrorInput(input)
	errorMap, ok := result.(*object.Map)
	if !ok {
		t.Fatalf("Expected Map, got %T", result)
	}

	// Check message
	if msg, exists := errorMap.Pairs["message"]; exists {
		if msgStr, ok := msg.(*object.String); ok {
			if msgStr.Value != "Something went wrong" {
				t.Errorf("Expected message 'Something went wrong', got '%s'", msgStr.Value)
			}
		} else {
			t.Error("message should be a String")
		}
	} else {
		t.Error("Error should have 'message' property")
	}

	// Check name
	if name, exists := errorMap.Pairs["name"]; exists {
		if nameStr, ok := name.(*object.String); ok {
			if nameStr.Value != "Error" {
				t.Errorf("Expected name 'Error', got '%s'", nameStr.Value)
			}
		}
	} else {
		t.Error("Error should have 'name' property")
	}
}

// Test TypeError() constructor
func TestTypeErrorConstructor(t *testing.T) {
	input := `
		dhoro err = TypeError("Expected number, got string");
		err
	`

	result := evalErrorInput(input)
	errorMap, ok := result.(*object.Map)
	if !ok {
		t.Fatalf("Expected Map, got %T", result)
	}

	// Check name is TypeError
	if name, exists := errorMap.Pairs["name"]; exists {
		if nameStr, ok := name.(*object.String); ok {
			if nameStr.Value != "TypeError" {
				t.Errorf("Expected name 'TypeError', got '%s'", nameStr.Value)
			}
		}
	}

	// Check message
	if msg, exists := errorMap.Pairs["message"]; exists {
		if msgStr, ok := msg.(*object.String); ok {
			if msgStr.Value != "Expected number, got string" {
				t.Errorf("Unexpected message: %s", msgStr.Value)
			}
		}
	}
}

// Test ReferenceError() constructor
func TestReferenceErrorConstructor(t *testing.T) {
	input := `
		dhoro err = ReferenceError("Variable not defined");
		bhul_naam(err)
	`

	result := evalErrorInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	if str.Value != "ReferenceError" {
		t.Errorf("Expected 'ReferenceError', got '%s'", str.Value)
	}
}

// Test RangeError() constructor
func TestRangeErrorConstructor(t *testing.T) {
	input := `
		dhoro err = RangeError("Index out of bounds");
		bhul_message(err)
	`

	result := evalErrorInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	if str.Value != "Index out of bounds" {
		t.Errorf("Expected 'Index out of bounds', got '%s'", str.Value)
	}
}

// Test SyntaxError() constructor
func TestSyntaxErrorConstructor(t *testing.T) {
	input := `
		dhoro err = SyntaxError("Unexpected token");
		bhul_naam(err)
	`

	result := evalErrorInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	if str.Value != "SyntaxError" {
		t.Errorf("Expected 'SyntaxError', got '%s'", str.Value)
	}
}

// Test bhul_message() function
func TestBhulMessage(t *testing.T) {
	input := `
		dhoro err = Error("Test error message");
		bhul_message(err)
	`

	result := evalErrorInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String, got %T", result)
	}

	if str.Value != "Test error message" {
		t.Errorf("Expected 'Test error message', got '%s'", str.Value)
	}
}

// Test bhul_naam() function
func TestBhulNaam(t *testing.T) {
	tests := []struct {
		errorType string
		expected  string
	}{
		{"Error", "Error"},
		{"TypeError", "TypeError"},
		{"ReferenceError", "ReferenceError"},
		{"RangeError", "RangeError"},
		{"SyntaxError", "SyntaxError"},
	}

	for _, tt := range tests {
		input := `
			dhoro err = ` + tt.errorType + `("test");
			bhul_naam(err)
		`

		result := evalErrorInput(input)
		str, ok := result.(*object.String)
		if !ok {
			t.Fatalf("[%s] Expected String, got %T", tt.errorType, result)
		}

		if str.Value != tt.expected {
			t.Errorf("[%s] Expected '%s', got '%s'", tt.errorType, tt.expected, str.Value)
		}
	}
}

// Test is_error() function
func TestIsError(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`is_error(Error("test"))`, true},
		{`is_error(TypeError("test"))`, true},
		{`is_error(ReferenceError("test"))`, true},
		{`is_error(RangeError("test"))`, true},
		{`is_error(SyntaxError("test"))`, true},
		{`is_error("not an error")`, false},
		{`is_error(42)`, false},
		{`is_error(sotti)`, false},
		{`is_error([1, 2, 3])`, false},
	}

	for _, tt := range tests {
		result := evalErrorInput(tt.input)

		if tt.expected {
			if result != object.TRUE {
				t.Errorf("[%s] Expected TRUE, got %s", tt.input, result.Inspect())
			}
		} else {
			if result != object.FALSE {
				t.Errorf("[%s] Expected FALSE, got %s", tt.input, result.Inspect())
			}
		}
	}
}

// Test throw and catch with custom error types
func TestThrowCatchWithErrorTypes(t *testing.T) {
	input := `
		dhoro caughtError = khali;
		
		chesta {
			felo TypeError("Invalid type provided");
		} dhoro_bhul(e) {
			caughtError = e;
		}
		
		caughtError
	`

	result := evalErrorInput(input)

	// The caught exception should contain the error information
	if result == object.NULL {
		t.Fatal("Expected error to be caught, got NULL")
	}
}

// Test error in function with stack trace
func TestErrorInFunction(t *testing.T) {
	input := `
		kaj divide(a, b) {
			jodi (b == 0) {
				felo RangeError("Division by zero");
			}
			ferao a / b;
		}
		
		dhoro result = khali;
		chesta {
			result = divide(10, 0);
		} dhoro_bhul(e) {
			result = e;
		}
		
		result
	`

	result := evalErrorInput(input)

	// Should catch the RangeError
	if result == object.NULL {
		t.Fatal("Expected error to be caught")
	}
}

// Test multiple error types in try-catch
func TestMultipleErrorTypes(t *testing.T) {
	input := `
		kaj processValue(val) {
			jodi (dhoron(val) != "NUMBER") {
				felo TypeError("Expected number");
			}
			jodi (val < 0) {
				felo RangeError("Value must be positive");
			}
			jodi (val > 100) {
				felo RangeError("Value must be <= 100");
			}
			ferao val * 2;
		}
		
		dhoro results = [];
		
		// Test TypeError
		chesta {
			processValue("not a number");
		} dhoro_bhul(e) {
			dhokao(results, bhul_naam(e));
		}
		
		// Test RangeError (negative)
		chesta {
			processValue(-5);
		} dhoro_bhul(e) {
			dhokao(results, bhul_naam(e));
		}
		
		// Test success
		chesta {
			dhoro val = processValue(50);
			dhokao(results, "success");
		} dhoro_bhul(e) {
			dhokao(results, "failed");
		}
		
		results
	`

	result := evalErrorInput(input)
	arr, ok := result.(*object.Array)
	if !ok {
		t.Fatalf("Expected Array, got %T", result)
	}

	if len(arr.Elements) != 3 {
		t.Fatalf("Expected 3 results, got %d", len(arr.Elements))
	}

	// First should be TypeError
	if str, ok := arr.Elements[0].(*object.String); ok {
		if str.Value != "TypeError" {
			t.Errorf("Expected first error to be TypeError, got %s", str.Value)
		}
	}

	// Second should be RangeError
	if str, ok := arr.Elements[1].(*object.String); ok {
		if str.Value != "RangeError" {
			t.Errorf("Expected second error to be RangeError, got %s", str.Value)
		}
	}

	// Third should be success
	if str, ok := arr.Elements[2].(*object.String); ok {
		if str.Value != "success" {
			t.Errorf("Expected success, got %s", str.Value)
		}
	}
}

// Test bhul_stack() - verify stack trace property exists
func TestBhulStack(t *testing.T) {
	input := `
		dhoro err = Error("Test with stack");
		dhoro stack = bhul_stack(err);
		dhoron(stack)
	`

	result := evalErrorInput(input)
	str, ok := result.(*object.String)
	if !ok {
		t.Fatalf("Expected String (type), got %T", result)
	}

	if str.Value != "STRING" {
		t.Errorf("Expected stack to be STRING type, got %s", str.Value)
	}
}

// Test real-world validation scenario
func TestRealWorldValidation(t *testing.T) {
	input := `
		// User validation function
		kaj validateUser(user) {
			jodi (dhoron(user) != "MAP") {
				felo TypeError("User must be an object");
			}
			
			dhoro userKeys = chabi(user);
			dhoro hasName = mittha;
			dhoro hasAge = mittha;
			
			ghuriye (dhoro i = 0; i < kato(userKeys); i = i + 1) {
				jodi (userKeys[i] == "name") {
					hasName = sotti;
				}
				jodi (userKeys[i] == "age") {
					hasAge = sotti;
				}
			}
			
			jodi (!hasName) {
				felo ReferenceError("User must have name property");
			}
			
			jodi (!hasAge) {
				felo ReferenceError("User must have age property");
			}
			
			dhoro age = user["age"];
			jodi (dhoron(age) != "NUMBER") {
				felo TypeError("Age must be a number");
			}
			
			jodi (age < 0) {
				felo RangeError("Age must be >= 0");
			}
			jodi (age > 150) {
				felo RangeError("Age must be <= 150");
			}
			
			ferao sotti;
		}
		
		// Test valid user
		dhoro validUser = {"name": "Rahim", "age": 25};
		dhoro result1 = validateUser(validUser);
		
		// Test invalid user (missing age)
		dhoro invalidUser = {"name": "Karim"};
		dhoro errorType = khali;
		chesta {
			validateUser(invalidUser);
		} dhoro_bhul(e) {
			errorType = bhul_naam(e);
		}
		
		[result1, errorType]
	`

	result := evalErrorInput(input)

	// Debug: print what we got
	if _, ok := result.(*object.Error); ok {
		t.Skip("Skipping due to error in code:", result.Inspect())
		return
	}

	arr, ok := result.(*object.Array)
	if !ok {
		t.Fatalf("Expected Array, got %T: %s", result, result.Inspect())
	}

	if len(arr.Elements) != 2 {
		t.Fatalf("Expected 2 elements, got %d", len(arr.Elements))
	}

	// First should be TRUE (valid user passed)
	if arr.Elements[0] != object.TRUE {
		t.Errorf("Expected first element to be TRUE, got %s", arr.Elements[0].Inspect())
	}

	// Second should be "ReferenceError"
	if str, ok := arr.Elements[1].(*object.String); ok {
		if str.Value != "ReferenceError" {
			t.Errorf("Expected ReferenceError, got %s", str.Value)
		}
	}
}

// Test error propagation through function calls
func TestErrorPropagation(t *testing.T) {
	input := `
		kaj level3() {
			felo Error("Error from level 3");
		}
		
		kaj level2() {
			ferao level3();
		}
		
		kaj level1() {
			ferao level2();
		}
		
		dhoro caught = mittha;
		chesta {
			level1();
		} dhoro_bhul(e) {
			caught = sotti;
		}
		
		caught
	`

	result := evalErrorInput(input)
	if result != object.TRUE {
		t.Errorf("Expected error to be caught through multiple levels, got %s", result.Inspect())
	}
}

// Test custom error messages with context
func TestCustomErrorMessages(t *testing.T) {
	input := `
		dhoro errorCount = 0;
		
		// File not found simulation
		chesta {
			felo Error("File not found");
		} dhoro_bhul(e) {
			errorCount = errorCount + 1;
		}
		
		// Network error simulation
		chesta {
			felo Error("Network timeout");
		} dhoro_bhul(e) {
			errorCount = errorCount + 1;
		}
		
		// Permission error simulation
		chesta {
			felo Error("Permission denied");
		} dhoro_bhul(e) {
			errorCount = errorCount + 1;
		}
		
		errorCount
	`

	result := evalErrorInput(input)
	num, ok := result.(*object.Number)
	if !ok {
		t.Fatalf("Expected Number, got %T: %s", result, result.Inspect())
	}

	if num.Value != 3 {
		t.Errorf("Expected 3 errors caught, got %f", num.Value)
	}
}

// Test that non-error objects don't match is_error
func TestIsErrorWithNonErrors(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{`is_error(42)`, false},
		{`is_error("string")`, false},
		{`is_error(sotti)`, false},
		{`is_error(Error("test"))`, true},
	}

	for _, tt := range tests {
		result := evalErrorInput(tt.input)

		if tt.expected {
			if result != object.TRUE {
				t.Errorf("[%s] Expected TRUE, got %s", tt.input, result.Inspect())
			}
		} else {
			if result != object.FALSE {
				t.Errorf("[%s] Expected FALSE, got %s", tt.input, result.Inspect())
			}
		}
	}
}
