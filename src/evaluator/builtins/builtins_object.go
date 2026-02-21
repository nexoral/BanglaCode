package builtins

import (
	"BanglaCode/src/object"
	"math"
)

func init() {
	registerObjectValues()
	registerObjectEntries()
	registerObjectAssign()
	registerObjectHasOwn()
	registerObjectFromEntries()
	registerObjectIs()
	registerObjectCreate()
	registerObjectFreeze()
}

func registerObjectValues() {
	Builtins["maan"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1", len(args))
		}
		if args[0].Type() != object.MAP_OBJ {
			return newError("argument to `maan` must be MAP, got %s", args[0].Type())
		}
		mapObj := args[0].(*object.Map)
		values := make([]object.Object, 0, len(mapObj.Pairs))
		for _, value := range mapObj.Pairs {
			values = append(values, value)
		}
		return &object.Array{Elements: values}
	}}
}

func registerObjectEntries() {
	Builtins["jora"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1", len(args))
		}
		if args[0].Type() != object.MAP_OBJ {
			return newError("argument to `jora` must be MAP, got %s", args[0].Type())
		}
		mapObj := args[0].(*object.Map)
		entries := make([]object.Object, 0, len(mapObj.Pairs))
		for key, value := range mapObj.Pairs {
			entry := &object.Array{Elements: []object.Object{&object.String{Value: key}, value}}
			entries = append(entries, entry)
		}
		return &object.Array{Elements: entries}
	}}
}

func registerObjectAssign() {
	Builtins["mishra"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) < 2 {
			return newError("wrong number of arguments. got=%d, want at least 2", len(args))
		}
		if args[0].Type() != object.MAP_OBJ {
			return newError("first argument to `mishra` must be MAP, got %s", args[0].Type())
		}
		target := args[0].(*object.Map)
		for i := 1; i < len(args); i++ {
			if args[i].Type() != object.MAP_OBJ {
				return newError("argument %d to `mishra` must be MAP, got %s", i+1, args[i].Type())
			}
			source := args[i].(*object.Map)
			for key, value := range source.Pairs {
				target.Pairs[key] = value
			}
		}
		return target
	}}
}

func registerObjectHasOwn() {
	Builtins["nijer_ache"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=2", len(args))
		}
		if args[0].Type() != object.MAP_OBJ {
			return newError("first argument to `nijer_ache` must be MAP, got %s", args[0].Type())
		}
		key := mapKeyFromObject(args[1])
		if key == "" && args[1].Type() != object.STRING_OBJ && args[1].Type() != object.NUMBER_OBJ {
			return object.FALSE
		}
		_, exists := args[0].(*object.Map).Pairs[key]
		return object.NativeBoolToBooleanObject(exists)
	}}
}

func registerObjectFromEntries() {
	Builtins["jora_theke"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1", len(args))
		}
		if args[0].Type() != object.ARRAY_OBJ {
			return newError("argument to `jora_theke` must be ARRAY, got %s", args[0].Type())
		}
		entries := args[0].(*object.Array)
		result := make(map[string]object.Object, len(entries.Elements))
		for i, entryObj := range entries.Elements {
			entry, ok := entryObj.(*object.Array)
			if !ok || len(entry.Elements) < 2 {
				return newError("entry at index %d must be [key, value]", i)
			}
			key := mapKeyFromObject(entry.Elements[0])
			if key == "" && entry.Elements[0].Type() != object.STRING_OBJ && entry.Elements[0].Type() != object.NUMBER_OBJ {
				return newError("entry key at index %d must be STRING or NUMBER", i)
			}
			result[key] = entry.Elements[1]
		}
		return &object.Map{Pairs: result}
	}}
}

func registerObjectIs() {
	Builtins["ekoi_ki"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=2", len(args))
		}
		if args[0].Type() == object.NUMBER_OBJ && args[1].Type() == object.NUMBER_OBJ {
			a := args[0].(*object.Number).Value
			b := args[1].(*object.Number).Value
			if math.IsNaN(a) && math.IsNaN(b) {
				return object.TRUE
			}
			return object.NativeBoolToBooleanObject(a == b)
		}
		return object.NativeBoolToBooleanObject(objectsEqual(args[0], args[1]))
	}}
}

func registerObjectCreate() {
	Builtins["notun_map"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) < 1 || len(args) > 2 {
			return newError("wrong number of arguments. got=%d, want=1 or 2", len(args))
		}
		if args[0].Type() != object.MAP_OBJ && args[0].Type() != object.NULL_OBJ {
			return newError("first argument to `notun_map` must be MAP or NULL, got %s", args[0].Type())
		}
		out := &object.Map{Pairs: make(map[string]object.Object)}
		if args[0].Type() == object.MAP_OBJ {
			for k, v := range args[0].(*object.Map).Pairs {
				out.Pairs[k] = v
			}
		}
		if len(args) == 2 {
			if args[1].Type() != object.MAP_OBJ {
				return newError("second argument to `notun_map` must be MAP, got %s", args[1].Type())
			}
			for k, v := range args[1].(*object.Map).Pairs {
				out.Pairs[k] = v
			}
		}
		return out
	}}
}

func registerObjectFreeze() {
	Builtins["joma"] = &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1", len(args))
		}
		if args[0].Type() != object.MAP_OBJ {
			return newError("argument to `joma` must be MAP, got %s", args[0].Type())
		}
		return args[0]
	}}
}
