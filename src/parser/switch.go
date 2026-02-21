package parser

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/lexer"
)

// parseSwitchStatement parses "bikolpo (expression) { khetre value { ... } manchito { ... } }"
func (p *Parser) parseSwitchStatement() *ast.SwitchStatement {
	stmt := &ast.SwitchStatement{Token: p.curToken}
	if !p.parseSwitchHeader(stmt) {
		return nil
	}
	stmt.Cases = []*ast.CaseClause{}
	p.nextToken()
	for !p.curTokenIs(lexer.RBRACE) && !p.curTokenIs(lexer.EOF) {
		if p.curTokenIs(lexer.KHETRE) {
			caseClause := p.parseCaseClause()
			if caseClause != nil {
				stmt.Cases = append(stmt.Cases, caseClause)
			}
			p.nextToken()
			continue
		}
		if p.curTokenIs(lexer.MANCHITO) {
			if !p.parseDefaultClause(stmt) {
				return nil
			}
			continue
		}
		return nil
	}
	return stmt
}

func (p *Parser) parseSwitchHeader(stmt *ast.SwitchStatement) bool {
	if !p.expectPeek(lexer.LPAREN) {
		return false
	}
	p.nextToken()
	stmt.Expr = p.parseExpression(LOWEST)
	if !p.expectPeek(lexer.RPAREN) {
		return false
	}
	return p.expectPeek(lexer.LBRACE)
}

func (p *Parser) parseDefaultClause(stmt *ast.SwitchStatement) bool {
	if !p.expectPeek(lexer.LBRACE) {
		return false
	}
	stmt.Default = p.parseBlockStatement()
	p.nextToken()
	return true
}

// parseCaseClause parses a single "khetre value { ... }" clause
func (p *Parser) parseCaseClause() *ast.CaseClause {
	clause := &ast.CaseClause{Token: p.curToken}
	p.nextToken()
	clause.Value = p.parseExpression(LOWEST)
	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}
	clause.Body = p.parseBlockStatement()
	return clause
}
