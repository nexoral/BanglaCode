package builtins

import (
	"BanglaCode/src/object"
	"math"
	"net/url"
	"strings"
)

func init() {
	registerParseInt()
	registerParseFloat()
	registerIsNaN()
	registerIsFinite()
	registerEncodeURI()
	registerDecodeURI()
	registerEncodeURIComponent()
	registerDecodeURIComponent()
}

func registerParseInt() {
	Builtins["purno_sonkhya"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) < 1 || len(args) > 2 {
			return newError("wrong number of arguments. got=%d, want=1 or 2", len(args))
		}
		if args[0].Type() != object.STRING_OBJ && args[0].Type() != object.NUMBER_OBJ {
			return nanNumber()
		}
		input := args[0].Inspect()
		radix := 0
		if len(args) == 2 {
			if args[1].Type() != object.NUMBER_OBJ {
				return nanNumber()
			}
			radix = int(args[1].(*object.Number).Value)
		}
		if v, ok := parseIntLikeJS(input, radix); ok {
			return &object.Number{Value: v}
		}
		return nanNumber()
	}}
}

func registerParseFloat() {
	Builtins["doshomik_sonkhya"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1", len(args))
		}
		if args[0].Type() == object.NUMBER_OBJ {
			return args[0]
		}
		if args[0].Type() != object.STRING_OBJ {
			return nanNumber()
		}
		if v, ok := parseLeadingFloat(args[0].(*object.String).Value); ok {
			return &object.Number{Value: v}
		}
		return nanNumber()
	}}
}

func registerIsNaN() {
	Builtins["sonkhya_na"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1", len(args))
		}
		n := coerceNumber(args[0])
		return object.NativeBoolToBooleanObject(math.IsNaN(n))
	}}
}

func registerIsFinite() {
	Builtins["sonkhya_shimito"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1", len(args))
		}
		n := coerceNumber(args[0])
		return object.NativeBoolToBooleanObject(!math.IsNaN(n) && !math.IsInf(n, 0))
	}}
}

func registerEncodeURI() {
	Builtins["uri_encode"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		s, err := requireSingleString("uri_encode", args)
		if err != nil {
			return err
		}
		return &object.String{Value: encodeURI(s)}
	}}
}

func registerDecodeURI() {
	Builtins["uri_decode"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		s, err := requireSingleString("uri_decode", args)
		if err != nil {
			return err
		}
		decoded, decErr := url.PathUnescape(s)
		if decErr != nil {
			return newError("invalid URI format")
		}
		return &object.String{Value: decoded}
	}}
}

func registerEncodeURIComponent() {
	Builtins["uri_ongsho_encode"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		s, err := requireSingleString("uri_ongsho_encode", args)
		if err != nil {
			return err
		}
		encoded := url.QueryEscape(s)
		encoded = strings.ReplaceAll(encoded, "+", "%20")
		return &object.String{Value: encoded}
	}}
}

func registerDecodeURIComponent() {
	Builtins["uri_ongsho_decode"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		s, err := requireSingleString("uri_ongsho_decode", args)
		if err != nil {
			return err
		}
		decoded, decErr := url.QueryUnescape(s)
		if decErr != nil {
			return newError("invalid URI component format")
		}
		return &object.String{Value: decoded}
	}}
}

func requireSingleString(name string, args []object.Object) (string, *object.Error) {
	if len(args) != 1 {
		return "", newError("wrong number of arguments. got=%d, want=1", len(args))
	}
	if args[0].Type() != object.STRING_OBJ {
		return "", newError("argument to `%s` must be STRING, got %s", name, args[0].Type())
	}
	return args[0].(*object.String).Value, nil
}

func coerceNumber(obj object.Object) float64 {
	switch v := obj.(type) {
	case *object.Number:
		return v.Value
	case *object.Boolean:
		if v.Value {
			return 1
		}
		return 0
	case *object.Null:
		return 0
	case *object.String:
		if parsed, ok := parseLeadingFloat(v.Value); ok {
			return parsed
		}
		return math.NaN()
	default:
		return math.NaN()
	}
}

func encodeURI(s string) string {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		c := s[i]
		if isAllowedInURI(c) {
			b.WriteByte(c)
			continue
		}
		b.WriteString("%")
		b.WriteString(strings.ToUpper(hexByte(c)))
	}
	return b.String()
}

func isAllowedInURI(c byte) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
		return true
	}
	switch c {
	case ';', ',', '/', '?', ':', '@', '&', '=', '+', '$', '#', '-', '_', '.', '!', '~', '*', '\'', '(', ')':
		return true
	default:
		return false
	}
}

func hexByte(c byte) string {
	const digits = "0123456789ABCDEF"
	return string([]byte{digits[c>>4], digits[c&0x0F]})
}
