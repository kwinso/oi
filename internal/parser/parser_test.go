package parser

import (
	"github.com/stretchr/testify/assert"
	"oilang/internal/lexer"
	"oilang/internal/parser/ast"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5
let y = 10
let z = 10_123.12`

	l := lexer.New(input)
	p, err := New(l).Parse()

	assert.Nil(t, err)
	assert.NotNil(t, p, "Returned program is nil")
	assert.Equal(t, 3, len(p.Statements), "Did not get expected 3 statements")

	tests := []string{"x", "y", "z"}

	for i, expected := range tests {
		stmt, ok := p.Statements[i].(*ast.LetStatement)

		assert.True(t, ok, "cannot convert to LetStatement")
		assert.Equal(t, "let", stmt.TokenLiteral())
		assert.Equal(t, expected, stmt.Name.Value)
		assert.Equal(t, expected, stmt.Name.TokenLiteral())
	}
}

func TestInvalidStatements(t *testing.T) {
	tests := []string{"let x 5", "let = 10;", "let 10_123.12"}

	for _, input := range tests {
		l := lexer.New(input)
		p, err := New(l).Parse()

		assert.Nil(t, p)
		assert.NotNil(t, err, "Should show an error")
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
return 0
return var;
return 12.10`

	l := lexer.New(input)
	p, err := New(l).Parse()

	assert.Nil(t, err)
	assert.NotNil(t, p, "Returned program is nil")
	assert.Equal(t, 3, len(p.Statements), "Did not get expected 3 statements")
}
