package evaluator

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/object"
)

// evalDoWhileStatement evaluates: do { ... } jotokkhon (condition);
func evalDoWhileStatement(stmt *ast.DoWhileStatement, env *object.Environment) object.Object {
	for {
		result := Eval(stmt.Body, env)
		if result != nil {
			switch result.Type() {
			case object.RETURN_OBJ, object.ERROR_OBJ, object.EXCEPTION_OBJ:
				return result
			case object.BREAK_OBJ:
				return object.NULL
			case object.CONTINUE_OBJ:
				// continue to condition check
			}
		}

		condition := Eval(stmt.Condition, env)
		if isError(condition) {
			return condition
		}
		if !isTruthy(condition) {
			break
		}
	}

	return object.NULL
}

func evalDeleteExpression(node *ast.DeleteExpression, env *object.Environment) object.Object {
	member, ok := node.Target.(*ast.MemberExpression)
	if !ok {
		return object.FALSE
	}

	obj := Eval(member.Object, env)
	if isError(obj) {
		return obj
	}

	switch o := obj.(type) {
	case *object.Map:
		key, ok := resolveMemberKey(member, env)
		if !ok {
			return object.FALSE
		}
		if _, exists := o.Pairs[key]; exists {
			delete(o.Pairs, key)
			return object.TRUE
		}
		return object.TRUE

	case *object.Instance:
		key, ok := resolveMemberKey(member, env)
		if !ok {
			return object.FALSE
		}
		delete(o.Properties, key)
		return object.TRUE

	case *object.Array:
		if !member.Computed {
			return object.FALSE
		}
		idxObj := Eval(member.Property, env)
		if isError(idxObj) || idxObj.Type() != object.NUMBER_OBJ {
			return object.FALSE
		}
		idx := int(idxObj.(*object.Number).Value)
		if idx < 0 || idx >= len(o.Elements) {
			return object.TRUE
		}
		o.Elements[idx] = object.NULL
		return object.TRUE
	}

	return object.FALSE
}

func evalInOperator(left, right object.Object) object.Object {
	switch r := right.(type) {
	case *object.Map:
		key := getMapKey(left)
		if key == "" && left.Type() != object.STRING_OBJ && left.Type() != object.NUMBER_OBJ {
			return object.FALSE
		}
		_, exists := r.Pairs[key]
		return object.NativeBoolToBooleanObject(exists)

	case *object.Array:
		if left.Type() != object.NUMBER_OBJ {
			return object.FALSE
		}
		idx := int(left.(*object.Number).Value)
		return object.NativeBoolToBooleanObject(idx >= 0 && idx < len(r.Elements))

	case *object.String:
		if left.Type() != object.NUMBER_OBJ {
			return object.FALSE
		}
		idx := int(left.(*object.Number).Value)
		return object.NativeBoolToBooleanObject(idx >= 0 && idx < len([]rune(r.Value)))
	}

	return object.FALSE
}

func evalInstanceofOperator(left, right object.Object) object.Object {
	instance, isInstance := left.(*object.Instance)
	classObj, isClass := right.(*object.Class)
	if !isInstance || !isClass {
		return object.FALSE
	}
	return object.NativeBoolToBooleanObject(instance.Class == classObj)
}

func resolveMemberKey(member *ast.MemberExpression, env *object.Environment) (string, bool) {
	if member.Computed {
		keyObj := Eval(member.Property, env)
		if isError(keyObj) {
			return "", false
		}
		key := getMapKey(keyObj)
		if key == "" && keyObj.Type() != object.STRING_OBJ && keyObj.Type() != object.NUMBER_OBJ {
			return "", false
		}
		return key, true
	}

	ident, ok := member.Property.(*ast.Identifier)
	if !ok {
		return "", false
	}
	return ident.Value, true
}
