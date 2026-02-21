package ast

import (
	"BanglaCode/src/lexer"
	"bytes"
)

// DoWhileStatement represents: do { ... } jotokkhon (condition);
type DoWhileStatement struct {
	Token     lexer.Token // the DO token
	Body      *BlockStatement
	Condition Expression
}

func (dws *DoWhileStatement) statementNode()       {}
func (dws *DoWhileStatement) TokenLiteral() string { return dws.Token.Literal }
func (dws *DoWhileStatement) String() string {
	var out bytes.Buffer
	out.WriteString("do ")
	out.WriteString(dws.Body.String())
	out.WriteString(" jotokkhon (")
	out.WriteString(dws.Condition.String())
	out.WriteString(");")
	return out.String()
}

// DeleteExpression represents: delete obj.prop or delete obj["prop"]
type DeleteExpression struct {
	Token  lexer.Token // the DELETE token
	Target Expression
}

func (de *DeleteExpression) expressionNode()      {}
func (de *DeleteExpression) TokenLiteral() string { return de.Token.Literal }
func (de *DeleteExpression) String() string {
	var out bytes.Buffer
	out.WriteString("delete ")
	out.WriteString(de.Target.String())
	return out.String()
}
