package ast

import "oilang/internal/token"

// TODO: Return undefined?

// ReturnStatement Represents return statement in code
//
// Although it seems like "return" should be an expression since it returns a value, but it's not, because it could lead to this confusing syntax:
//
//	let x = return a
type ReturnStatement struct {
	Node
	Token       token.Token
	ReturnValue Expression
}

func (*ReturnStatement) statementNode() {}
func (rs *ReturnStatement) String() string {
	out := "return"

	if rs.ReturnValue != nil {
		out += " " + rs.ReturnValue.String()
	}

	out += ";"

	return out
}
