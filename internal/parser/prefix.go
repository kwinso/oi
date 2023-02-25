package parser

import "oilang/internal/ast"

func (p *Parser) parsePrefixExpression() (ast.Expression, *ParsingError) {
	exp := &ast.PrefixExpression{Token: p.curToken}

	p.nextToken()

	// Parse next expression with the prefix precedence
	v, err := p.parseExpression(PREFIX)
	if err != nil {
		return nil, err
	}

	exp.Operand = v
	return exp, nil
}
