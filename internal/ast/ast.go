package ast

import (
	"oilang/internal/token"
)

type Node interface {
	String() string
}

// Statement is a type of node that does not return value, but just declares something (like let statement)
type Statement interface {
	Node
	// Any struct implementing Statement interface should declare this function
	statementNode()
}

// Expression is a type of node that returns value
type Expression interface {
	Node
	// Any struct implementing Expression interface should declare this function
	expressionNode()
}

// ExpressionStatement is a type of expression that could appear at top level.
//
// For example, this code should be valid with ExpressionStatement:
//
//	let x = 1
//	x + 1 // this line
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (*ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// Program is represents sequence of parsed statements that create a program
type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var out = ""

	for _, s := range p.Statements {
		out += s.String()
	}

	return out
}
