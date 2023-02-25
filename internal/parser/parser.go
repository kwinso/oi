package parser

import (
	"oilang/internal/ast"
	"oilang/internal/lexer"
	"oilang/internal/token"
)

// TODO: Implement converting to string with errored input line

// ParsingError represents an error that occuring during parsing, as well as token that has caused an error
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

// Maps each token that could appear in infix position to its precedence level
var precedences = map[token.TokenType]int{
	token.EQ:  EQUALS,
	token.NEQ: EQUALS,
	token.OR:  OR,
	token.AND: AND,
	token.NOT: NOT,

	token.LT:  COMPARISON,
	token.GT:  COMPARISON,
	token.LTE: COMPARISON,
	token.GTE: COMPARISON,

	token.PLUS:     SUM,
	token.MINUS:    SUM,
	token.DIVIDE:   PRODUCT,
	token.MULTIPLY: PRODUCT,
	token.POWER:    EXP,
}

// Generic function for parsing expressions for different token positions
type (
	// Parsing when token is encountered in prefix position(e.g. -, not, literals and identifiers)
	prefixParseFn func() (ast.Expression, *ParsingError)
	// Parsing when token it's encountered in binary position (e.g. 5 + 5)
	infixParseFn func(ast.Expression) (ast.Expression, *ParsingError)
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

// Parse Goes through lexical tokens and turns them into AST
func (p *Parser) Parse() (*ast.Program, *ParsingError) {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

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

// All parser functions are located in the separate files
//
// Note: all parsing function should stop advancing tokens on the last token that belongs to a given expression
func (p *Parser) setupParsers() {
	p.prefixParsers = make(map[token.TokenType]prefixParseFn)
	p.registerPrefixParser(token.IDENT, p.parseIdentifier)
	p.registerPrefixParser(token.INT, p.parseInt)
	p.registerPrefixParser(token.FLOAT, p.parseFloat)
	p.registerPrefixParser(token.NOT, p.parsePrefixExpression)
	p.registerPrefixParser(token.MINUS, p.parsePrefixExpression)

	p.infixParsers = make(map[token.TokenType]infixParseFn)
	for k := range precedences {
		p.registerInfixParser(k, p.parseInfixExpression)
	}
}

func (p *Parser) registerPrefixParser(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParsers[tokenType] = fn
}

func (p *Parser) registerInfixParser(tokenType token.TokenType, fn infixParseFn) {
	p.infixParsers[tokenType] = fn
}

// Steps one token forward, advancing current and peek tokens
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// Decides which function to call to turn given token into program statement
func (p *Parser) parseStatement() (ast.Statement, *ParsingError) {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	// TODO: Return statement should not appear outside of function body
	case token.RETURN:
		return p.parseReturnStatement(), nil
	case token.NEWLINE, token.EOF:
		break
	default:
		return p.parseExpressionStatement()
	}

	return nil, nil
}

// Creates and error for the peek token
func (p *Parser) createPeekError(msg string) *ParsingError {
	return &ParsingError{msg, p.peekToken}
}

// Peeks next token if it matches the supplied type and returns whether it matched
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	}

	return false
}

// Returns precedence of the supplied token
func (p *Parser) tokPrecedence(tok token.Token) int {
	if p, ok := precedences[tok.Type]; ok {
		return p
	}

	return LOWEST
}

// Returns precedence of the peek token
func (p *Parser) peekPrecedence() int {
	return p.tokPrecedence(p.peekToken)
}

// Returns precedence of the current token
func (p *Parser) curPrecedence() int {
	return p.tokPrecedence(p.curToken)
}

// Is token considered as an end of statement
func (p *Parser) isEndOfStatementToken(tok token.Token) bool {
	return tok.Type == token.SEMICOLON || tok.Type == token.EOF || tok.Type == token.NEWLINE
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}
