package parser

import (
	"BanglaCode/ast"
	"BanglaCode/lexer"
	"fmt"
	"strconv"
)

// Operator precedence levels
const (
	_ int = iota
	LOWEST
	ASSIGN      // =, +=, -=, *=, /=
	OR          // ba (||)
	AND         // ebong (&&)
	EQUALS      // ==, !=
	LESSGREATER // <, >, <=, >=
	SUM         // +, -
	PRODUCT     // *, /, %
	PREFIX      // -x, !x, na x
	CALL        // function(x)
	INDEX       // array[index], obj.prop
)

var precedences = map[lexer.TokenType]int{
	lexer.ASSIGN:         ASSIGN,
	lexer.PLUS_ASSIGN:    ASSIGN,
	lexer.MINUS_ASSIGN:   ASSIGN,
	lexer.ASTERISK_ASSIGN: ASSIGN,
	lexer.SLASH_ASSIGN:   ASSIGN,
	lexer.BA:             OR,
	lexer.EBONG:          AND,
	lexer.EQ:             EQUALS,
	lexer.NOT_EQ:         EQUALS,
	lexer.LT:             LESSGREATER,
	lexer.GT:             LESSGREATER,
	lexer.LTE:            LESSGREATER,
	lexer.GTE:            LESSGREATER,
	lexer.PLUS:           SUM,
	lexer.MINUS:          SUM,
	lexer.ASTERISK:       PRODUCT,
	lexer.SLASH:          PRODUCT,
	lexer.PERCENT:        PRODUCT,
	lexer.LPAREN:         CALL,
	lexer.LBRACKET:       INDEX,
	lexer.DOT:            INDEX,
}

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

// Parser represents the parser
type Parser struct {
	l      *lexer.Lexer
	errors []string

	curToken  lexer.Token
	peekToken lexer.Token

	prefixParseFns map[lexer.TokenType]prefixParseFn
	infixParseFns  map[lexer.TokenType]infixParseFn
}

// New creates a new parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// Register prefix parse functions
	p.prefixParseFns = make(map[lexer.TokenType]prefixParseFn)
	p.registerPrefix(lexer.IDENT, p.parseIdentifier)
	p.registerPrefix(lexer.NUMBER, p.parseNumberLiteral)
	p.registerPrefix(lexer.STRING, p.parseStringLiteral)
	p.registerPrefix(lexer.SOTTI, p.parseBooleanLiteral)
	p.registerPrefix(lexer.MITTHA, p.parseBooleanLiteral)
	p.registerPrefix(lexer.KHALI, p.parseNullLiteral)
	p.registerPrefix(lexer.BANG, p.parseUnaryExpression)
	p.registerPrefix(lexer.MINUS, p.parseUnaryExpression)
	p.registerPrefix(lexer.NA, p.parseUnaryExpression)
	p.registerPrefix(lexer.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(lexer.LBRACKET, p.parseArrayLiteral)
	p.registerPrefix(lexer.LBRACE, p.parseMapLiteral)
	p.registerPrefix(lexer.KAJ, p.parseFunctionLiteral)
	p.registerPrefix(lexer.NOTUN, p.parseNewExpression)

	// Register infix parse functions
	p.infixParseFns = make(map[lexer.TokenType]infixParseFn)
	p.registerInfix(lexer.PLUS, p.parseBinaryExpression)
	p.registerInfix(lexer.MINUS, p.parseBinaryExpression)
	p.registerInfix(lexer.ASTERISK, p.parseBinaryExpression)
	p.registerInfix(lexer.SLASH, p.parseBinaryExpression)
	p.registerInfix(lexer.PERCENT, p.parseBinaryExpression)
	p.registerInfix(lexer.EQ, p.parseBinaryExpression)
	p.registerInfix(lexer.NOT_EQ, p.parseBinaryExpression)
	p.registerInfix(lexer.LT, p.parseBinaryExpression)
	p.registerInfix(lexer.GT, p.parseBinaryExpression)
	p.registerInfix(lexer.LTE, p.parseBinaryExpression)
	p.registerInfix(lexer.GTE, p.parseBinaryExpression)
	p.registerInfix(lexer.EBONG, p.parseBinaryExpression)
	p.registerInfix(lexer.BA, p.parseBinaryExpression)
	p.registerInfix(lexer.ASSIGN, p.parseAssignmentExpression)
	p.registerInfix(lexer.PLUS_ASSIGN, p.parseAssignmentExpression)
	p.registerInfix(lexer.MINUS_ASSIGN, p.parseAssignmentExpression)
	p.registerInfix(lexer.ASTERISK_ASSIGN, p.parseAssignmentExpression)
	p.registerInfix(lexer.SLASH_ASSIGN, p.parseAssignmentExpression)
	p.registerInfix(lexer.LPAREN, p.parseCallExpression)
	p.registerInfix(lexer.LBRACKET, p.parseMemberExpression)
	p.registerInfix(lexer.DOT, p.parseMemberExpression)

	// Read two tokens to initialize curToken and peekToken
	p.nextToken()
	p.nextToken()

	return p
}

