package builtins

import "BanglaCode/src/object"

func init() {
	registerArrayConcat()
	registerArrayFlat()
	registerArrayReduceRight()
	registerArrayFind()
	registerArrayFindIndex()
	registerArrayFindLast()
	registerArrayFindLastIndex()
	registerArrayEvery()
	registerArraySome()
	registerArrayFlatMap()
	registerArrayAt()
	registerArrayLastIndexOf()
}

func registerArrayConcat() {
	Builtins["joro_array"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) < 1 {
			return newError("wrong number of arguments. got=%d, want>=1", len(args))
		}
		if args[0].Type() != object.ARRAY_OBJ {
			return newError("first argument to `joro_array` must be ARRAY, got %s", args[0].Type())
		}

		base := args[0].(*object.Array)
		result := make([]object.Object, 0, len(base.Elements))
		result = append(result, base.Elements...)
		for i := 1; i < len(args); i++ {
			if args[i].Type() == object.ARRAY_OBJ {
				result = append(result, args[i].(*object.Array).Elements...)
			} else {
				result = append(result, args[i])
			}
		}
		return &object.Array{Elements: result}
	}}
}

func registerArrayFlat() {
	Builtins["somtol"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) < 1 || len(args) > 2 {
			return newError("wrong number of arguments. got=%d, want=1 or 2", len(args))
		}
		if args[0].Type() != object.ARRAY_OBJ {
			return newError("first argument to `somtol` must be ARRAY, got %s", args[0].Type())
		}

		depth := 1
		if len(args) == 2 {
			if args[1].Type() != object.NUMBER_OBJ {
				return newError("second argument to `somtol` must be NUMBER, got %s", args[1].Type())
			}
			depth = int(args[1].(*object.Number).Value)
			if depth < 0 {
				depth = 0
			}
		}
		return &object.Array{Elements: flattenArray(args[0].(*object.Array).Elements, depth)}
	}}
}

func registerArrayReduceRight() {
	Builtins["sonkuchito_dan"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) < 2 || len(args) > 3 {
			return newError("wrong number of arguments. got=%d, want=2 or 3", len(args))
		}
		if args[0].Type() != object.ARRAY_OBJ {
			return newError("first argument to `sonkuchito_dan` must be ARRAY, got %s", args[0].Type())
		}
		if args[1].Type() != object.FUNCTION_OBJ {
			return newError("second argument to `sonkuchito_dan` must be FUNCTION, got %s", args[1].Type())
		}

		arr := args[0].(*object.Array)
		handler := args[1].(*object.Function)
		if len(arr.Elements) == 0 && len(args) == 2 {
			return newError("reduceRight of empty array with no initial value")
		}

		var accumulator object.Object
		start := len(arr.Elements) - 1
		if len(args) == 3 {
			accumulator = args[2]
		} else {
			accumulator = arr.Elements[start]
			start--
		}
		for i := start; i >= 0; i-- {
			next := EvalFunc(handler, []object.Object{accumulator, arr.Elements[i], &object.Number{Value: float64(i)}, arr})
			if next.Type() == object.ERROR_OBJ {
				return next
			}
			accumulator = next
		}
		return accumulator
	}}
}

func registerArrayFind() {
	Builtins["khojo_prothom"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		arr, callback, err := requireArrayAndCallback("khojo_prothom", args)
		if err != nil {
			return err
		}
		for i, el := range arr.Elements {
			keep := EvalFunc(callback, []object.Object{el, &object.Number{Value: float64(i)}, arr})
			if keep.Type() == object.ERROR_OBJ {
				return keep
			}
			if isTruthy(keep) {
				return el
			}
		}
		return object.NULL
	}}
}

func registerArrayFindIndex() {
	Builtins["khojo_index"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		arr, callback, err := requireArrayAndCallback("khojo_index", args)
		if err != nil {
			return err
		}
		for i, el := range arr.Elements {
			keep := EvalFunc(callback, []object.Object{el, &object.Number{Value: float64(i)}, arr})
			if keep.Type() == object.ERROR_OBJ {
				return keep
			}
			if isTruthy(keep) {
				return &object.Number{Value: float64(i)}
			}
		}
		return &object.Number{Value: -1}
	}}
}

func registerArrayFindLast() {
	Builtins["khojo_shesh"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		arr, callback, err := requireArrayAndCallback("khojo_shesh", args)
		if err != nil {
			return err
		}
		for i := len(arr.Elements) - 1; i >= 0; i-- {
			el := arr.Elements[i]
			keep := EvalFunc(callback, []object.Object{el, &object.Number{Value: float64(i)}, arr})
			if keep.Type() == object.ERROR_OBJ {
				return keep
			}
			if isTruthy(keep) {
				return el
			}
		}
		return object.NULL
	}}
}

