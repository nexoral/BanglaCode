package evaluator

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/object"
)

// evalMemberAssignment handles assignment to object properties or array elements
func evalMemberAssignment(member *ast.MemberExpression, operator string, value ast.Expression, env *object.Environment) object.Object {
	obj := Eval(member.Object, env)
	if isError(obj) {
		return obj
	}

	val := Eval(value, env)
	if isError(val) {
		return val
	}

	switch o := obj.(type) {
	case *object.Array:
		return assignArrayMember(o, member, operator, val, env)

	case *object.Map:
		return assignMapMember(o, member, operator, val, env)

	case *object.Instance:
		return assignInstanceMember(o, member, operator, val)

	default:
		return newError("cannot assign to %s", obj.Type())
	}
}

// evalMemberExpression evaluates member access (obj.prop or arr[idx])
func evalMemberExpression(me *ast.MemberExpression, env *object.Environment) object.Object {
	obj := Eval(me.Object, env)
	if isError(obj) {
		return obj
	}

	switch o := obj.(type) {
	case *object.Array:
		return accessArrayMember(o, me, env)

	case *object.Map:
		return accessMapMember(o, me, env)

	case *object.Instance:
		return accessInstanceMember(o, me)

	default:
		return newError("member access not supported on %s", obj.Type())
	}
}

func assignArrayMember(arr *object.Array, member *ast.MemberExpression, operator string, val object.Object, env *object.Environment) object.Object {
	index := Eval(member.Property, env)
	if isError(index) {
		return index
	}
	if index.Type() != object.NUMBER_OBJ {
		return newError("array index must be a number, got %s", index.Type())
	}
	idx := int(index.(*object.Number).Value)
	if idx < 0 || idx >= len(arr.Elements) {
		return newError("array index out of bounds: %d", idx)
	}

	if operator != "=" {
		current := arr.Elements[idx]
		op := string(operator[0])
		val = evalBinaryExpression(op, current, val)
		if isError(val) {
			return val
		}
	}

	arr.Elements[idx] = val
	return val
}

func assignMapMember(m *object.Map, member *ast.MemberExpression, operator string, val object.Object, env *object.Environment) object.Object {
	key, errObj := resolveMapMemberKey(member, env)
	if errObj != nil {
		return errObj
	}

	if operator != "=" {
		current, ok := m.Pairs[key]
		if !ok {
			return newError("key '%s' not found in map", key)
		}
		op := string(operator[0])
		val = evalBinaryExpression(op, current, val)
		if isError(val) {
			return val
		}
	}

	m.Pairs[key] = val
	return val
}

func assignInstanceMember(inst *object.Instance, member *ast.MemberExpression, operator string, val object.Object) object.Object {
	ident, ok := member.Property.(*ast.Identifier)
	if !ok {
		return newError("invalid property name")
	}

	if operator != "=" {
		current, ok := inst.Properties[ident.Value]
		if !ok {
			return newError("property '%s' not found", ident.Value)
		}
		op := string(operator[0])
		val = evalBinaryExpression(op, current, val)
		if isError(val) {
			return val
		}
	}

	inst.Properties[ident.Value] = val
	return val
}

func accessArrayMember(arr *object.Array, me *ast.MemberExpression, env *object.Environment) object.Object {
	index := Eval(me.Property, env)
	if isError(index) {
		return index
	}
	return evalArrayIndex(arr, index)
}

func accessMapMember(m *object.Map, me *ast.MemberExpression, env *object.Environment) object.Object {
	key, errObj := resolveMapMemberKey(me, env)
	if errObj != nil {
		return errObj
	}
	if val, ok := m.Pairs[key]; ok {
		return val
	}
	return object.NULL
}

func accessInstanceMember(inst *object.Instance, me *ast.MemberExpression) object.Object {
	ident, ok := me.Property.(*ast.Identifier)
	if !ok {
		return newError("invalid property name")
	}

	if val, ok := inst.Properties[ident.Value]; ok {
		return val
	}

	if method, ok := inst.Class.Methods[ident.Value]; ok {
		boundEnv := object.NewEnclosedEnvironment(method.Env)
		boundEnv.Set("ei", inst)
		return &object.Function{
			Parameters: method.Parameters,
			Body:       method.Body,
			Env:        boundEnv,
			Name:       method.Name,
		}
	}

	return object.NULL
}

func resolveMapMemberKey(member *ast.MemberExpression, env *object.Environment) (string, *object.Error) {
	if member.Computed {
		keyObj := Eval(member.Property, env)
		if isError(keyObj) {
			return "", keyObj.(*object.Error)
		}
		return getMapKey(keyObj), nil
	}
	ident, ok := member.Property.(*ast.Identifier)
	if !ok {
		return "", newError("invalid map key")
	}
	return ident.Value, nil
}

// evalArrayIndex evaluates array indexing
func evalArrayIndex(array *object.Array, index object.Object) object.Object {
	if index.Type() != object.NUMBER_OBJ {
		return newError("array index must be a number, got %s", index.Type())
	}

	idx := int(index.(*object.Number).Value)
	max := len(array.Elements) - 1

	if idx < 0 || idx > max {
		return object.NULL
	}

	return array.Elements[idx]
}
