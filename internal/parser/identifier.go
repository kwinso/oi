package parser

import (
	"oilang/internal/ast"
)

func (p *Parser) parseIdentifier() (ast.Expression, *ParsingError) {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}, nil
}
