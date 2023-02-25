package parser

import (
	"oilang/internal/ast"
	"oilang/internal/token"
)

func (p *Parser) parseIfExpression() (ast.Expression, *ParsingError) {
	exp := &ast.IfExpression{Token: p.curToken}
	p.nextToken()

	cond, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}
	exp.Condition = cond

	if !p.tryPeek(token.LBRACE) {
		return nil, p.createPeekError("expected { for main if branch")
	}

	exp.Consequnce, err = p.parseBlockStatement()
	if err != nil {
		return nil, err
	}

	if p.peekTokenIs(token.ELSE) {
		p.nextToken()

		if !p.tryPeek(token.LBRACE) {
			return nil, p.createPeekError("expected { for else branch")
		}

		exp.Alternative, err = p.parseBlockStatement()
		if err != nil {
			return nil, err
		}
	}

	return exp, nil
}
