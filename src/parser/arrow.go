package parser

import (
	"BanglaCode/src/ast"
	"BanglaCode/src/lexer"
)

// parseArrowFunctionExpression parses: x => expr  OR  x => { ... }
func (p *Parser) parseArrowFunctionExpression(left ast.Expression) ast.Expression {
	params := []*ast.Identifier{}

	switch l := left.(type) {
	case *ast.Identifier:
		params = append(params, l)
	case *ast.ArrowParamList:
		params = append(params, l.Params...)
	default:
		p.errors = append(p.errors, "invalid arrow function parameters")
		return nil
	}

	fn := &ast.FunctionLiteral{
		Token:      p.curToken, // ARROW token
		Parameters: params,
	}

	p.nextToken()

	if p.curTokenIs(lexer.LBRACE) {
		fn.Body = p.parseBlockStatement()
		return fn
	}

	// Expression body has implicit return
	bodyExpr := p.parseExpression(LOWEST)
	fn.Body = &ast.BlockStatement{
		Token: lexer.NewToken(lexer.LBRACE, "{", p.curToken.Line, p.curToken.Column),
		Statements: []ast.Statement{
			&ast.ReturnStatement{
				Token:       lexer.NewToken(lexer.FERAO, "ferao", p.curToken.Line, p.curToken.Column),
				ReturnValue: bodyExpr,
			},
		},
	}

	return fn
}
