package builtins

import (
	"BanglaCode/src/object"
	"sort"
)

func init() {
	// Push - dhokao (ঢোকাও - insert)
	Builtins["dhokao"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `dhokao` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			arr.Elements = append(arr.Elements, args[1])
			return arr
		},
	}

	// Pop - berKoro (বের করো - take out)
	Builtins["berKoro"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `berKoro` must be ARRAY, got %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				lastElement := arr.Elements[length-1]
				arr.Elements = arr.Elements[:length-1]
				return lastElement
			}
			return object.NULL
		},
	}

	// Keys - chabi (চাবি - keys)
	Builtins["chabi"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.MAP_OBJ {
				return newError("argument to `chabi` must be MAP, got %s", args[0].Type())
			}

			mapObj := args[0].(*object.Map)
			keys := make([]object.Object, 0, len(mapObj.Pairs))
			for key := range mapObj.Pairs {
				keys = append(keys, &object.String{Value: key})
			}
			return &object.Array{Elements: keys}
		},
	}

	// Slice - kato (কাটো - cut)
	Builtins["kato"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 2 || len(args) > 3 {
				return newError("wrong number of arguments. got=%d, want=2 or 3", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to `kato` must be ARRAY, got %s", args[0].Type())
			}
			if args[1].Type() != object.NUMBER_OBJ {
				return newError("second argument to `kato` must be NUMBER, got %s", args[1].Type())
			}
			arr := args[0].(*object.Array)
			start := int(args[1].(*object.Number).Value)
			end := len(arr.Elements)
			if len(args) == 3 {
				if args[2].Type() != object.NUMBER_OBJ {
					return newError("third argument to `kato` must be NUMBER, got %s", args[2].Type())
				}
				end = int(args[2].(*object.Number).Value)
			}
			if start < 0 {
				start = 0
			}
			if end > len(arr.Elements) {
				end = len(arr.Elements)
			}
			if start > end {
				return &object.Array{Elements: []object.Object{}}
			}
			newElements := make([]object.Object, end-start)
			copy(newElements, arr.Elements[start:end])
			return &object.Array{Elements: newElements}
		},
	}

	// Reverse - ulto (উল্টো - reverse)
	Builtins["ulto"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `ulto` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			newElements := make([]object.Object, length)
			for i := 0; i < length; i++ {
				newElements[i] = arr.Elements[length-1-i]
			}
			return &object.Array{Elements: newElements}
		},
	}

	// Includes - ache (আছে - exists)
	Builtins["ache"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("first argument to `ache` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			target := args[1]
			for _, el := range arr.Elements {
				if objectsEqual(el, target) {
					return object.TRUE
				}
			}
			return object.FALSE
		},
	}

	// Sort - saja (সাজা - arrange)
	Builtins["saja"] = &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `saja` must be ARRAY, got %s", args[0].Type())
			}
			arr := args[0].(*object.Array)
			newElements := make([]object.Object, len(arr.Elements))
			copy(newElements, arr.Elements)
			sort.Slice(newElements, func(i, j int) bool {
				if newElements[i].Type() == object.NUMBER_OBJ && newElements[j].Type() == object.NUMBER_OBJ {
					return newElements[i].(*object.Number).Value < newElements[j].(*object.Number).Value
				}
				return newElements[i].Inspect() < newElements[j].Inspect()
			})
			return &object.Array{Elements: newElements}
		},
	}
}
