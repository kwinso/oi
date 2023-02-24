package ast

import "oilang/internal/token"

// TODO: Return undefined?

type ReturnStatement struct {
	Node
	Token token.Token
	Value Expression
}

func (*ReturnStatement) statementNode()          {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }
