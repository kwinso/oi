package ast

import (
	"github.com/stretchr/testify/assert"
	"oilang/internal/token"
	"testing"
)

func TestProgram_String(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "var"},
					Value: "var",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "another"},
					Value: "another",
				},
			},
			&ReturnStatement{
				Token: token.Token{Type: token.RETURN, Literal: "return"},
				ReturnValue: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "var"},
					Value: "var",
				},
			},
		},
	}

	assert.Equal(t, "let var = another;return var;", program.String())

}
