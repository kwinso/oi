package ast

import "oilang/internal/token"

type BoolExpression struct {
	Token token.Token
	Value bool
}

func (*BoolExpression) expressionNode()   {}
func (be *BoolExpression) String() string { return be.Token.Literal }
