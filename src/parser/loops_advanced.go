package parser

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/lexer"
)

func (p *Parser) parseForInOrForOf(forToken lexer.Token) ast.Statement {
	// Support: ghuriye (item of iterable) { ... }
	// Support: ghuriye (key in object) { ... }
	if !p.curTokenIs(lexer.IDENT) {
		return nil
	}
	if !(p.peekTokenIs(lexer.OF) || p.peekTokenIs(lexer.IN)) {
		return nil
	}
	varName := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	p.nextToken()
	loopKind := p.curToken.Type

	p.nextToken()
	iterableOrObject := p.parseExpression(LOWEST)
	if iterableOrObject == nil {
		return nil
	}

	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	body := p.parseBlockStatement()

	if loopKind == lexer.OF {
		return &ast.ForOfStatement{
			Token:    forToken,
			VarName:  varName,
			Iterable: iterableOrObject,
			Body:     body,
		}
	}

	return &ast.ForInStatement{
		Token:   forToken,
		VarName: varName,
		Object:  iterableOrObject,
		Body:    body,
	}
}
