package test

import (
	"testing"
)

// TestRestParameter tests the rest parameter functionality
func TestRestParameter(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		// Basic rest parameter
		{
			`kaj sum(...numbers) {
				dhoro total = 0;
				ghuriye (dhoro i = 0; i < dorghyo(numbers); i = i + 1) {
					total = total + numbers[i];
				}
				ferao total;
			}
			sum(1, 2, 3);`,
			6.0,
		},
		// Rest parameter with no arguments
		{
			`kaj collect(...items) {
				ferao dorghyo(items);
			}
			collect();`,
			0.0,
		},
		// Rest parameter with many arguments
		{
			`kaj collect(...items) {
				ferao dorghyo(items);
			}
			collect(1, 2, 3, 4, 5, 6, 7, 8, 9, 10);`,
			10.0,
		},
		// Mixed regular and rest parameters
		{
			`kaj greet(greeting, ...names) {
				ferao dorghyo(names);
			}
			greet("Hello", "Alice", "Bob", "Charlie");`,
			3.0,
		},
		// Rest parameter is always an array
		{
			`kaj getType(...args) {
				ferao dhoron(args);
			}
			getType(1, 2, 3);`,
			"ARRAY",
		},
		// Access rest parameter elements
		{
			`kaj first(...args) {
				jodi (dorghyo(args) > 0) {
					ferao args[0];
				}
				ferao khali;
			}
			first(42, 100, 200);`,
			42.0,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)

		switch expected := tt.expected.(type) {
		case float64:
			testNumberObject(t, evaluated, expected)
		case string:
			testStringObject(t, evaluated, expected)
		}
	}
}

// TestSpreadInFunctionCall tests spread operator in function calls
func TestSpreadInFunctionCall(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
	}{
		// Spread array as arguments
		{
			`kaj sum(...numbers) {
				dhoro total = 0;
				ghuriye (dhoro i = 0; i < dorghyo(numbers); i = i + 1) {
					total = total + numbers[i];
				}
				ferao total;
			}
			dhoro nums = [1, 2, 3, 4, 5];
			sum(...nums);`,
			15.0,
		},
		// Spread with other arguments
		{
			`kaj sum(...numbers) {
				dhoro total = 0;
				ghuriye (dhoro i = 0; i < dorghyo(numbers); i = i + 1) {
					total = total + numbers[i];
				}
				ferao total;
			}
			dhoro nums = [2, 3];
			sum(1, ...nums, 4, 5);`,
			15.0,
		},
		// Multiple spreads
		{
			`kaj sum(...numbers) {
				dhoro total = 0;
				ghuriye (dhoro i = 0; i < dorghyo(numbers); i = i + 1) {
					total = total + numbers[i];
				}
				ferao total;
			}
			dhoro a = [1, 2];
			dhoro b = [3, 4];
			sum(...a, ...b);`,
			10.0,
		},
		// Spread empty array
		{
			`kaj sum(...numbers) {
				dhoro total = 0;
				ghuriye (dhoro i = 0; i < dorghyo(numbers); i = i + 1) {
					total = total + numbers[i];
				}
				ferao total;
			}
			dhoro empty = [];
			sum(...empty);`,
			0.0,
		},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testNumberObject(t, evaluated, tt.expected)
	}
}

// TestSpreadInArrayLiteral tests spread operator in array literals
func TestSpreadInArrayLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected []float64
	}{
		// Basic spread in array
		{
			`dhoro arr = [1, 2];
			dhoro result = [...arr];
			result;`,
			[]float64{1, 2},
		},
		// Combine two arrays
		{
			`dhoro a = [1, 2];
			dhoro b = [3, 4];
			dhoro result = [...a, ...b];
			result;`,
			[]float64{1, 2, 3, 4},
		},
		// Spread with additional elements
		{
			`dhoro arr = [2, 3];
			dhoro result = [1, ...arr, 4, 5];
			result;`,
			[]float64{1, 2, 3, 4, 5},
		},
		// Multiple spreads with elements
		{
			`dhoro a = [1];
			dhoro b = [3];
			dhoro result = [...a, 2, ...b, 4];
			result;`,
			[]float64{1, 2, 3, 4},
		},
		// Spread empty array
		{
			`dhoro empty = [];
			dhoro result = [1, ...empty, 2];
			result;`,
			[]float64{1, 2},
		},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		testArrayObject(t, evaluated, tt.expected, i)
	}
}

// TestSpreadErrors tests error cases for spread operator
func TestSpreadErrors(t *testing.T) {
	tests := []struct {
		input         string
		expectedError string
	}{
		// Spread non-array
		{
			`dhoro x = 5;
			dhoro result = [...x];`,
			"spread operator requires an array",
		},
		// Spread string (not allowed)
		{
			`dhoro s = "hello";
			dhoro result = [...s];`,
			"spread operator requires an array",
		},
	}

	for i, tt := range tests {
		evaluated := testEval(tt.input)
		testErrorObject(t, evaluated, tt.expectedError, i)
	}
}
