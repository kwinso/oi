package parser

import (
	"github.com/stretchr/testify/assert"
	"oilang/internal/ast"
	"oilang/internal/lexer"
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
		assert.Equal(t, "let", stmt.Token.Literal)
		assert.Equal(t, expected, stmt.Name.Value)
		assert.Equal(t, expected, stmt.Name.Token.Literal)
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

func TestIdentifierExpression(t *testing.T) {
	input := "var"
	l := lexer.New(input)
	p, err := New(l).Parse()

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, 1, len(p.Statements))

	stmt, ok := p.Statements[0].(*ast.ExpressionStatement)
	assert.True(t, ok, "Program first statement is not expression statement")

	id, ok := stmt.Expression.(*ast.Identifier)
	assert.True(t, ok, "Statement's expression is not an identifier")
	assert.Equal(t, "var", id.Value)
	assert.Equal(t, "var", id.Token.Literal)
}

func TestIntegerLiterals(t *testing.T) {
	input := `5
50;
123_10`

	l := lexer.New(input)
	p, err := New(l).Parse()
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, 3, len(p.Statements))

	tests := []struct {
		lit string
		val int64
	}{
		{"5", 5},
		{"50", 50},
		{"12310", 12310},
	}

	for i, expected := range tests {
		stmt, ok := p.Statements[i].(*ast.ExpressionStatement)
		assert.True(t, ok, "cannot convert to ExpressionStatement")

		id, ok := stmt.Expression.(*ast.IntegerLiteral)
		assert.True(t, ok, "Statement's expression is not an integer literal")
		assert.Equal(t, expected.lit, id.Token.Literal)
		assert.Equal(t, expected.val, id.Value)
	}
}

func TestFloatLiterals(t *testing.T) {
	input := `5.0
50.123;
123_100.456`

	l := lexer.New(input)
	p, err := New(l).Parse()
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, 3, len(p.Statements))

	tests := []struct {
		lit string
		val float64
	}{
		{"5.0", 5},
		{"50.123", 50.123},
		{"123100.456", 123100.456},
	}

	for i, expected := range tests {
		stmt, ok := p.Statements[i].(*ast.ExpressionStatement)
		assert.True(t, ok, "cannot convert to ExpressionStatement")

		id, ok := stmt.Expression.(*ast.FloatLiteral)
		assert.True(t, ok, "Statement's expression is not a float literal")
		assert.Equal(t, expected.lit, id.Token.Literal)
		assert.Equal(t, expected.val, id.Value)
	}
}
