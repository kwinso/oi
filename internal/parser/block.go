package parser

import (
	"oilang/internal/ast"
	"oilang/internal/token"
)

func (p *Parser) parseBlockStatement() (*ast.BlockStatement, *ParsingError) {
	block := &ast.BlockStatement{Token: p.curToken}
	block.Statements = []ast.Statement{}

	p.nextToken()

	for !p.curTokenIs(token.RBRACE) && !p.curTokenIs(token.EOF) {
		stmt, err := p.parseStatement()
		if err != nil {
			return nil, err
		}

		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}

		p.nextToken()
	}

	// Check that block is closed properly
	if !p.curTokenIs(token.RBRACE) {
		return nil, p.createCurrentTokenError("expected } at the end of block")
	}

	return block, nil
}
