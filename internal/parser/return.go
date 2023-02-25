package parser

import (
	"oilang/internal/ast"
	"oilang/internal/token"
)

func (p *Parser) parseReturnStatement() (*ast.ReturnStatement, *ParsingError) {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	if !p.isEndOfStatementToken(p.peekToken) && !p.peekTokenIs(token.RBRACE) {
		p.nextToken()

		ret, err := p.parseExpression(LOWEST)
		if err != nil {
			return nil, err
		}
		stmt.ReturnValue = ret

	}
	if p.isEndOfStatementToken(p.peekToken) {
		p.nextToken()
	}

	return stmt, nil
}
