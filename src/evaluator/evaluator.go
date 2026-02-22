// Package evaluator implements the tree-walking interpreter for BanglaCode.
// It evaluates AST nodes and produces object values.
package evaluator

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/evaluator/builtins"
	"BanglaCode/src/evaluator/builtins/collections"
	"BanglaCode/src/evaluator/builtins/events"
	"BanglaCode/src/evaluator/builtins/streams"
	"BanglaCode/src/evaluator/builtins/worker"
	"BanglaCode/src/object"
)

func init() {
	// Set up EvalFunc for builtins that need to call back into the evaluator
	builtins.EvalFunc = evalFunctionCall
	events.SetEvalFunc(evalFunctionCall)
	worker.SetEvalFunc(Eval)
	streams.SetEvalFunc(Eval)
	collections.SetEvalFunc(evalFunctionCall)
}

// evalFunctionCall evaluates a function with the given arguments
// Used by builtins that need to call back into the evaluator
func evalFunctionCall(handler *object.Function, args []object.Object) object.Object {
	env := object.NewEnclosedEnvironment(handler.Env)
	for i, param := range handler.Parameters {
		if i < len(args) {
			env.Set(param.Value, args[i])
		}
	}
	result := Eval(handler.Body, env)
	return unwrapReturnValue(result)
}

// Eval evaluates an AST node and returns the resulting object
func Eval(node ast.Node, env *object.Environment) object.Object {
	if out, ok := evalStatementNode(node, env); ok {
		return out
	}
	if out, ok := evalControlNode(node, env); ok {
		return out
	}
	if out, ok := evalLiteralNode(node, env); ok {
		return out
	}
	if out, ok := evalExpressionNode(node, env); ok {
		return out
	}
	return nil
}

func evalStatementNode(node ast.Node, env *object.Environment) (object.Object, bool) {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node.Statements, env), true
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env), true
	case *ast.BlockStatement:
		return evalBlockStatement(node, env), true
	case *ast.VariableDeclaration:
		val := Eval(node.Value, env)
		if isError(val) {
			return val, true
		}
		if node.IsConstant {
			env.SetConstant(node.Name.Value, val)
		} else if node.IsGlobal {
			env.SetGlobal(node.Name.Value, val)
		} else {
			env.Set(node.Name.Value, val)
		}
		return val, true
	case *ast.ArrayDestructuringDeclaration:
		return evalArrayDestructuringDeclaration(node, env), true
	case *ast.ObjectDestructuringDeclaration:
		return evalObjectDestructuringDeclaration(node, env), true
	}
	return evalFlowStatementNode(node, env)
}

func evalFlowStatementNode(node ast.Node, env *object.Environment) (object.Object, bool) {
	switch node := node.(type) {
	case *ast.IfStatement:
		return evalIfStatement(node, env), true
	case *ast.WhileStatement:
		return evalWhileStatement(node, env), true
	case *ast.DoWhileStatement:
		return evalDoWhileStatement(node, env), true
	case *ast.ForStatement:
		return evalForStatement(node, env), true
	case *ast.ForOfStatement:
		return evalForOfStatement(node, env), true
	case *ast.ForInStatement:
		return evalForInStatement(node, env), true
	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val, true
		}
		return &object.ReturnValue{Value: val}, true
	case *ast.BreakStatement:
		return object.BREAK, true
	case *ast.ContinueStatement:
		return object.CONTINUE, true
	case *ast.SwitchStatement:
		return evalSwitchStatement(node, env), true
	}
	return nil, false
}

func evalControlNode(node ast.Node, env *object.Environment) (object.Object, bool) {
	switch node := node.(type) {
	case *ast.ClassDeclaration:
		return evalClassDeclaration(node, env), true
	case *ast.ImportStatement:
		return evalImportStatement(node, env), true
	case *ast.ExportStatement:
		return evalExportStatement(node, env), true
	case *ast.TryCatchStatement:
		return evalTryCatchStatement(node, env), true
	case *ast.ThrowStatement:
		return evalThrowStatement(node, env), true
	}
	return nil, false
}

func evalLiteralNode(node ast.Node, env *object.Environment) (object.Object, bool) {
	switch node := node.(type) {
	case *ast.NumberLiteral:
		return &object.Number{Value: node.Value}, true
	case *ast.StringLiteral:
		return &object.String{Value: node.Value}, true
	case *ast.TemplateLiteral:
		return evalTemplateLiteral(node, env), true
	case *ast.BooleanLiteral:
		return object.NativeBoolToBooleanObject(node.Value), true
	case *ast.NullLiteral:
		return object.NULL, true
	case *ast.ArrayLiteral:
		elements := evalExpressions(node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0], true
		}
		return &object.Array{Elements: elements}, true
	case *ast.MapLiteral:
		return evalMapLiteral(node, env), true
	}
	return nil, false
}

func evalExpressionNode(node ast.Node, env *object.Environment) (object.Object, bool) {
	switch node := node.(type) {
	case *ast.Identifier:
		return evalIdentifier(node, env), true
	case *ast.UnaryExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right, true
		}
		return evalUnaryExpression(node.Operator, right), true
	case *ast.BinaryExpression:
		return evalBinaryNode(node, env), true
	case *ast.DeleteExpression:
		return evalDeleteExpression(node, env), true
	case *ast.AssignmentExpression:
		return evalAssignmentExpression(node, env), true
	case *ast.CallExpression:
		return evalCallExpression(node, env), true
	case *ast.MemberExpression:
		return evalMemberExpression(node, env), true
	case *ast.FunctionLiteral:
		return buildFunctionLiteral(node, env), true
	case *ast.YieldExpression:
		return newError("utpadan (yield) can only be used inside generator function"), true
	case *ast.NewExpression:
		return evalNewExpression(node, env), true
	case *ast.SpreadElement:
		return evalSpreadElement(node, env), true
	case *ast.AsyncFunctionLiteral:
		return evalAsyncFunctionLiteral(node, env), true
	case *ast.AwaitExpression:
		return evalAwaitExpression(node, env), true
	}
	return nil, false
}

func evalBinaryNode(node *ast.BinaryExpression, env *object.Environment) object.Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}
	right := Eval(node.Right, env)
	if isError(right) {
		return right
	}
	return evalBinaryExpression(node.Operator, left, right)
}

func evalCallExpression(node *ast.CallExpression, env *object.Environment) object.Object {
	function := Eval(node.Function, env)
	if isError(function) {
		return function
	}
	args := evalExpressions(node.Arguments, env)
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}
	return applyFunctionWithPosition(function, args, env, node.Token.Line, node.Token.Column, node.Function)
}

func buildFunctionLiteral(node *ast.FunctionLiteral, env *object.Environment) object.Object {
	name := ""
	if node.Name != nil {
		name = node.Name.Value
	}
	fn := &object.Function{
		Parameters:    node.Parameters,
		RestParameter: node.RestParameter,
		Env:           env,
		Body:          node.Body,
		Name:          name,
		IsGenerator:   node.IsGenerator,
	}
	if name != "" {
		env.Set(name, fn)
	}
	return fn
}
