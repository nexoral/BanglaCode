package builtins

import (
	"BanglaCode/src/object"
	"strings"
)

func init() {
	// Length - dorghyo (দৈর্ঘ্য - length)
	Builtins["dorghyo"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Number{Value: float64(len(arg.Value))}
			case *object.Array:
				return &object.Number{Value: float64(len(arg.Elements))}
			default:
				return newError("argument to `dorghyo` not supported, got %s", args[0].Type())
			}
		},
	}

	// Upper - boroHater (বড় হাতের - uppercase)
	Builtins["boroHater"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `boroHater` must be STRING, got %s", args[0].Type())
			}
			str := args[0].(*object.String).Value
			return &object.String{Value: strings.ToUpper(str)}
		},
	}

	// Lower - chotoHater (ছোট হাতের - lowercase)
	Builtins["chotoHater"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `chotoHater` must be STRING, got %s", args[0].Type())
			}
			str := args[0].(*object.String).Value
			return &object.String{Value: strings.ToLower(str)}
		},
	}

	// Split - bhag (ভাগ - divide)
	Builtins["bhag"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
				return newError("arguments to `bhag` must be STRINGs")
			}
			str := args[0].(*object.String).Value
			sep := args[1].(*object.String).Value
			parts := strings.Split(str, sep)
			elements := make([]object.Object, len(parts))
			for i, p := range parts {
				elements[i] = &object.String{Value: p}
			}
			return &object.Array{Elements: elements}
		},
	}

	// Join - joro (জোড়ো - join)
	Builtins["joro"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to `joro` must be ARRAY, got %s", args[0].Type())
			}
			if args[1].Type() != object.STRING_OBJ {
				return newError("second argument to `joro` must be STRING, got %s", args[1].Type())
			}
			arr := args[0].(*object.Array)
			sep := args[1].(*object.String).Value
			parts := make([]string, len(arr.Elements))
			for i, el := range arr.Elements {
				parts[i] = el.Inspect()
			}
			return &object.String{Value: strings.Join(parts, sep)}
		},
	}

	// Trim - chhanto (ছাঁটো - trim)
	Builtins["chhanto"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("argument to `chhanto` must be STRING, got %s", args[0].Type())
			}
			str := args[0].(*object.String).Value
			return &object.String{Value: strings.TrimSpace(str)}
		},
	}

	// Index of - khojo (খোঁজো - search)
	Builtins["khojo"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ {
				return newError("arguments to `khojo` must be STRINGs")
			}
			str := args[0].(*object.String).Value
			substr := args[1].(*object.String).Value
			return &object.Number{Value: float64(strings.Index(str, substr))}
		},
	}

	// Substring - angsho (অংশ - portion)
	Builtins["angsho"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 || len(args) > 3 {
				return newError("wrong number of arguments. got=%d, want=2 or 3", len(args))
			}
			if args[0].Type() != object.STRING_OBJ {
				return newError("first argument to `angsho` must be STRING, got %s", args[0].Type())
			}
			if args[1].Type() != object.NUMBER_OBJ {
				return newError("second argument to `angsho` must be NUMBER, got %s", args[1].Type())
			}
			str := args[0].(*object.String).Value
			start := int(args[1].(*object.Number).Value)
			end := len(str)
			if len(args) == 3 {
				if args[2].Type() != object.NUMBER_OBJ {
					return newError("third argument to `angsho` must be NUMBER, got %s", args[2].Type())
				}
				end = int(args[2].(*object.Number).Value)
			}
			if start < 0 {
				start = 0
			}
			if end > len(str) {
				end = len(str)
			}
			if start > end {
				return &object.String{Value: ""}
			}
			return &object.String{Value: str[start:end]}
		},
	}

	// Replace - bodlo (বদলো - change)
	Builtins["bodlo"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 3 {
				return newError("wrong number of arguments. got=%d, want=3", len(args))
			}
			if args[0].Type() != object.STRING_OBJ || args[1].Type() != object.STRING_OBJ || args[2].Type() != object.STRING_OBJ {
				return newError("all arguments to `bodlo` must be STRINGs")
			}
			str := args[0].(*object.String).Value
			old := args[1].(*object.String).Value
			new := args[2].(*object.String).Value
			return &object.String{Value: strings.ReplaceAll(str, old, new)}
		},
	}
}
