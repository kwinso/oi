package ast

import (
	"fmt"
	"oilang/internal/token"
	"strings"
)

type CallExpression struct {
	Token            token.Token
	CalledExpression Expression
	Arguments        []Expression
}

func (*CallExpression) expressionNode() {}
func (ce *CallExpression) String() string {
	var params []string
	for _, p := range ce.Arguments {
		params = append(params, p.String())
	}

	return fmt.Sprintf("%s(%v)", ce.CalledExpression, strings.Join(params, ", "))
}
