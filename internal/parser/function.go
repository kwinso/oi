package parser

import (
	"oilang/internal/ast"
	"oilang/internal/token"
)

func (p *Parser) parseFunctionLiteral() (ast.Expression, *ParsingError) {
	f := &ast.FunctionLiteral{}
	f.Token = p.curToken
	f.IsPipelineStage = p.curTokenIs(token.STAGE_FN)

	if p.tryPeek(token.IDENT) {
		f.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	}

	if !p.tryPeek(token.LPAREN) {
		return nil, p.createPeekError("expected function name or (")
	}

	params, err := p.parseFunctionParameters()
	if err != nil {
		return nil, err
	}
	f.Parameters = params

	if !p.tryPeek(token.LBRACE) {
		return nil, p.createPeekError("expected { at the start of function body")
	}

	f.Body, err = p.parseBlockStatement()

	return f, err
}

func (p *Parser) parseFunctionParameters() ([]*ast.Identifier, *ParsingError) {
	var ids []*ast.Identifier

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return ids, nil
	}

	if !p.tryPeek(token.IDENT) {
		return nil, p.createPeekError("expected parameter identifier")
	}

	ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	ids = append(ids, ident)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()

		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		ids = append(ids, ident)
	}

	if !p.tryPeek(token.RPAREN) {
		return nil, p.createPeekError("expected )")
	}

	return ids, nil
}
