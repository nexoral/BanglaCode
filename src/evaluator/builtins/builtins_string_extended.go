package builtins

import (
	"BanglaCode/src/object"
	"strings"
	"unicode/utf8"
)

func init() {
	registerStringIncludes()
	registerStringStartsWith()
	registerStringEndsWith()
	registerStringRepeat()
	registerStringPad()
	registerStringAt()
	registerStringCharCodeAt()
	registerStringTrim()
	registerStringLastIndexOf()
	registerStringCodePointAt()
	registerStringLocaleCompare()
	registerStringNormalize()
}

func registerStringIncludes() {
	Builtins["ache_text"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		str, part, err := requireTwoStrings("ache_text", args)
		if err != nil {
			return err
		}
		return object.NativeBoolToBooleanObject(strings.Contains(str, part))
	}}
}

func registerStringStartsWith() {
	Builtins["shuru_diye"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		str, prefix, err := requireTwoStrings("shuru_diye", args)
		if err != nil {
			return err
		}
		return object.NativeBoolToBooleanObject(strings.HasPrefix(str, prefix))
	}}
}

func registerStringEndsWith() {
	Builtins["shesh_diye"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		str, suffix, err := requireTwoStrings("shesh_diye", args)
		if err != nil {
			return err
		}
		return object.NativeBoolToBooleanObject(strings.HasSuffix(str, suffix))
	}}
}

func registerStringRepeat() {
	Builtins["baro"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=2", len(args))
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("first argument to `baro` must be STRING, got %s", args[0].Type())
		}
		if args[1].Type() != object.NUMBER_OBJ {
			return newError("second argument to `baro` must be NUMBER, got %s", args[1].Type())
		}
		n := int(args[1].(*object.Number).Value)
		if n < 0 {
			return newError("repeat count must be >= 0")
		}
		return &object.String{Value: strings.Repeat(args[0].(*object.String).Value, n)}
	}}
}

func registerStringPad() {
	Builtins["agey_bhoro"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		return padString("agey_bhoro", true, args)
	}}
	Builtins["pichoney_bhoro"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		return padString("pichoney_bhoro", false, args)
	}}
}

func registerStringAt() {
	Builtins["okkhor"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		return textAt("okkhor", false, args)
	}}
	Builtins["text_at"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		return textAt("text_at", true, args)
	}}
}

func registerStringCharCodeAt() {
	Builtins["okkhor_code"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=2", len(args))
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("first argument to `okkhor_code` must be STRING, got %s", args[0].Type())
		}
		if args[1].Type() != object.NUMBER_OBJ {
			return newError("second argument to `okkhor_code` must be NUMBER, got %s", args[1].Type())
		}
		runes := []rune(args[0].(*object.String).Value)
		idx := int(args[1].(*object.Number).Value)
		if idx < 0 || idx >= len(runes) {
			return nanNumber()
		}
		return &object.Number{Value: float64(runes[idx])}
	}}
}

func registerStringTrim() {
	Builtins["chhanto_shuru"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1", len(args))
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("argument to `chhanto_shuru` must be STRING, got %s", args[0].Type())
		}
		return &object.String{Value: strings.TrimLeftFunc(args[0].(*object.String).Value, unicodeSpace)}
	}}

	Builtins["chhanto_shesh"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1", len(args))
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("argument to `chhanto_shesh` must be STRING, got %s", args[0].Type())
		}
		return &object.String{Value: strings.TrimRightFunc(args[0].(*object.String).Value, unicodeSpace)}
	}}
}

func registerStringLastIndexOf() {
	Builtins["shesh_khojo"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		str, part, err := requireTwoStrings("shesh_khojo", args)
		if err != nil {
			return err
		}
		return &object.Number{Value: float64(strings.LastIndex(str, part))}
	}}
}

