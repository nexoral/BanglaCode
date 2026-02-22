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

	case *object.Class:
		return assignClassMember(o, member, operator, val)

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

	case *object.Class:
		return accessClassMember(o, me)

	case *object.URL:
		return accessURLMember(o, me)

	case *object.Stream:
		return accessStreamMember(o, me)

	case *object.Buffer:
		return accessBufferMember(o, me)

	case *object.Generator:
		return accessGeneratorMember(o, me)

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

	propName := ident.Value

	// Check if setter exists for this property
	if setter, ok := inst.Class.Setters[propName]; ok {
		// Execute setter with 'ei' bound to instance and value as parameter
		boundEnv := object.NewEnclosedEnvironment(setter.Env)
		boundEnv.Set("ei", inst)

		// Bind the value to the setter parameter
		if len(setter.Parameters) > 0 {
			boundEnv.Set(setter.Parameters[0].Value, val)
		}

		result := Eval(setter.Body, boundEnv)
		if isError(result) {
			return result
		}
		return val // Return the assigned value
	}

	// Check if property is private (starts with _)
	if len(propName) > 0 && propName[0] == '_' {
		// Assign to private field
		if operator != "=" {
			current, ok := inst.PrivateFields[propName]
			if !ok {
				return newError("private property '%s' not found", propName)
			}
			op := string(operator[0])
			val = evalBinaryExpression(op, current, val)
			if isError(val) {
				return val
			}
		}
		inst.PrivateFields[propName] = val
		return val
	}

	// Normal property assignment
	if operator != "=" {
		current, ok := inst.Properties[propName]
		if !ok {
			return newError("property '%s' not found", propName)
		}
		op := string(operator[0])
		val = evalBinaryExpression(op, current, val)
		if isError(val) {
			return val
		}
	}

	inst.Properties[propName] = val
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

	propName := ident.Value

	// Check if property is private (starts with _)
	if len(propName) > 0 && propName[0] == '_' {
		// Access private field
		if val, ok := inst.PrivateFields[propName]; ok {
			return val
		}
		// If not found in private fields, continue to normal properties
	}

	// Check if property exists
	if val, ok := inst.Properties[propName]; ok {
		return val
	}

	// Check if getter exists for this property
	if getter, ok := inst.Class.Getters[propName]; ok {
		// Execute getter with 'ei' bound to instance
		boundEnv := object.NewEnclosedEnvironment(getter.Env)
		boundEnv.Set("ei", inst)
		result := Eval(getter.Body, boundEnv)
		return unwrapReturnValue(result)
	}

	// Check if method exists
	if method, ok := inst.Class.Methods[propName]; ok {
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

// accessClassMember accesses static properties of a class
func accessClassMember(class *object.Class, me *ast.MemberExpression) object.Object {
	ident, ok := me.Property.(*ast.Identifier)
	if !ok {
		return newError("invalid property name")
	}

	propName := ident.Value

	// Check if static property exists
	if val, ok := class.StaticProperties[propName]; ok {
		return val
	}

	return newError("class '%s' has no static property '%s'", class.Name, propName)
}

func accessGeneratorMember(gen *object.Generator, me *ast.MemberExpression) object.Object {
	ident, ok := me.Property.(*ast.Identifier)
	if !ok {
		return newError("invalid generator member")
	}

	switch ident.Value {
	case "next":
		return &object.Builtin{
			Fn: func(args ...object.Object) object.Object {
				return generatorNext(gen)
			},
		}
	case "return":
		return &object.Builtin{
			Fn: func(args ...object.Object) object.Object {
				if len(args) > 0 {
					return generatorReturn(gen, args[0])
				}
				return generatorReturn(gen, object.NULL)
			},
		}
	case "throw":
		return &object.Builtin{
			Fn: func(args ...object.Object) object.Object {
				if len(args) > 0 {
					return generatorThrow(gen, args[0])
				}
				return generatorThrow(gen, &object.String{Value: "generator throw"})
			},
		}
	default:
		return object.NULL
	}
}

// assignClassMember assigns to static properties of a class
func assignClassMember(class *object.Class, member *ast.MemberExpression, operator string, val object.Object) object.Object {
	ident, ok := member.Property.(*ast.Identifier)
	if !ok {
		return newError("invalid property name")
	}

	propName := ident.Value

	if operator != "=" {
		current, ok := class.StaticProperties[propName]
		if !ok {
			return newError("static property '%s' not found in class '%s'", propName, class.Name)
		}
		op := string(operator[0])
		val = evalBinaryExpression(op, current, val)
		if isError(val) {
			return val
		}
	}

	class.StaticProperties[propName] = val
	return val
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

// accessURLMember accesses URL object properties
func accessURLMember(url *object.URL, me *ast.MemberExpression) object.Object {
	if me.Computed {
		return newError("computed member access not supported on URL")
	}

	prop := me.Property.(*ast.Identifier).Value

	switch prop {
	case "Href", "href":
		return &object.String{Value: url.Href}
	case "Protocol", "protocol":
		return &object.String{Value: url.Protocol}
	case "Username", "username":
		return &object.String{Value: url.Username}
	case "Password", "password":
		return &object.String{Value: url.Password}
	case "Hostname", "hostname":
		return &object.String{Value: url.Hostname}
	case "Port", "port":
		return &object.String{Value: url.Port}
	case "Host", "host":
		return &object.String{Value: url.Host}
	case "Pathname", "pathname":
		return &object.String{Value: url.Pathname}
	case "Search", "search":
		return &object.String{Value: url.Search}
	case "Hash", "hash":
		return &object.String{Value: url.Hash}
	case "Origin", "origin":
		return &object.String{Value: url.Origin}
	default:
		return newError("URL has no property '%s'", prop)
	}
}

// accessStreamMember accesses Stream object properties
func accessStreamMember(stream *object.Stream, me *ast.MemberExpression) object.Object {
	if me.Computed {
		return newError("computed member access not supported on Stream")
	}

	prop := me.Property.(*ast.Identifier).Value

	switch prop {
	case "Buffer":
		// Return buffer as Buffer object
		return &object.Buffer{Data: stream.Buffer}
	case "Type", "StreamType":
		return &object.String{Value: stream.StreamType}
	case "IsClosed":
		return object.NativeBoolToBooleanObject(stream.IsClosed)
	case "IsEnded":
		return object.NativeBoolToBooleanObject(stream.IsEnded)
	case "HighWaterMark":
		return &object.Number{Value: float64(stream.HighWaterMark)}
	default:
		return newError("Stream has no property '%s'", prop)
	}
}

// accessBufferMember accesses Buffer object properties
func accessBufferMember(buffer *object.Buffer, me *ast.MemberExpression) object.Object {
	if me.Computed {
		return newError("computed member access not supported on Buffer")
	}

	prop := me.Property.(*ast.Identifier).Value

	switch prop {
	case "Length", "length":
		return &object.Number{Value: float64(len(buffer.Data))}
	default:
		return newError("Buffer has no property '%s'", prop)
	}
}
