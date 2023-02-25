package ast

import (
	"oilang/internal/token"
)

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (*IntegerLiteral) expressionNode()   {}
func (il *IntegerLiteral) String() string { return il.Token.Literal }

type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (*FloatLiteral) expressionNode()   {}
func (fl *FloatLiteral) String() string { return fl.Token.Literal }