func registerStringCodePointAt() {
	Builtins["codepoint_at"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=2", len(args))
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("first argument to `codepoint_at` must be STRING, got %s", args[0].Type())
		}
		if args[1].Type() != object.NUMBER_OBJ {
			return newError("second argument to `codepoint_at` must be NUMBER, got %s", args[1].Type())
		}
		runes := []rune(args[0].(*object.String).Value)
		idx := int(args[1].(*object.Number).Value)
		if idx < 0 || idx >= len(runes) {
			return nanNumber()
		}
		return &object.Number{Value: float64(runes[idx])}
	}}
}

func registerStringLocaleCompare() {
	Builtins["tulona_text"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		left, right, err := requireTwoStrings("tulona_text", args)
		if err != nil {
			return err
		}
		if left < right {
			return &object.Number{Value: -1}
		}
		if left > right {
			return &object.Number{Value: 1}
		}
		return &object.Number{Value: 0}
	}}
}

func registerStringNormalize() {
	Builtins["shadharon_text"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1", len(args))
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("argument to `shadharon_text` must be STRING, got %s", args[0].Type())
		}
		return &object.String{Value: args[0].(*object.String).Value}
	}}
}

func requireTwoStrings(name string, args []object.Object) (string, string, *object.Error) {
	if len(args) != 2 {
		return "", "", newError("wrong number of arguments. got=%d, want=2", len(args))
	}
	if args[0].Type() != object.STRING_OBJ {
		return "", "", newError("first argument to `%s` must be STRING, got %s", name, args[0].Type())
	}
	if args[1].Type() != object.STRING_OBJ {
		return "", "", newError("second argument to `%s` must be STRING, got %s", name, args[1].Type())
	}
	return args[0].(*object.String).Value, args[1].(*object.String).Value, nil
}

func padString(name string, left bool, args []object.Object) object.Object {
	if len(args) < 2 || len(args) > 3 {
		return newError("wrong number of arguments. got=%d, want=2 or 3", len(args))
	}
	if args[0].Type() != object.STRING_OBJ {
		return newError("first argument to `%s` must be STRING, got %s", name, args[0].Type())
	}
	if args[1].Type() != object.NUMBER_OBJ {
		return newError("second argument to `%s` must be NUMBER, got %s", name, args[1].Type())
	}

	str := args[0].(*object.String).Value
	target := int(args[1].(*object.Number).Value)
	if target <= utf8.RuneCountInString(str) {
		return &object.String{Value: str}
	}

	pad := " "
	if len(args) == 3 {
		if args[2].Type() != object.STRING_OBJ {
			return newError("third argument to `%s` must be STRING, got %s", name, args[2].Type())
		}
		pad = args[2].(*object.String).Value
		if pad == "" {
			pad = " "
		}
	}

	need := target - utf8.RuneCountInString(str)
	repeatCount := need/utf8.RuneCountInString(pad) + 2
	filler := []rune(strings.Repeat(pad, repeatCount))
	if len(filler) > need {
		filler = filler[:need]
	}
	if left {
		return &object.String{Value: string(filler) + str}
	}
	return &object.String{Value: str + string(filler)}
}

func textAt(name string, allowNegative bool, args []object.Object) object.Object {
	if len(args) != 2 {
		return newError("wrong number of arguments. got=%d, want=2", len(args))
	}
	if args[0].Type() != object.STRING_OBJ {
		return newError("first argument to `%s` must be STRING, got %s", name, args[0].Type())
	}
	if args[1].Type() != object.NUMBER_OBJ {
		return newError("second argument to `%s` must be NUMBER, got %s", name, args[1].Type())
	}

	runes := []rune(args[0].(*object.String).Value)
	idx := int(args[1].(*object.Number).Value)
	if allowNegative && idx < 0 {
		idx = len(runes) + idx
	}
	if idx < 0 || idx >= len(runes) {
		return &object.String{Value: ""}
	}
	return &object.String{Value: string(runes[idx])}
}

func unicodeSpace(r rune) bool {
	return r == ' ' || r == '\n' || r == '\r' || r == '\t' || r == '\f' || r == '\v'
}