// Errors returns parser errors
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t lexer.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead at line %d, column %d",
		t, p.peekToken.Type, p.peekToken.Line, p.peekToken.Column)
	p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t lexer.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t lexer.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t lexer.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) registerPrefix(tokenType lexer.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType lexer.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

// ParseProgram parses the entire program
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(lexer.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case lexer.DHORO:
		return p.parseVariableDeclaration()
	case lexer.JODI:
		return p.parseIfStatement()
	case lexer.JOTOKKHON:
		return p.parseWhileStatement()
	case lexer.GHURIYE:
		return p.parseForStatement()
	case lexer.FERAO:
		return p.parseReturnStatement()
	case lexer.CLASS:
		return p.parseClassDeclaration()
	case lexer.THAMO:
		return p.parseBreakStatement()
	case lexer.CHHARO:
		return p.parseContinueStatement()
	case lexer.ANO:
		return p.parseImportStatement()
	case lexer.PATHAO:
		return p.parseExportStatement()
	case lexer.CHESTA:
		return p.parseTryCatchStatement()
	case lexer.FELO:
		return p.parseThrowStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseVariableDeclaration() *ast.VariableDeclaration {
	stmt := &ast.VariableDeclaration{Token: p.curToken}

	if !p.expectPeek(lexer.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(lexer.ASSIGN) {
		return nil
	}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseIfStatement() *ast.IfStatement {
	stmt := &ast.IfStatement{Token: p.curToken}

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	stmt.Consequence = p.parseBlockStatement()

	if p.peekTokenIs(lexer.NAHOLE) {
		p.nextToken()

		// Check for "nahole jodi" (else if)
		if p.peekTokenIs(lexer.JODI) {
			p.nextToken()
			// Parse else if as a nested if statement
			ifStmt := p.parseIfStatement()
			stmt.Alternative = &ast.BlockStatement{
				Token:      p.curToken,
				Statements: []ast.Statement{ifStmt},
			}
		} else {
			if !p.expectPeek(lexer.LBRACE) {
				return nil
			}
			stmt.Alternative = p.parseBlockStatement()
		}
	}

	return stmt
}

func (p *Parser) parseWhileStatement() *ast.WhileStatement {
	stmt := &ast.WhileStatement{Token: p.curToken}

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	p.nextToken()
	stmt.Condition = p.parseExpression(LOWEST)

	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

func (p *Parser) parseForStatement() *ast.ForStatement {
	stmt := &ast.ForStatement{Token: p.curToken}

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	p.nextToken()

	// Parse init statement
	if !p.curTokenIs(lexer.SEMICOLON) {
		stmt.Init = p.parseStatement()
	}

	if !p.curTokenIs(lexer.SEMICOLON) {
		if !p.expectPeek(lexer.SEMICOLON) {
			return nil
		}
	}

	p.nextToken()

	// Parse condition
	if !p.curTokenIs(lexer.SEMICOLON) {
		stmt.Condition = p.parseExpression(LOWEST)
	}

	if !p.expectPeek(lexer.SEMICOLON) {
		return nil
	}

	p.nextToken()

	// Parse update
	if !p.curTokenIs(lexer.RPAREN) {
		stmt.Update = p.parseExpression(LOWEST)
	}

	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	stmt.Body = p.parseBlockStatement()

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	if !p.curTokenIs(lexer.SEMICOLON) && !p.curTokenIs(lexer.RBRACE) {
		stmt.ReturnValue = p.parseExpression(LOWEST)
	}

	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseClassDeclaration() *ast.ClassDeclaration {
	stmt := &ast.ClassDeclaration{Token: p.curToken}

	if !p.expectPeek(lexer.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	p.nextToken()

	stmt.Methods = []*ast.FunctionLiteral{}

	for !p.curTokenIs(lexer.RBRACE) && !p.curTokenIs(lexer.EOF) {
		if p.curTokenIs(lexer.KAJ) {
			method := p.parseFunctionLiteral()
			if fn, ok := method.(*ast.FunctionLiteral); ok {
				stmt.Methods = append(stmt.Methods, fn)
			}
		}
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseBreakStatement() *ast.BreakStatement {
	stmt := &ast.BreakStatement{Token: p.curToken}
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseContinueStatement() *ast.ContinueStatement {
	stmt := &ast.ContinueStatement{Token: p.curToken}
	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}
	return stmt
}

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(lexer.RBRACE) && !p.curTokenIs(lexer.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextToken()
	}

	return block
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(lexer.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}

func (p *Parser) noPrefixParseFnError(t lexer.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found at line %d, column %d",
		t, p.curToken.Line, p.curToken.Column)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseNumberLiteral() ast.Expression {
	lit := &ast.NumberLiteral{Token: p.curToken}

	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as number", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}

	lit.Value = value
	return lit
}

func (p *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseBooleanLiteral() ast.Expression {
	return &ast.BooleanLiteral{Token: p.curToken, Value: p.curTokenIs(lexer.SOTTI)}
}

func (p *Parser) parseNullLiteral() ast.Expression {
	return &ast.NullLiteral{Token: p.curToken}
}

func (p *Parser) parseUnaryExpression() ast.Expression {
	expression := &ast.UnaryExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	return exp
}

func (p *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: p.curToken}
	array.Elements = p.parseExpressionList(lexer.RBRACKET)
	return array
}

func (p *Parser) parseMapLiteral() ast.Expression {
	mapLit := &ast.MapLiteral{Token: p.curToken}
	mapLit.Pairs = make(map[ast.Expression]ast.Expression)

	for !p.peekTokenIs(lexer.RBRACE) {
		p.nextToken()
		key := p.parseExpression(LOWEST)

		if !p.expectPeek(lexer.COLON) {
			return nil
		}

		p.nextToken()
		value := p.parseExpression(LOWEST)

		mapLit.Pairs[key] = value

		if !p.peekTokenIs(lexer.RBRACE) && !p.expectPeek(lexer.COMMA) {
			return nil
		}
	}

	if !p.expectPeek(lexer.RBRACE) {
		return nil
	}

	return mapLit
}

func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.curToken}

	// Check if function has a name
	if p.peekTokenIs(lexer.IDENT) {
		p.nextToken()
		lit.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	if !p.expectPeek(lexer.LPAREN) {
		return nil
	}

	lit.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if p.peekTokenIs(lexer.RPAREN) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	identifiers = append(identifiers, ident)

	for p.peekTokenIs(lexer.COMMA) {
		p.nextToken()
		p.nextToken()
		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		identifiers = append(identifiers, ident)
	}

	if !p.expectPeek(lexer.RPAREN) {
		return nil
	}

	return identifiers
}

func (p *Parser) parseNewExpression() ast.Expression {
	exp := &ast.NewExpression{Token: p.curToken}

	p.nextToken()
	exp.Class = p.parseExpression(CALL)

	if p.peekTokenIs(lexer.LPAREN) {
		p.nextToken()
		exp.Arguments = p.parseExpressionList(lexer.RPAREN)
	}

	return exp
}

func (p *Parser) parseBinaryExpression(left ast.Expression) ast.Expression {
	expression := &ast.BinaryExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)

	return expression
}

func (p *Parser) parseAssignmentExpression(left ast.Expression) ast.Expression {
	expression := &ast.AssignmentExpression{
		Token:    p.curToken,
		Name:     left,
		Operator: p.curToken.Literal,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Value = p.parseExpression(precedence)

	return expression
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseExpressionList(lexer.RPAREN)
	return exp
}

func (p *Parser) parseMemberExpression(left ast.Expression) ast.Expression {
	exp := &ast.MemberExpression{
		Token:  p.curToken,
		Object: left,
	}

	if p.curTokenIs(lexer.LBRACKET) {
		exp.Computed = true
		p.nextToken()
		exp.Property = p.parseExpression(LOWEST)
		if !p.expectPeek(lexer.RBRACKET) {
			return nil
		}
	} else {
		exp.Computed = false
		if !p.expectPeek(lexer.IDENT) {
			return nil
		}
		exp.Property = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	return exp
}

func (p *Parser) parseExpressionList(end lexer.TokenType) []ast.Expression {
	list := []ast.Expression{}

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}

	p.nextToken()
	list = append(list, p.parseExpression(LOWEST))

	for p.peekTokenIs(lexer.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}

func (p *Parser) parseImportStatement() *ast.ImportStatement {
	stmt := &ast.ImportStatement{Token: p.curToken}

	if !p.expectPeek(lexer.STRING) {
		return nil
	}

	stmt.Path = &ast.StringLiteral{Token: p.curToken, Value: p.curToken.Literal}

	// Check for optional "as alias" syntax
	if p.peekTokenIs(lexer.IDENT) && p.peekToken.Literal == "as" {
		p.nextToken() // move to "as"
		if !p.expectPeek(lexer.IDENT) {
			return nil
		}
		stmt.Alias = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExportStatement() *ast.ExportStatement {
	stmt := &ast.ExportStatement{Token: p.curToken}

	p.nextToken()

	// Parse the statement being exported
	stmt.Statement = p.parseStatement()

	return stmt
}

func (p *Parser) parseTryCatchStatement() *ast.TryCatchStatement {
	stmt := &ast.TryCatchStatement{Token: p.curToken}

	if !p.expectPeek(lexer.LBRACE) {
		return nil
	}

	stmt.TryBlock = p.parseBlockStatement()

	// Parse catch block
	if p.peekTokenIs(lexer.DHORO_BHUL) {
		p.nextToken()

		if p.peekTokenIs(lexer.LPAREN) {
			p.nextToken()
			if !p.expectPeek(lexer.IDENT) {
				return nil
			}
			stmt.CatchParam = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
			if !p.expectPeek(lexer.RPAREN) {
				return nil
			}
		}

		if !p.expectPeek(lexer.LBRACE) {
			return nil
		}

		stmt.CatchBlock = p.parseBlockStatement()
	}

	// Parse finally block
	if p.peekTokenIs(lexer.SHESH) {
		p.nextToken()

		if !p.expectPeek(lexer.LBRACE) {
			return nil
		}

		stmt.FinallyBlock = p.parseBlockStatement()
	}

	return stmt
}

func (p *Parser) parseThrowStatement() *ast.ThrowStatement {
	stmt := &ast.ThrowStatement{Token: p.curToken}

	p.nextToken()

	stmt.Value = p.parseExpression(LOWEST)

	if p.peekTokenIs(lexer.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
