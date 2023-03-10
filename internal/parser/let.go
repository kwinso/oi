package parser

import (
	"oilang/internal/ast"
	"oilang/internal/token"
)

// TODO: Allow not setting values
// parseLetStatement expects peek token to be an identifier followed by ASSIGN token
//
// After this, it assign statement's value to expression after the ASSIGN token
func (p *Parser) parseLetStatement() (*ast.LetStatement, *ParsingError) {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.tryPeek(token.IDENT) {
		return nil, p.createPeekError("Identifier expected")
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.tryPeek(token.ASSIGN) {
		return nil, p.createPeekError("Assign operator expected")
	}

	p.nextToken()

	val, err := p.parseExpression(LOWEST)

	if err != nil {
		return nil, err
	}

	stmt.Value = val

	for !p.isEndOfStatementToken(p.curToken) {
		p.nextToken()
	}

	return stmt, nil
}
