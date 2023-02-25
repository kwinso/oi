package parser

import (
	"oilang/internal/ast"
	"strconv"
)

// parseInt converts current token to 64-bit integer
func (p *Parser) parseInt() (ast.Expression, *ParsingError) {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	v, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		return nil, p.createCurrentTokenError("unable to parse integer")
	}

	lit.Value = v

	return lit, nil
}

// parseFloat converts current token to 64-bit float
func (p *Parser) parseFloat() (ast.Expression, *ParsingError) {
	lit := &ast.FloatLiteral{Token: p.curToken}

	v, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		return nil, p.createCurrentTokenError("unable to parse float number")
	}

	lit.Value = v

	return lit, nil
}
