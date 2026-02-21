package builtins

import (
	"BanglaCode/src/object"
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

// newError creates a new error object without position info
func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

// objectsEqual checks if two objects are equal
func objectsEqual(left, right object.Object) bool {
	if left.Type() != right.Type() {
		return false
	}

	switch l := left.(type) {
	case *object.Number:
		return l.Value == right.(*object.Number).Value
	case *object.String:
		return l.Value == right.(*object.String).Value
	case *object.Boolean:
		return l.Value == right.(*object.Boolean).Value
	case *object.Null:
		return true
	default:
		return left == right
	}
}

// isTruthy matches evaluator truthiness rules for builtin callbacks
func isTruthy(obj object.Object) bool {
	if obj == nil {
		return false
	}
	if obj == object.NULL || obj == object.FALSE {
		return false
	}
	if b, ok := obj.(*object.Boolean); ok {
		return b.Value
	}
	return true
}

// parseLeadingFloat parses leading numeric content similar to JS parseFloat
func parseLeadingFloat(input string) (float64, bool) {
	s := strings.TrimSpace(input)
	if s == "" {
		return 0, false
	}

	maxLen := 0
	for i := 1; i <= len(s); i++ {
		if _, err := strconv.ParseFloat(s[:i], 64); err == nil {
			maxLen = i
		}
	}
	if maxLen == 0 {
		return 0, false
	}
	v, err := strconv.ParseFloat(s[:maxLen], 64)
	if err != nil {
		return 0, false
	}
	return v, true
}

// parseIntLikeJS parses input like JS parseInt with optional radix
func parseIntLikeJS(input string, radix int) (float64, bool) {
	s := strings.TrimSpace(input)
	if s == "" {
		return 0, false
	}

	sign := 1.0
	if s[0] == '+' || s[0] == '-' {
		if s[0] == '-' {
			sign = -1
		}
		s = s[1:]
		if s == "" {
			return 0, false
		}
	}

	base := radix
	if base == 0 {
		base = 10
		if len(s) >= 2 && (s[:2] == "0x" || s[:2] == "0X") {
			base = 16
			s = s[2:]
		}
	}
	if base < 2 || base > 36 {
		return 0, false
	}

	if base == 16 && len(s) >= 2 && (s[:2] == "0x" || s[:2] == "0X") {
		s = s[2:]
	}

	i := 0
	for i < len(s) && isValidDigitForBase(rune(s[i]), base) {
		i++
	}
	if i == 0 {
		return 0, false
	}

	val, err := strconv.ParseInt(s[:i], base, 64)
	if err != nil {
		return 0, false
	}
	return sign * float64(val), true
}

func isValidDigitForBase(ch rune, base int) bool {
	if unicode.IsDigit(ch) {
		return int(ch-'0') < base
	}
	if ch >= 'a' && ch <= 'z' {
		return int(ch-'a')+10 < base
	}
	if ch >= 'A' && ch <= 'Z' {
		return int(ch-'A')+10 < base
	}
	return false
}

func nanNumber() *object.Number {
	return &object.Number{Value: math.NaN()}
}

func mapKeyFromObject(key object.Object) string {
	switch k := key.(type) {
	case *object.String:
		return k.Value
	case *object.Number:
		return k.Inspect()
	default:
		return ""
	}
}
