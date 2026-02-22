package evaluator

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/object"
)

type yieldResult struct {
	Value object.Object
}

func (yr *yieldResult) Type() object.ObjectType { return object.GENERATOR_OBJ }
func (yr *yieldResult) Inspect() string         { return "yield " + yr.Value.Inspect() }

func evalGeneratorFunction(fn *object.Function, args []object.Object, env *object.Environment) object.Object {
	extendedEnv := extendFunctionEnv(fn, args)
	return &object.Generator{
		Function: fn,
		Env:      extendedEnv,
		State:    "suspended",
		Value:    object.NULL,
		Index:    0,
		Done:     false,
	}
}

func generatorNext(gen *object.Generator) *object.Map {
	if gen.Done {
		return generatorResult(object.NULL, true)
	}

	gen.State = "executing"
	stmts := gen.Function.Body.Statements
	for i := gen.Index; i < len(stmts); i++ {
		result := evalGeneratorStatement(stmts[i], gen)

		switch v := result.(type) {
		case *yieldResult:
			gen.Index = i + 1
			gen.State = "suspended"
			gen.Value = v.Value
			return generatorResult(v.Value, false)
		case *object.ReturnValue:
			gen.Done = true
			gen.State = "completed"
			gen.Value = v.Value
			return generatorResult(v.Value, true)
		case *object.Exception:
			gen.Done = true
			gen.State = "completed"
			return &object.Map{
				Pairs: map[string]object.Object{
					"value": &object.String{Value: v.Inspect()},
					"done":  object.TRUE,
				},
			}
		}

		if isError(result) {
			gen.Done = true
			gen.State = "completed"
			return generatorResult(result, true)
		}
	}

	gen.Done = true
	gen.State = "completed"
	gen.Value = object.NULL
	return generatorResult(object.NULL, true)
}

func generatorReturn(gen *object.Generator, value object.Object) *object.Map {
	if value == nil {
		value = object.NULL
	}
	gen.Done = true
	gen.State = "completed"
	gen.Value = value
	return generatorResult(value, true)
}

func generatorThrow(gen *object.Generator, thrown object.Object) object.Object {
	if thrown == nil {
		thrown = &object.String{Value: "generator throw"}
	}
	gen.Done = true
	gen.State = "completed"
	return &object.Exception{Message: thrown.Inspect(), Value: thrown}
}

func generatorResult(value object.Object, done bool) *object.Map {
	doneObj := object.FALSE
	if done {
		doneObj = object.TRUE
	}
	if value == nil {
		value = object.NULL
	}
	return &object.Map{
		Pairs: map[string]object.Object{
			"value": value,
			"done":  doneObj,
		},
	}
}

func evalGeneratorStatement(stmt ast.Statement, gen *object.Generator) object.Object {
	switch s := stmt.(type) {
	case *ast.ExpressionStatement:
		return evalGeneratorExpression(s.Expression, gen)
	case *ast.ReturnStatement:
		val := evalGeneratorExpression(s.ReturnValue, gen)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}
	case *ast.VariableDeclaration:
		val := Eval(s.Value, gen.Env)
		if isError(val) {
			return val
		}
		if s.IsConstant {
			gen.Env.SetConstant(s.Name.Value, val)
		} else if s.IsGlobal {
			gen.Env.SetGlobal(s.Name.Value, val)
		} else {
			gen.Env.Set(s.Name.Value, val)
		}
		return val
	default:
		return Eval(stmt, gen.Env)
	}
}

func evalGeneratorExpression(expr ast.Expression, gen *object.Generator) object.Object {
	if expr == nil {
		return object.NULL
	}
	switch e := expr.(type) {
	case *ast.YieldExpression:
		var value object.Object = object.NULL
		if e.Expression != nil {
			value = Eval(e.Expression, gen.Env)
			if isError(value) {
				return value
			}
		}
		return &yieldResult{Value: value}
	default:
		return Eval(expr, gen.Env)
	}
}
