package parser

import "oilang/internal/ast"

func (p *Parser) parsePrefixExpression() (ast.Expression, error) {
	exp := &ast.PrefixExpression{Token: p.curToken}

	p.nextToken()

	v, err := p.parseExpression(PREFIX)
	if err != nil {
		return nil, err
	}

	exp.Operand = v
	return exp, nil
}
