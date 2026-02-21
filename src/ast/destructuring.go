package ast

import (
	"BanglaCode/src/lexer"
	"bytes"
	"strings"
)

// ArrayDestructuringDeclaration represents: dhoro [a, b] = expr;
type ArrayDestructuringDeclaration struct {
	Token      lexer.Token // DHORO/STHIR/BISHWO token
	Names      []*Identifier
	Source     Expression
	IsConstant bool
	IsGlobal   bool
}

func (ad *ArrayDestructuringDeclaration) statementNode()       {}
func (ad *ArrayDestructuringDeclaration) TokenLiteral() string { return ad.Token.Literal }
func (ad *ArrayDestructuringDeclaration) String() string {
	var out bytes.Buffer
	if ad.IsConstant {
		out.WriteString("sthir ")
	} else if ad.IsGlobal {
		out.WriteString("bishwo ")
	} else {
		out.WriteString("dhoro ")
	}
	names := make([]string, 0, len(ad.Names))
	for _, n := range ad.Names {
		names = append(names, n.Value)
	}
	out.WriteString("[")
	out.WriteString(strings.Join(names, ", "))
	out.WriteString("] = ")
	out.WriteString(ad.Source.String())
	out.WriteString(";")
	return out.String()
}

// ObjectDestructuringDeclaration represents: dhoro {x, y} = expr;
type ObjectDestructuringDeclaration struct {
	Token      lexer.Token // DHORO/STHIR/BISHWO token
	Keys       []string
	Names      []*Identifier
	Source     Expression
	IsConstant bool
	IsGlobal   bool
}

func (od *ObjectDestructuringDeclaration) statementNode()       {}
func (od *ObjectDestructuringDeclaration) TokenLiteral() string { return od.Token.Literal }
func (od *ObjectDestructuringDeclaration) String() string {
	var out bytes.Buffer
	if od.IsConstant {
		out.WriteString("sthir ")
	} else if od.IsGlobal {
		out.WriteString("bishwo ")
	} else {
		out.WriteString("dhoro ")
	}
	out.WriteString("{")
	out.WriteString(strings.Join(od.Keys, ", "))
	out.WriteString("} = ")
	out.WriteString(od.Source.String())
	out.WriteString(";")
	return out.String()
}

// ArrowParamList is an internal expression node used to parse (a, b) => ...
type ArrowParamList struct {
	Token  lexer.Token // LPAREN token
	Params []*Identifier
}

func (ap *ArrowParamList) expressionNode()      {}
func (ap *ArrowParamList) TokenLiteral() string { return ap.Token.Literal }
func (ap *ArrowParamList) String() string {
	names := make([]string, 0, len(ap.Params))
	for _, p := range ap.Params {
		names = append(names, p.Value)
	}
	return "(" + strings.Join(names, ", ") + ")"
}
