package ast

import "oilang/internal/token"

type Identifier struct {
	Token token.Token // Identifier token
	Value string      // Name of the identifier
}

func (*Identifier) expressionNode()        {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
func (i *Identifier) String() string       { return i.Value }
