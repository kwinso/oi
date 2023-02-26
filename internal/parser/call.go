package parser

import (
	"oilang/internal/ast"
	"oilang/internal/token"
)

func (p *Parser) parseCallExpression(called ast.Expression) (ast.Expression, *ParsingError) {
	call := &ast.CallExpression{Token: p.curToken, CalledExpression: called}
	args, err := p.parseCallArguments()
	if err != nil {
		return nil, err
	}
	call.Arguments = args

	return call, nil
}

// FIXME: Looks awful
func (p *Parser) parseCallArguments() ([]ast.Expression, *ParsingError) {
	var args []ast.Expression
	var err *ParsingError

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return args, nil
	}

	p.nextToken()
	exp, err := p.parseExpression(LOWEST)
	if err == nil {
		args = append(args, exp)

		for p.peekTokenIs(token.COMMA) {
			p.nextToken()
			p.nextToken()

			exp, err = p.parseExpression(LOWEST)
			if err != nil {
				break
			}

			args = append(args, exp)
		}
	}

	if err == nil && !p.tryPeek(token.RPAREN) {
		return nil, p.createPeekError("expected )")
	}

	return args, err
}