func registerArrayFindLastIndex() {
	Builtins["khojo_shesh_index"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		arr, callback, err := requireArrayAndCallback("khojo_shesh_index", args)
		if err != nil {
			return err
		}
		for i := len(arr.Elements) - 1; i >= 0; i-- {
			el := arr.Elements[i]
			keep := EvalFunc(callback, []object.Object{el, &object.Number{Value: float64(i)}, arr})
			if keep.Type() == object.ERROR_OBJ {
				return keep
			}
			if isTruthy(keep) {
				return &object.Number{Value: float64(i)}
			}
		}
		return &object.Number{Value: -1}
	}}
}

func registerArrayEvery() {
	Builtins["prottek"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		arr, callback, err := requireArrayAndCallback("prottek", args)
		if err != nil {
			return err
		}
		for i, el := range arr.Elements {
			keep := EvalFunc(callback, []object.Object{el, &object.Number{Value: float64(i)}, arr})
			if keep.Type() == object.ERROR_OBJ {
				return keep
			}
			if !isTruthy(keep) {
				return object.FALSE
			}
		}
		return object.TRUE
	}}
}

func registerArraySome() {
	Builtins["kono"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		arr, callback, err := requireArrayAndCallback("kono", args)
		if err != nil {
			return err
		}
		for i, el := range arr.Elements {
			keep := EvalFunc(callback, []object.Object{el, &object.Number{Value: float64(i)}, arr})
			if keep.Type() == object.ERROR_OBJ {
				return keep
			}
			if isTruthy(keep) {
				return object.TRUE
			}
		}
		return object.FALSE
	}}
}

func registerArrayFlatMap() {
	Builtins["somtol_manchitro"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		arr, callback, err := requireArrayAndCallback("somtol_manchitro", args)
		if err != nil {
			return err
		}
		result := make([]object.Object, 0, len(arr.Elements))
		for i, el := range arr.Elements {
			mapped := EvalFunc(callback, []object.Object{el, &object.Number{Value: float64(i)}, arr})
			if mapped.Type() == object.ERROR_OBJ {
				return mapped
			}
			if mappedArr, ok := mapped.(*object.Array); ok {
				result = append(result, mappedArr.Elements...)
			} else {
				result = append(result, mapped)
			}
		}
		return &object.Array{Elements: result}
	}}
}

func registerArrayAt() {
	Builtins["array_at"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=2", len(args))
		}
		if args[0].Type() != object.ARRAY_OBJ {
			return newError("first argument to `array_at` must be ARRAY, got %s", args[0].Type())
		}
		if args[1].Type() != object.NUMBER_OBJ {
			return newError("second argument to `array_at` must be NUMBER, got %s", args[1].Type())
		}
		arr := args[0].(*object.Array)
		idx := int(args[1].(*object.Number).Value)
		if idx < 0 {
			idx = len(arr.Elements) + idx
		}
		if idx < 0 || idx >= len(arr.Elements) {
			return object.NULL
		}
		return arr.Elements[idx]
	}}
}

func registerArrayLastIndexOf() {
	Builtins["shesh_index_of"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=2", len(args))
		}
		if args[0].Type() != object.ARRAY_OBJ {
			return newError("first argument to `shesh_index_of` must be ARRAY, got %s", args[0].Type())
		}
		arr := args[0].(*object.Array)
		target := args[1]
		for i := len(arr.Elements) - 1; i >= 0; i-- {
			if objectsEqual(arr.Elements[i], target) {
				return &object.Number{Value: float64(i)}
			}
		}
		return &object.Number{Value: -1}
	}}
}

func flattenArray(elements []object.Object, depth int) []object.Object {
	if depth == 0 {
		out := make([]object.Object, len(elements))
		copy(out, elements)
		return out
	}

	result := make([]object.Object, 0, len(elements))
	for _, el := range elements {
		if arr, ok := el.(*object.Array); ok {
			result = append(result, flattenArray(arr.Elements, depth-1)...)
		} else {
			result = append(result, el)
		}
	}
	return result
}

func requireArrayAndCallback(name string, args []object.Object) (*object.Array, *object.Function, *object.Error) {
	if len(args) != 2 {
		return nil, nil, newError("wrong number of arguments. got=%d, want=2", len(args))
	}
	if args[0].Type() != object.ARRAY_OBJ {
		return nil, nil, newError("first argument to `%s` must be ARRAY, got %s", name, args[0].Type())
	}
	if args[1].Type() != object.FUNCTION_OBJ {
		return nil, nil, newError("second argument to `%s` must be FUNCTION, got %s", name, args[1].Type())
	}
	return args[0].(*object.Array), args[1].(*object.Function), nil
}
