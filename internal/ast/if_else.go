package ast

import (
	"fmt"
	"oilang/internal/token"
)

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequnce  *BlockStatement
	Alternative *BlockStatement
}

func (*IfExpression) expressionNode() {}
func (ie *IfExpression) String() string {
	first := fmt.Sprintf("if "+ie.Condition.String()+" { %s }", ie.Consequnce)

	if ie.Alternative != nil {
		first += fmt.Sprintf(" else { %s }", ie.Alternative)
	}

	return first
}
