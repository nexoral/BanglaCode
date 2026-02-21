package parser

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/lexer"
)

// parseDoWhileStatement parses: do { ... } jotokkhon (condition);
func (p *Parser) parseDoWhileStatement() *ast.DoWhileStatement {
	stmt := &ast.DoWhileStatement{Token: p.curToken}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	stmt.Body = p.parseBlockStatement()

	if !p.expectPeek(lexer.JOTOKKHON) {
		return nil
	}
	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseDeleteExpression parses: delete targetExpression
func (p *Parser) parseDeleteExpression() ast.Expression {
	exp := &ast.DeleteExpression{Token: p.curToken}
	p.nextToken()
	exp.Target = p.parseExpression(PREFIX)
	return exp
}
