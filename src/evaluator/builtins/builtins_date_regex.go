package builtins

import (
	"BanglaCode/src/object"
	"regexp"
	"time"
)

func init() {
	registerDateNow()
	registerDateParse()
	registerDateFormat()
	registerRegexTest()
	registerRegexMatch()
	registerRegexMatchAll()
	registerRegexSearch()
	registerRegexReplace()
	registerMatchWrappers()
}

func registerDateNow() {
	Builtins["tarikh_ekhon"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 0 {
			return newError("wrong number of arguments. got=%d, want=0", len(args))
		}
		return &object.Number{Value: float64(time.Now().UnixMilli())}
	}}
}

func registerDateParse() {
	Builtins["tarikh_parse"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1", len(args))
		}
		if args[0].Type() != object.STRING_OBJ {
			return newError("argument to `tarikh_parse` must be STRING, got %s", args[0].Type())
		}
		text := args[0].(*object.String).Value
		layouts := []string{time.RFC3339, time.RFC1123, time.RFC822, "2006-01-02 15:04:05", "2006-01-02"}
		for _, layout := range layouts {
			if parsed, err := time.Parse(layout, text); err == nil {
				return &object.Number{Value: float64(parsed.UnixMilli())}
			}
		}
		return nanNumber()
	}}
}

func registerDateFormat() {
	Builtins["tarikh_format"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) < 1 || len(args) > 2 {
			return newError("wrong number of arguments. got=%d, want=1 or 2", len(args))
		}
		if args[0].Type() != object.NUMBER_OBJ {
			return newError("first argument to `tarikh_format` must be NUMBER, got %s", args[0].Type())
		}
		layout := time.RFC3339
		if len(args) == 2 {
			if args[1].Type() != object.STRING_OBJ {
				return newError("second argument to `tarikh_format` must be STRING, got %s", args[1].Type())
			}
			layout = args[1].(*object.String).Value
		}
		ts := int64(args[0].(*object.Number).Value)
		return &object.String{Value: time.UnixMilli(ts).UTC().Format(layout)}
	}}
}

func registerRegexTest() {
	Builtins["regex_test"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		re, text, errObj := regexPatternText("regex_test", args)
		if errObj != nil {
			return errObj
		}
		return object.NativeBoolToBooleanObject(re.MatchString(text))
	}}
}

func registerRegexMatch() {
	Builtins["regex_match"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		re, text, errObj := regexPatternText("regex_match", args)
		if errObj != nil {
			return errObj
		}
		match := re.FindStringSubmatch(text)
		if match == nil {
			return object.NULL
		}
		out := make([]object.Object, 0, len(match))
		for _, m := range match {
			out = append(out, &object.String{Value: m})
		}
		return &object.Array{Elements: out}
	}}
}

func registerRegexMatchAll() {
	Builtins["regex_match_all"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		re, text, errObj := regexPatternText("regex_match_all", args)
		if errObj != nil {
			return errObj
		}
		matches := re.FindAllString(text, -1)
		out := make([]object.Object, 0, len(matches))
		for _, m := range matches {
			out = append(out, &object.String{Value: m})
		}
		return &object.Array{Elements: out}
	}}
}

func registerRegexSearch() {
	Builtins["regex_search"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		re, text, errObj := regexPatternText("regex_search", args)
		if errObj != nil {
			return errObj
		}
		loc := re.FindStringIndex(text)
		if loc == nil {
			return &object.Number{Value: -1}
		}
		return &object.Number{Value: float64(loc[0])}
	}}
}

func registerRegexReplace() {
	Builtins["regex_replace"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) < 3 || len(args) > 4 {
			return newError("wrong number of arguments. got=%d, want=3 or 4", len(args))
		}
		reArgs := []object.Object{args[0], args[1]}
		if len(args) == 4 {
			reArgs = append(reArgs, args[3])
		}
		re, text, errObj := regexPatternText("regex_replace", reArgs)
		if errObj != nil {
			return errObj
		}
		if args[2].Type() != object.STRING_OBJ {
			return newError("third argument to `regex_replace` must be STRING, got %s", args[2].Type())
		}
		replacement := args[2].(*object.String).Value
		return &object.String{Value: re.ReplaceAllString(text, replacement)}
	}}
}

func registerMatchWrappers() {
	Builtins["match"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		regexArgs, errObj := validateMatchLikeArgs(args)
		if errObj != nil {
			return errObj
		}
		return Builtins["regex_match"].Fn(regexArgs...)
	}}

	Builtins["matchAll"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		regexArgs, errObj := validateMatchLikeArgs(args)
		if errObj != nil {
			return errObj
		}
		return Builtins["regex_match_all"].Fn(regexArgs...)
	}}

	Builtins["search"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		regexArgs, errObj := validateMatchLikeArgs(args)
		if errObj != nil {
			return errObj
		}
		return Builtins["regex_search"].Fn(regexArgs...)
	}}
}

func validateMatchLikeArgs(args []object.Object) ([]object.Object, *object.Error) {
	if len(args) < 2 || len(args) > 3 {
		return nil, newError("wrong number of arguments. got=%d, want=2 or 3", len(args))
	}
	regexArgs := []object.Object{args[1], args[0]}
	if len(args) == 3 {
		regexArgs = append(regexArgs, args[2])
	}
	return regexArgs, nil
}

func regexPatternText(name string, args []object.Object) (*regexp.Regexp, string, *object.Error) {
	if len(args) < 2 || len(args) > 3 {
		return nil, "", newError("wrong number of arguments. got=%d, want=2 or 3", len(args))
	}
	if args[0].Type() != object.STRING_OBJ {
		return nil, "", newError("first argument to `%s` must be STRING pattern, got %s", name, args[0].Type())
	}
	if args[1].Type() != object.STRING_OBJ {
		return nil, "", newError("second argument to `%s` must be STRING text, got %s", name, args[1].Type())
	}
	pattern := args[0].(*object.String).Value
	if len(args) == 3 {
		if args[2].Type() != object.STRING_OBJ {
			return nil, "", newError("third argument to `%s` must be STRING flags, got %s", name, args[2].Type())
		}
		pattern = applyRegexFlags(pattern, args[2].(*object.String).Value)
	}

	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, "", newError("invalid regex pattern: %s", err.Error())
	}
	return re, args[1].(*object.String).Value, nil
}

func applyRegexFlags(pattern, flags string) string {
	withFlags := pattern
	needsPrefix := false
	for _, ch := range flags {
		switch ch {
		case 'i':
			needsPrefix = true
			withFlags = "(?i)" + withFlags
		case 'm':
			needsPrefix = true
			withFlags = "(?m)" + withFlags
		case 's':
			needsPrefix = true
			withFlags = "(?s)" + withFlags
		}
	}
	if !needsPrefix {
		return pattern
	}
	return withFlags
}
