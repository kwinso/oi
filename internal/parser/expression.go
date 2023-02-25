package parser

import (
	"fmt"
	"oilang/internal/ast"
)

func (p *Parser) parseExpressionStatement() (*ast.ExpressionStatement, *ParsingError) {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	exp, err := p.parseExpression(LOWEST)
	if err != nil {
		return nil, &ParsingError{Token: p.curToken, Message: err.Error()}
	}
	stmt.Expression = exp

	if p.isEndOfStatementToken(p.peekToken) {
		p.nextToken()
	}

	return stmt, nil
}

func (p *Parser) parseExpression(precedence int) (ast.Expression, error) {
	prefix, ok := p.prefixParsers[p.curToken.Type]
	if !ok {
		return nil, fmt.Errorf("cannot parse %s token", p.curToken.Type)
	}

	leftExp, err := prefix()
	if err != nil {
		return nil, err
	}

	for !p.isEndOfStatementToken(p.peekToken) && precedence < p.peekPrecedence() {
		infix, ok := p.infixParsers[p.peekToken.Type]
		if !ok {
			break
		}

		p.nextToken()

		leftExp, err = infix(leftExp)
		if err != nil {
			return nil, err
		}
	}

	return leftExp, nil
}
