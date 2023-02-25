package ast

import (
	"fmt"
	"oilang/internal/token"
)

type PrefixExpression struct {
	Token   token.Token // Represents prefix token, e.g. "not" or "-"
	Operand Expression
}

func (*PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) Operator() string {
	return pe.Token.Literal
}
func (pe *PrefixExpression) String() string {
	return fmt.Sprintf("(%s %s)", pe.Operator(), pe.Operand)
}
