package ast

import (
	"fmt"
	"oilang/internal/token"
	"strings"
)

type FunctionLiteral struct {
	Token      token.Token
	Name       *Identifier
	Parameters []*Identifier
	Body       *BlockStatement
	// Tells if the function should be run only in pipeline
	IsPipelineStage bool
}

func (*FunctionLiteral) expressionNode() {}
func (fl *FunctionLiteral) String() string {
	var params []string
	for _, p := range fl.Parameters {
		params = append(params, p.String())
	}

	var name string
	if fl.Name != nil {
		name = fl.Name.String()
	}

	return fmt.Sprintf("%v %s(%v) { %s }", fl.Token.Literal, name, strings.Join(params, ", "), fl.Body)
}
