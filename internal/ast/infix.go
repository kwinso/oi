package ast

import (
	"fmt"
	"oilang/internal/token"
)

// InfixExpression is an expression that holds to expression on each side of the token (addition, product, power and so on)
type InfixExpression struct {
	Token token.Token // Represents infix token, e.g. "+" in "5 + 5" or "and" in "true and false"
	Left  Expression
	Right Expression
}

func (*InfixExpression) expressionNode() {}
func (pe *InfixExpression) Operator() string {
	return pe.Token.Literal
}
func (pe *InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", pe.Left, pe.Operator(), pe.Right)
}
