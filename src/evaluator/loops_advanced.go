package evaluator

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/object"
	"sort"
)

func evalForOfStatement(stmt *ast.ForOfStatement, env *object.Environment) object.Object {
	iterable := Eval(stmt.Iterable, env)
	if isError(iterable) {
		return iterable
	}

	loopEnv := object.NewEnclosedEnvironment(env)
	elements, err := toForOfElements(iterable)
	if err != nil {
		return err
	}

	for _, el := range elements {
		loopEnv.Update(stmt.VarName.Value, el)
		result := Eval(stmt.Body, loopEnv)
		if result != nil {
			switch result.Type() {
			case object.RETURN_OBJ, object.ERROR_OBJ, object.EXCEPTION_OBJ:
				return result
			case object.BREAK_OBJ:
				return object.NULL
			case object.CONTINUE_OBJ:
				continue
			}
		}
	}

	return object.NULL
}

func evalForInStatement(stmt *ast.ForInStatement, env *object.Environment) object.Object {
	target := Eval(stmt.Object, env)
	if isError(target) {
		return target
	}

	loopEnv := object.NewEnclosedEnvironment(env)
	keys, err := toForInKeys(target)
	if err != nil {
		return err
	}

	for _, key := range keys {
		loopEnv.Update(stmt.VarName.Value, key)
		result := Eval(stmt.Body, loopEnv)
		if result != nil {
			switch result.Type() {
			case object.RETURN_OBJ, object.ERROR_OBJ, object.EXCEPTION_OBJ:
				return result
			case object.BREAK_OBJ:
				return object.NULL
			case object.CONTINUE_OBJ:
				continue
			}
		}
	}

	return object.NULL
}

func toForOfElements(iterable object.Object) ([]object.Object, *object.Error) {
	switch it := iterable.(type) {
	case *object.Array:
		elements := make([]object.Object, len(it.Elements))
		copy(elements, it.Elements)
		return elements, nil

	case *object.String:
		runes := []rune(it.Value)
		elements := make([]object.Object, 0, len(runes))
		for _, r := range runes {
			elements = append(elements, &object.String{Value: string(r)})
		}
		return elements, nil

	case *object.Map:
		keys := make([]string, 0, len(it.Pairs))
		for k := range it.Pairs {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		elements := make([]object.Object, 0, len(keys))
		for _, k := range keys {
			elements = append(elements, it.Pairs[k])
		}
		return elements, nil
	}

	return nil, newError("for...of target must be ARRAY, STRING, or MAP, got %s", iterable.Type())
}

func toForInKeys(target object.Object) ([]object.Object, *object.Error) {
	switch it := target.(type) {
	case *object.Map:
		keys := make([]string, 0, len(it.Pairs))
		for k := range it.Pairs {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		out := make([]object.Object, 0, len(keys))
		for _, k := range keys {
			out = append(out, &object.String{Value: k})
		}
		return out, nil

	case *object.Array:
		out := make([]object.Object, 0, len(it.Elements))
		for i := range it.Elements {
			out = append(out, &object.Number{Value: float64(i)})
		}
		return out, nil

	case *object.String:
		runes := []rune(it.Value)
		out := make([]object.Object, 0, len(runes))
		for i := range runes {
			out = append(out, &object.Number{Value: float64(i)})
		}
		return out, nil
	}

	return nil, newError("for...in target must be MAP, ARRAY, or STRING, got %s", target.Type())
}
