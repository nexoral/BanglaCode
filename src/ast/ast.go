package ast

import (
	"BanglaCode/src/lexer"
	"bytes"
	"strings"
)

// Node represents any node in the AST
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement represents a statement node
type Statement interface {
	Node
	statementNode()
}

// Expression represents an expression node
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of the AST
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// VariableDeclaration represents: dhoro x = 5;
type VariableDeclaration struct {
	Token lexer.Token // the DHORO token
	Name  *Identifier
	Value Expression
}

func (vd *VariableDeclaration) statementNode()       {}
func (vd *VariableDeclaration) TokenLiteral() string { return vd.Token.Literal }
func (vd *VariableDeclaration) String() string {
	var out bytes.Buffer
	out.WriteString("dhoro ")
	out.WriteString(vd.Name.String())
	out.WriteString(" = ")
	if vd.Value != nil {
		out.WriteString(vd.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

// Identifier represents a variable name
type Identifier struct {
	Token lexer.Token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }

// NumberLiteral represents a numeric value
type NumberLiteral struct {
	Token lexer.Token
	Value float64
}

func (nl *NumberLiteral) expressionNode()      {}
func (nl *NumberLiteral) TokenLiteral() string { return nl.Token.Literal }
func (nl *NumberLiteral) String() string       { return nl.Token.Literal }

// StringLiteral represents a string value
type StringLiteral struct {
	Token lexer.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string       { return sl.Token.Literal }

// BooleanLiteral represents sotti or mittha
type BooleanLiteral struct {
	Token lexer.Token
	Value bool
}

func (bl *BooleanLiteral) expressionNode()      {}
func (bl *BooleanLiteral) TokenLiteral() string { return bl.Token.Literal }
func (bl *BooleanLiteral) String() string       { return bl.Token.Literal }

// NullLiteral represents khali
type NullLiteral struct {
	Token lexer.Token
}

func (nl *NullLiteral) expressionNode()      {}
func (nl *NullLiteral) TokenLiteral() string { return nl.Token.Literal }
func (nl *NullLiteral) String() string       { return "khali" }

// ArrayLiteral represents [1, 2, 3]
type ArrayLiteral struct {
	Token    lexer.Token // the '[' token
	Elements []Expression
}

func (al *ArrayLiteral) expressionNode()      {}
func (al *ArrayLiteral) TokenLiteral() string { return al.Token.Literal }
func (al *ArrayLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}
	for _, el := range al.Elements {
		elements = append(elements, el.String())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

// MapLiteral represents {"key": "value"}
type MapLiteral struct {
	Token lexer.Token // the '{' token
	Pairs map[Expression]Expression
}

func (ml *MapLiteral) expressionNode()      {}
func (ml *MapLiteral) TokenLiteral() string { return ml.Token.Literal }
func (ml *MapLiteral) String() string {
	var out bytes.Buffer
	pairs := []string{}
	for key, value := range ml.Pairs {
		pairs = append(pairs, key.String()+":"+value.String())
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

// BinaryExpression represents operations like a + b, a == b
type BinaryExpression struct {
	Token    lexer.Token // the operator token
	Left     Expression
	Operator string
	Right    Expression
}

func (be *BinaryExpression) expressionNode()      {}
func (be *BinaryExpression) TokenLiteral() string { return be.Token.Literal }
func (be *BinaryExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(be.Left.String())
	out.WriteString(" " + be.Operator + " ")
	out.WriteString(be.Right.String())
	out.WriteString(")")
	return out.String()
}

// UnaryExpression represents !x, -x, na x
type UnaryExpression struct {
	Token    lexer.Token // the operator token
	Operator string
	Right    Expression
}

func (ue *UnaryExpression) expressionNode()      {}
func (ue *UnaryExpression) TokenLiteral() string { return ue.Token.Literal }
func (ue *UnaryExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ue.Operator)
	out.WriteString(ue.Right.String())
	out.WriteString(")")
	return out.String()
}

// AssignmentExpression represents x = 5, x += 5
type AssignmentExpression struct {
	Token    lexer.Token // the = token
	Name     Expression
	Operator string // =, +=, -=, *=, /=
	Value    Expression
}

func (ae *AssignmentExpression) expressionNode()      {}
func (ae *AssignmentExpression) TokenLiteral() string { return ae.Token.Literal }
func (ae *AssignmentExpression) String() string {
	var out bytes.Buffer
	out.WriteString(ae.Name.String())
	out.WriteString(" " + ae.Operator + " ")
	out.WriteString(ae.Value.String())
	return out.String()
}

// CallExpression represents function calls: add(1, 2)
type CallExpression struct {
	Token     lexer.Token // the '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}
	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

// MemberExpression represents property access: obj.prop, arr[0]
type MemberExpression struct {
	Token    lexer.Token // the '.' or '[' token
	Object   Expression
	Property Expression
	Computed bool // true for arr[0], false for obj.prop
}

func (me *MemberExpression) expressionNode()      {}
func (me *MemberExpression) TokenLiteral() string { return me.Token.Literal }
func (me *MemberExpression) String() string {
	var out bytes.Buffer
	out.WriteString(me.Object.String())
	if me.Computed {
		out.WriteString("[")
		out.WriteString(me.Property.String())
		out.WriteString("]")
	} else {
		out.WriteString(".")
		out.WriteString(me.Property.String())
	}
	return out.String()
}

// ExpressionStatement wraps an expression as a statement
type ExpressionStatement struct {
	Token      lexer.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// BlockStatement represents a code block: { ... }
type BlockStatement struct {
	Token      lexer.Token // the '{' token
	Statements []Statement
}

func (bs *BlockStatement) statementNode()       {}
func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BlockStatement) String() string {
	var out bytes.Buffer
	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

// IfStatement represents: jodi (condition) { ... } nahole { ... }
type IfStatement struct {
	Token       lexer.Token // the JODI token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (is *IfStatement) statementNode()       {}
func (is *IfStatement) TokenLiteral() string { return is.Token.Literal }
func (is *IfStatement) String() string {
	var out bytes.Buffer
	out.WriteString("jodi ")
	out.WriteString(is.Condition.String())
	out.WriteString(" ")
	out.WriteString(is.Consequence.String())
	if is.Alternative != nil {
		out.WriteString(" nahole ")
		out.WriteString(is.Alternative.String())
	}
	return out.String()
}

// WhileStatement represents: jotokkhon (condition) { ... }
type WhileStatement struct {
	Token     lexer.Token // the JOTOKKHON token
	Condition Expression
	Body      *BlockStatement
}

func (ws *WhileStatement) statementNode()       {}
func (ws *WhileStatement) TokenLiteral() string { return ws.Token.Literal }
func (ws *WhileStatement) String() string {
	var out bytes.Buffer
	out.WriteString("jotokkhon ")
	out.WriteString(ws.Condition.String())
	out.WriteString(" ")
	out.WriteString(ws.Body.String())
	return out.String()
}

// ForStatement represents: ghuriye (init; condition; update) { ... }
type ForStatement struct {
	Token     lexer.Token // the GHURIYE token
	Init      Statement
	Condition Expression
	Update    Expression
	Body      *BlockStatement
}

func (fs *ForStatement) statementNode()       {}
func (fs *ForStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForStatement) String() string {
	var out bytes.Buffer
	out.WriteString("ghuriye (")
	if fs.Init != nil {
		out.WriteString(fs.Init.String())
	}
	out.WriteString("; ")
	if fs.Condition != nil {
		out.WriteString(fs.Condition.String())
	}
	out.WriteString("; ")
	if fs.Update != nil {
		out.WriteString(fs.Update.String())
	}
	out.WriteString(") ")
	out.WriteString(fs.Body.String())
	return out.String()
}

// FunctionLiteral represents: kaj(a, b) { ... }
type FunctionLiteral struct {
	Token      lexer.Token // the KAJ token
	Name       *Identifier // optional function name
	Parameters []*Identifier
	Body       *BlockStatement
}

func (fl *FunctionLiteral) expressionNode()      {}
func (fl *FunctionLiteral) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionLiteral) String() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("kaj")
	if fl.Name != nil {
		out.WriteString(" " + fl.Name.String())
	}
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(fl.Body.String())
	return out.String()
}

// ReturnStatement represents: ferao x;
type ReturnStatement struct {
	Token       lexer.Token // the FERAO token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString("ferao ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

// ClassDeclaration represents: class Manush { ... }
type ClassDeclaration struct {
	Token   lexer.Token // the CLASS token
	Name    *Identifier
	Methods []*FunctionLiteral
}

func (cd *ClassDeclaration) statementNode()       {}
func (cd *ClassDeclaration) TokenLiteral() string { return cd.Token.Literal }
func (cd *ClassDeclaration) String() string {
	var out bytes.Buffer
	out.WriteString("class ")
	out.WriteString(cd.Name.String())
	out.WriteString(" { ")
	for _, method := range cd.Methods {
		out.WriteString(method.String())
	}
	out.WriteString(" }")
	return out.String()
}

// NewExpression represents: notun Manush(args)
type NewExpression struct {
	Token     lexer.Token // the NOTUN token
	Class     Expression
	Arguments []Expression
}

func (ne *NewExpression) expressionNode()      {}
func (ne *NewExpression) TokenLiteral() string { return ne.Token.Literal }
func (ne *NewExpression) String() string {
	var out bytes.Buffer
	args := []string{}
	for _, a := range ne.Arguments {
		args = append(args, a.String())
	}
	out.WriteString("notun ")
	out.WriteString(ne.Class.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")
	return out.String()
}

// BreakStatement represents: thamo;
type BreakStatement struct {
	Token lexer.Token // the THAMO token
}

func (bs *BreakStatement) statementNode()       {}
func (bs *BreakStatement) TokenLiteral() string { return bs.Token.Literal }
func (bs *BreakStatement) String() string       { return "thamo;" }

// ContinueStatement represents: chharo;
type ContinueStatement struct {
	Token lexer.Token // the CHHARO token
}

func (cs *ContinueStatement) statementNode()       {}
func (cs *ContinueStatement) TokenLiteral() string { return cs.Token.Literal }
func (cs *ContinueStatement) String() string       { return "chharo;" }

// ImportStatement represents: anayat "module.bang";
type ImportStatement struct {
	Token    lexer.Token // the ANO token
	Path     *StringLiteral
	Alias    *Identifier // optional: ano "module.bang" as mod;
}

func (is *ImportStatement) statementNode()       {}
func (is *ImportStatement) TokenLiteral() string { return is.Token.Literal }
func (is *ImportStatement) String() string {
	var out bytes.Buffer
	out.WriteString("ano ")
	out.WriteString("\"" + is.Path.Value + "\"")
	if is.Alias != nil {
		out.WriteString(" as " + is.Alias.Value)
	}
	out.WriteString(";")
	return out.String()
}

// ExportStatement represents: pathao kaj funcName() { }
type ExportStatement struct {
	Token     lexer.Token // the PATHAO token
	Statement Statement   // the exported statement (function, class, or variable)
}

func (es *ExportStatement) statementNode()       {}
func (es *ExportStatement) TokenLiteral() string { return es.Token.Literal }
func (es *ExportStatement) String() string {
	var out bytes.Buffer
	out.WriteString("pathao ")
	out.WriteString(es.Statement.String())
	return out.String()
}

// TryCatchStatement represents: chesta { } dhoro_bhul (e) { } shesh { }
type TryCatchStatement struct {
	Token       lexer.Token     // the CHESTA token
	TryBlock    *BlockStatement
	CatchParam  *Identifier     // parameter in catch block
	CatchBlock  *BlockStatement
	FinallyBlock *BlockStatement // optional
}

func (tcs *TryCatchStatement) statementNode()       {}
func (tcs *TryCatchStatement) TokenLiteral() string { return tcs.Token.Literal }
func (tcs *TryCatchStatement) String() string {
	var out bytes.Buffer
	out.WriteString("chesta ")
	out.WriteString(tcs.TryBlock.String())
	if tcs.CatchBlock != nil {
		out.WriteString(" dhoro_bhul (")
		if tcs.CatchParam != nil {
			out.WriteString(tcs.CatchParam.Value)
		}
		out.WriteString(") ")
		out.WriteString(tcs.CatchBlock.String())
	}
	if tcs.FinallyBlock != nil {
		out.WriteString(" shesh ")
		out.WriteString(tcs.FinallyBlock.String())
	}
	return out.String()
}

// ThrowStatement represents: felo "error message";
type ThrowStatement struct {
	Token lexer.Token // the FELO token
	Value Expression
}

func (ts *ThrowStatement) statementNode()       {}
func (ts *ThrowStatement) TokenLiteral() string { return ts.Token.Literal }
func (ts *ThrowStatement) String() string {
	var out bytes.Buffer
	out.WriteString("felo ")
	if ts.Value != nil {
		out.WriteString(ts.Value.String())
	}
	out.WriteString(";")
	return out.String()
}
