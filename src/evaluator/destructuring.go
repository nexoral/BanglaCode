package evaluator

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/object"
)

func evalArrayDestructuringDeclaration(node *ast.ArrayDestructuringDeclaration, env *object.Environment) object.Object {
	source := Eval(node.Source, env)
	if isError(source) {
		return source
	}
	arr, ok := source.(*object.Array)
	if !ok {
		return newError("array destructuring source must be ARRAY, got %s", source.Type())
	}

	for i, name := range node.Names {
		var val object.Object = object.NULL
		if i < len(arr.Elements) {
			val = arr.Elements[i]
		}
		bindValue(env, name.Value, val, node.IsConstant, node.IsGlobal)
	}

	return source
}

func evalObjectDestructuringDeclaration(node *ast.ObjectDestructuringDeclaration, env *object.Environment) object.Object {
	source := Eval(node.Source, env)
	if isError(source) {
		return source
	}
	m, ok := source.(*object.Map)
	if !ok {
		return newError("object destructuring source must be MAP, got %s", source.Type())
	}

	for i, key := range node.Keys {
		val, exists := m.Pairs[key]
		if !exists {
			val = object.NULL
		}
		bindValue(env, node.Names[i].Value, val, node.IsConstant, node.IsGlobal)
	}

	return source
}

func bindValue(env *object.Environment, name string, val object.Object, isConstant, isGlobal bool) {
	if isConstant {
		env.SetConstant(name, val)
		return
	}
	if isGlobal {
		env.SetGlobal(name, val)
		return
	}
	env.Set(name, val)
}
