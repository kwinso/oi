package ast

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	// Dummy method to identify Statement node
	statementNode()
}
type Expression interface {
	Node
	// Dummy method to identify Expression node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

//func (p *Program) TokenLiteral() {
//	if len(p.Statements) > 0 {
//
//	}
//}
