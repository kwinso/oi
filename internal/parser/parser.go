package parser

import (
	"errors"
	"oilang/internal/ast"
	"oilang/internal/lexer"
	"oilang/internal/token"
)

type ParsingError struct {
	Message string
	Token   token.Token
}

// Operator precedence levels
const (
	LOWEST = iota
	OR
	AND
	NOT
	EQUALS
	COMPARISON
	SUM
	PRODUCT
	PREFIX
	EXP
	CALL
)

// Generic function for parsing expressions for different token positions
type (
	// Parsing when token is encountered in prefix position(e.g. -, not, literals and identifiers)
	prefixParseFn func() (ast.Expression, error)
	// Parsing when token it's encountered in binary position (e.g. 5 + 5)
	infixParseFn func(ast.Expression) (ast.Expression, error)
)

type Parser struct {
	l *lexer.Lexer

	errors ParsingError

	curToken  token.Token
	peekToken token.Token

	prefixParsers map[token.TokenType]prefixParseFn
	infixParsers  map[token.TokenType]infixParseFn
}

// New Creates token parser based on lexer
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.setupParsers()

	// Set both current and peek tokens by reading twice
	p.nextToken()
	p.nextToken()

	return p
}

// All parser functions are located in the separate files
func (p *Parser) setupParsers() {
	p.prefixParsers = make(map[token.TokenType]prefixParseFn)
	p.prefixParsers[token.IDENT] = p.parseIdentifier
	p.prefixParsers[token.INT] = p.parseInt
	p.prefixParsers[token.FLOAT] = p.parseFloat
	p.prefixParsers[token.NOT] = p.parsePrefixExpression
	p.prefixParsers[token.MINUS] = p.parsePrefixExpression
}

func (p *Parser) registerPrefix(token token.TokenType, fn prefixParseFn) {
	p.prefixParsers[token] = fn
}

func (p *Parser) registerInfix(token token.TokenType, fn infixParseFn) {
	p.infixParsers[token] = fn
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Parse parsers tokens from lexer into AST
func (p *Parser) Parse() (*ast.Program, *ParsingError) {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// FIXME: add error handling with showing where was the error
	for p.curToken.Type != token.EOF {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program, nil
}

func (p *Parser) parseStatement() (ast.Statement, *ParsingError) {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	// TODO: Return statement should not appear outside of function body
	case token.RETURN:
		return p.parseReturnStatement(), nil
	default:
		if !p.isEndOfStatementToken(p.curToken) {
			return p.parseExpressionStatement()
		}
	}

	return nil, nil
}

// TODO: Allow not setting values
func (p *Parser) parseLetStatement() (*ast.LetStatement, *ParsingError) {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil, p.createPeekError("Identifier expected")
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(token.ASSIGN) {
		return nil, p.createPeekError("Assign operator expected")
	}

	// TODO: Value is not stored
	for !p.isEndOfStatementToken(p.curToken) {
		p.nextToken()
	}

	return stmt, nil
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TODO: Value is not stored
	for !p.isEndOfStatementToken(p.curToken) {
		p.nextToken()
	}

	return stmt
}

// TODO:
func (p *Parser) parseExpressionStatement() (*ast.ExpressionStatement, *ParsingError) {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	exp, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, &ParsingError{Token: p.curToken, Message: err.Error()}
	}
	stmt.Expression = exp

	if p.isEndOfStatementToken(p.peekToken) {
		p.nextToken()
	}

	return stmt, nil
}

func (p *Parser) parseExpression(precedence int) (ast.Expression, error) {
	prefix := p.prefixParsers[p.curToken.Type]
	if prefix == nil {
		return nil, errors.New("cannot parse this token")
	}

	return prefix()
}

func (p *Parser) createPeekError(msg string) *ParsingError {
	return &ParsingError{msg, p.peekToken}
}

// Peeks next token if it matches the supplied type and returns true
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	return false
}

func (p *Parser) isEndOfStatementToken(tok token.Token) bool {
	return tok.Type == token.SEMICOLON || tok.Type == token.EOF || tok.Type == token.NEWLINE
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
