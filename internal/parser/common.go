package parser

// #####################
// This package contains some basic parsers that are too small to put in a different file
// #####################

import (
	"oilang/internal/ast"
	"oilang/internal/token"
)

func (p *Parser) parseIdentifier() (ast.Expression, *ParsingError) {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}, nil
}

func (p *Parser) parseBool() (ast.Expression, *ParsingError) {
	return &ast.BoolExpression{Token: p.curToken, Value: p.curTokenIs(token.TRUE)}, nil
}

func (p *Parser) createPrefixParserWithPrecedence(precedence int) prefixParseFn {
	return func() (ast.Expression, *ParsingError) {
		exp := &ast.PrefixExpression{Token: p.curToken}

		p.nextToken()

		v, err := p.parseExpression(precedence)
		if err != nil {
			return nil, err
		}

		exp.Operand = v
		return exp, nil
	}
}

func (p *Parser) parseInfixExpression(left ast.Expression) (ast.Expression, *ParsingError) {
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

func (p *Parser) parseGroupedExpression() (ast.Expression, *ParsingError) {
	p.nextToken()

	// Create something like a local scope inside parentheses and reset precedence
	exp, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, err
	}

	if !p.tryPeek(token.RPAREN) {
		return nil, p.createPeekError("expected to get closing parenthesis")
	}

	return exp, nil
}
