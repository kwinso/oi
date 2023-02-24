package parser

import (
	"oilang/internal/lexer"
	"oilang/internal/parser/ast"
	"oilang/internal/token"
)

type ParsingError struct {
	Message string
	Token   token.Token
}

type Parser struct {
	l *lexer.Lexer

	errors ParsingError

	curToken  token.Token
	peekToken token.Token
}

// New Creates token parser based on lexer
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Set both current and peek tokens by reading twice
	p.nextToken()
	p.nextToken()

	return p
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
	case token.RETURN:
		return p.parseReturnStatement(), nil
	default:
		return nil, nil
	}
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

	return stmt, nil
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TODO: Value is not stored
	for !p.curTokenIs(token.NEWLINE) && !p.curTokenIs(token.EOF) {
		p.nextToken()
	}

	return stmt
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

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
