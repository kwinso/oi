package parser

import "oilang/internal/ast"

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	// TODO: Value is not stored
	for !p.isEndOfStatementToken(p.curToken) {
		p.nextToken()
	}

	return stmt
}
