package parser

import "BanglaCode/src/lexer"

// Operator precedence levels
const (
	_ int = iota
	LOWEST
	ASSIGN      // =, +=, -=, *=, /=
	ARROWP      // =>
	OR          // ba (||)
	AND         // ebong (&&)
	INOP        // in, instanceof
	EQUALS      // ==, !=
	LESSGREATER // <, >, <=, >=
	SUM         // +, -
	PRODUCT     // *, /, %
	PREFIX      // -x, !x, na x
	CALL        // function(x)
	INDEX       // array[index], obj.prop
)

// precedences maps token types to their precedence levels
var precedences = map[lexer.TokenType]int{
	lexer.ASSIGN:          ASSIGN,
	lexer.PLUS_ASSIGN:     ASSIGN,
	lexer.MINUS_ASSIGN:    ASSIGN,
	lexer.ASTERISK_ASSIGN: ASSIGN,
	lexer.SLASH_ASSIGN:    ASSIGN,
	lexer.ARROW:           ARROWP,
	lexer.BA:              OR,
	lexer.EBONG:           AND,
	lexer.IN:              INOP,
	lexer.INSTANCEOF:      INOP,
	lexer.EQ:              EQUALS,
	lexer.NOT_EQ:          EQUALS,
	lexer.LT:              LESSGREATER,
	lexer.GT:              LESSGREATER,
	lexer.LTE:             LESSGREATER,
	lexer.GTE:             LESSGREATER,
	lexer.PLUS:            SUM,
	lexer.MINUS:           SUM,
	lexer.ASTERISK:        PRODUCT,
	lexer.SLASH:           PRODUCT,
	lexer.PERCENT:         PRODUCT,
	lexer.LPAREN:          CALL,
	lexer.LBRACKET:        INDEX,
	lexer.DOT:             INDEX,
}

// peekPrecedence returns the precedence of the next token
func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

// curPrecedence returns the precedence of the current token
func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}
