package ast

import (
	"BanglaCode/src/lexer"
	"bytes"
)

// ForOfStatement represents: ghuriye (item of iterable) { ... }
type ForOfStatement struct {
	Token    lexer.Token // GHURIYE token
	VarName  *Identifier
	Iterable Expression
	Body     *BlockStatement
}

func (fs *ForOfStatement) statementNode()       {}
func (fs *ForOfStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForOfStatement) String() string {
	var out bytes.Buffer
	out.WriteString("ghuriye (")
	out.WriteString(fs.VarName.String())
	out.WriteString(" of ")
	out.WriteString(fs.Iterable.String())
	out.WriteString(") ")
	out.WriteString(fs.Body.String())
	return out.String()
}

// ForInStatement represents: ghuriye (key in object) { ... }
type ForInStatement struct {
	Token   lexer.Token // GHURIYE token
	VarName *Identifier
	Object  Expression
	Body    *BlockStatement
}

func (fs *ForInStatement) statementNode()       {}
func (fs *ForInStatement) TokenLiteral() string { return fs.Token.Literal }
func (fs *ForInStatement) String() string {
	var out bytes.Buffer
	out.WriteString("ghuriye (")
	out.WriteString(fs.VarName.String())
	out.WriteString(" in ")
	out.WriteString(fs.Object.String())
	out.WriteString(") ")
	out.WriteString(fs.Body.String())
	return out.String()
}
