package ast

import (
	"oilang/internal/token"
)

type LetStatement struct {
	Token token.Token // Token that represents let keyword
	Name  *Identifier
	Value Expression
}

func (*LetStatement) statementNode()          {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	var out = ls.TokenLiteral() + " " + ls.Name.String()

	if ls.Value != nil {
		out += " = " + ls.Value.String()
	}

	out += ";"

	return out
}
