package parser

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/lexer"
)

func (p *Parser) parseArrayDestructuringDeclaration(token lexer.Token, isConstant, isGlobal bool) *ast.ArrayDestructuringDeclaration {
	stmt := &ast.ArrayDestructuringDeclaration{
		Token:      token,
		IsConstant: isConstant,
		IsGlobal:   isGlobal,
		Names:      []*ast.Identifier{},
	}

	// current token is '['
	for {
		p.nextToken()
		if p.curTokenIs(lexer.RBRACKET) {
			break
		}
		if !p.curTokenIs(lexer.IDENT) {
			p.errors = append(p.errors, "array destructuring expects identifiers")
			return nil
		}
		stmt.Names = append(stmt.Names, &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal})
		if p.peekTokenIs(lexer.COMMA) {
			p.nextToken()
			continue
		}
		if p.peekTokenIs(lexer.RBRACKET) {
			p.nextToken()
			break
		}
		p.errors = append(p.errors, "array destructuring expects ',' or ']'")
		return nil
	}

	if !p.expectPeek(lexer.ASSIGN) {
		return nil
	}
	p.nextToken()
	stmt.Source = p.parseExpression(LOWEST)
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseObjectDestructuringDeclaration(token lexer.Token, isConstant, isGlobal bool) *ast.ObjectDestructuringDeclaration {
	stmt := &ast.ObjectDestructuringDeclaration{
		Token:      token,
		IsConstant: isConstant,
		IsGlobal:   isGlobal,
		Keys:       []string{},
		Names:      []*ast.Identifier{},
	}

	if !p.parseObjectDestructuringBindings(stmt) {
		return nil
	}

	if !p.expectPeek(lexer.ASSIGN) {
		return nil
	}
	p.nextToken()
	stmt.Source = p.parseExpression(LOWEST)
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseObjectDestructuringBindings(stmt *ast.ObjectDestructuringDeclaration) bool {
	// current token is '{'
	for {
		p.nextToken()
		if p.curTokenIs(lexer.RBRACE) {
			return true
		}

		key, name, ok := p.parseObjectDestructuringPair()
		if !ok {
			return false
		}
		stmt.Keys = append(stmt.Keys, key)
		stmt.Names = append(stmt.Names, name)

		if p.peekTokenIs(lexer.COMMA) {
			p.nextToken()
			continue
		}
		if p.peekTokenIs(lexer.RBRACE) {
			p.nextToken()
			return true
		}
		p.errors = append(p.errors, "object destructuring expects ',' or '}'")
		return false
	}
}

func (p *Parser) parseObjectDestructuringPair() (string, *ast.Identifier, bool) {
	if !p.curTokenIs(lexer.IDENT) {
		p.errors = append(p.errors, "object destructuring expects identifier keys")
		return "", nil, false
	}

	key := p.curToken.Literal
	name := key
	token := p.curToken

	if p.peekTokenIs(lexer.COLON) {
		p.nextToken()
		if !p.expectPeek(lexer.IDENT) {
			return "", nil, false
		}
		name = p.curToken.Literal
		token = p.curToken
	}

	return key, &ast.Identifier{Token: token, Value: name}, true
}
