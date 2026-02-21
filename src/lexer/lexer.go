package lexer

import (
	"unicode"
)

// Lexer represents the lexical analyzer
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
	line         int  // current line number
	column       int  // current column number
}

// New creates a new Lexer instance
func New(input string) *Lexer {
	l := &Lexer{
		input:  input,
		line:   1,
		column: 0,
	}
	l.readChar()
	return l
}

// readChar advances the lexer position and updates current character
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0 // ASCII code for "NUL" (end of input)
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
	l.column++

	// Track newlines for error reporting
	if l.ch == '\n' {
		l.line++
		l.column = 0
	}
}

// peekChar returns the next character without advancing position
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

// NextToken returns the next token from the input
func (l *Lexer) NextToken() Token {
	l.skipWhitespace()
	if l.consumeComment() {
		return l.NextToken()
	}
	if tok, ok := l.readStringOrTemplateToken(); ok {
		return tok
	}
	if tok, ok := l.readIdentifierOrNumberToken(); ok {
		return tok
	}
	tok, advance := l.readSymbolToken()
	if advance {
		l.readChar()
	}
	return tok
}

func (l *Lexer) consumeComment() bool {
	if l.ch != '/' || l.peekChar() != '/' {
		return false
	}
	l.skipComment()
	return true
}

func (l *Lexer) readStringOrTemplateToken() (Token, bool) {
	switch l.ch {
	case '"':
		return l.readQuotedToken('"'), true
	case '\'':
		return l.readQuotedToken('\''), true
	case '`':
		return l.readTemplateToken(), true
	default:
		return Token{}, false
	}
}

func (l *Lexer) readQuotedToken(quote byte) Token {
	tok := Token{Type: STRING, Line: l.line, Column: l.column}
	tok.Literal = l.readString(quote)
	l.readChar()
	return tok
}

func (l *Lexer) readTemplateToken() Token {
	tok := Token{Type: TEMPLATE, Line: l.line, Column: l.column}
	tok.Literal = l.readTemplate()
	l.readChar()
	return tok
}

func (l *Lexer) readIdentifierOrNumberToken() (Token, bool) {
	if isLetter(l.ch) {
		tok := Token{Line: l.line, Column: l.column}
		tok.Literal = l.readIdentifier()
		tok.Type = LookupIdent(tok.Literal)
		return tok, true
	}
	if isDigit(l.ch) {
		tok := Token{Type: NUMBER, Line: l.line, Column: l.column}
		tok.Literal = l.readNumber()
		return tok, true
	}
	return Token{}, false
}

func (l *Lexer) readSymbolToken() (Token, bool) {
	if tok, ok := l.readTwoCharOperator(); ok {
		return tok, true
	}
	switch l.ch {
	case '.':
		return l.readDotToken()
	case 0:
		return NewToken(EOF, "", l.line, l.column), false
	case ',', ';', ':', '(', ')', '{', '}', '[', ']', '%':
		return NewToken(singleCharTokenType(l.ch), string(l.ch), l.line, l.column), true
	default:
		return NewToken(ILLEGAL, string(l.ch), l.line, l.column), true
	}
}

