package parser

import "oilang/internal/ast"

func (p *Parser) parseInfixExpression(left ast.Expression) (ast.Expression, error) {
	exp := &ast.InfixExpression{Token: p.curToken, Left: left}
	precedence := p.curPrecedence()

	p.nextToken()

	right, err := p.parseExpression(precedence)
	if err != nil {
		return nil, err
	}

	exp.Right = right
	return exp, nil
}