func (l *Lexer) readTwoCharOperator() (Token, bool) {
	switch l.ch {
	case '=':
		if l.peekChar() == '>' {
			return l.makeTwoCharToken(ARROW), true
		}
		if l.peekChar() == '=' {
			return l.makeTwoCharToken(EQ), true
		}
		return NewToken(ASSIGN, string(l.ch), l.line, l.column), true
	case '+':
		if l.peekChar() == '=' {
			return l.makeTwoCharToken(PLUS_ASSIGN), true
		}
		return NewToken(PLUS, string(l.ch), l.line, l.column), true
	case '-':
		if l.peekChar() == '=' {
			return l.makeTwoCharToken(MINUS_ASSIGN), true
		}
		return NewToken(MINUS, string(l.ch), l.line, l.column), true
	case '*':
		if l.peekChar() == '=' {
			return l.makeTwoCharToken(ASTERISK_ASSIGN), true
		}
		return NewToken(ASTERISK, string(l.ch), l.line, l.column), true
	case '/':
		if l.peekChar() == '=' {
			return l.makeTwoCharToken(SLASH_ASSIGN), true
		}
		return NewToken(SLASH, string(l.ch), l.line, l.column), true
	case '!':
		if l.peekChar() == '=' {
			return l.makeTwoCharToken(NOT_EQ), true
		}
		return NewToken(BANG, string(l.ch), l.line, l.column), true
	case '<':
		if l.peekChar() == '=' {
			return l.makeTwoCharToken(LTE), true
		}
		return NewToken(LT, string(l.ch), l.line, l.column), true
	case '>':
		if l.peekChar() == '=' {
			return l.makeTwoCharToken(GTE), true
		}
		return NewToken(GT, string(l.ch), l.line, l.column), true
	default:
		return Token{}, false
	}
}

func (l *Lexer) makeTwoCharToken(tokenType TokenType) Token {
	ch := l.ch
	line := l.line
	column := l.column
	l.readChar()
	return NewToken(tokenType, string(ch)+string(l.ch), line, column)
}

func (l *Lexer) readDotToken() (Token, bool) {
	if l.peekChar() == '.' && l.readPosition+1 < len(l.input) && l.input[l.readPosition+1] == '.' {
		line := l.line
		column := l.column
		l.readChar()
		l.readChar()
		return NewToken(DOTDOTDOT, "...", line, column), true
	}
	return NewToken(DOT, string(l.ch), l.line, l.column), true
}

func singleCharTokenType(ch byte) TokenType {
	switch ch {
	case ',':
		return COMMA
	case ';':
		return SEMICOLON
	case ':':
		return COLON
	case '(':
		return LPAREN
	case ')':
		return RPAREN
	case '{':
		return LBRACE
	case '}':
		return RBRACE
	case '[':
		return LBRACKET
	case ']':
		return RBRACKET
	default:
		return PERCENT
	}
}

// readIdentifier reads an identifier (variable name, keyword, etc.)
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// readNumber reads a numeric literal (integer or float)
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	// Handle decimal numbers
	if l.ch == '.' && isDigit(l.peekChar()) {
		l.readChar() // consume '.'
		for isDigit(l.ch) {
			l.readChar()
		}
	}

	return l.input[position:l.position]
}

// readString reads a string literal
func (l *Lexer) readString(quote byte) string {
	position := l.position + 1 // skip opening quote
	for {
		l.readChar()
		if l.ch == quote || l.ch == 0 {
			break
		}
		// Handle escape sequences
		if l.ch == '\\' {
			l.readChar() // skip escaped character
		}
	}
	str := l.input[position:l.position]
	return str
}

// readTemplate reads a template literal with ${expression} interpolation
func (l *Lexer) readTemplate() string {
	position := l.position + 1 // skip opening backtick
	braceDepth := 0

	for {
		l.readChar()
		if l.ch == 0 {
			break // end of input
		}

		// Track brace depth for ${...} expressions
		if l.ch == '$' && l.peekChar() == '{' {
			l.readChar() // consume '{'
			braceDepth++
		} else if l.ch == '{' && braceDepth > 0 {
			braceDepth++
		} else if l.ch == '}' && braceDepth > 0 {
			braceDepth--
		} else if l.ch == '`' && braceDepth == 0 {
			break // found closing backtick (not in expression)
		}

		// Handle escape sequences
		if l.ch == '\\' {
			l.readChar() // skip escaped character
		}
	}

	str := l.input[position:l.position]
	return str
}

// skipWhitespace skips whitespace characters
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// skipComment skips single-line comments
func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

// isLetter checks if a character is a letter or underscore
func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

// isDigit checks if a character is a digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
