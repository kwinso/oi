package parser

import (
	"github.com/stretchr/testify/assert"
	"oilang/internal/ast"
	"oilang/internal/lexer"
	"reflect"
	"testing"
)

func testValidProgram(t *testing.T, program *ast.Program, err *ParsingError, amount int) {
	assert.Nil(t, err)
	assert.NotNil(t, program)
	assert.Len(t, program.Statements, amount)
}

func getAsInstanceOf[T any](t *testing.T, val any) *T {
	stmt, ok := val.(*T)
	assert.Truef(t, ok, "cannot be converted %v", reflect.TypeOf(*new(T)))

	return stmt
}

func TestLetStatements(t *testing.T) {
	input := `
let x = 5
let y = 10
let z = 10_123.12`

	l := lexer.New(input)
	p, err := New(l).Parse()

	testValidProgram(t, p, err, 3)

	tests := []string{"x", "y", "z"}

	for i, expected := range tests {
		let := getAsInstanceOf[ast.LetStatement](t, p.Statements[i])
		assert.Equal(t, "let", let.Token.Literal)
		assert.Equal(t, expected, let.Name.Value)
		assert.Equal(t, expected, let.Name.Token.Literal)
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

	testValidProgram(t, p, err, 3)
}

func TestIdentifierExpression(t *testing.T) {
	input := "var"
	l := lexer.New(input)
	p, err := New(l).Parse()

	testValidProgram(t, p, err, 1)

	stmt := getAsInstanceOf[ast.ExpressionStatement](t, p.Statements[0])
	id := getAsInstanceOf[ast.Identifier](t, stmt.Expression)

	assert.Equal(t, "var", id.Value)
	assert.Equal(t, "var", id.Token.Literal)
}

func TestIntegerLiterals(t *testing.T) {
	input := `5
50;
123_10`

	l := lexer.New(input)
	p, err := New(l).Parse()
	testValidProgram(t, p, err, 3)

	tests := []struct {
		lit string
		val int64
	}{
		{"5", 5},
		{"50", 50},
		{"12310", 12310},
	}

	for i, expected := range tests {
		stmt := getAsInstanceOf[ast.ExpressionStatement](t, p.Statements[i])
		id := getAsInstanceOf[ast.IntegerLiteral](t, stmt.Expression)

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
	testValidProgram(t, p, err, 3)

	tests := []struct {
		lit string
		val float64
	}{
		{"5.0", 5},
		{"50.123", 50.123},
		{"123100.456", 123100.456},
	}

	for i, expected := range tests {
		stmt := getAsInstanceOf[ast.ExpressionStatement](t, p.Statements[i])
		float := getAsInstanceOf[ast.FloatLiteral](t, stmt.Expression)

		assert.Equal(t, expected.lit, float.Token.Literal)
		assert.Equal(t, expected.val, float.Value)
	}
}

func TestPrefixOperators(t *testing.T) {
	tests := []struct {
		input    string
		operator string
		operand  string
	}{
		{"not 123", "not", "123"},
		{"-var", "-", "var"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p, err := New(l).Parse()

		testValidProgram(t, p, err, 1)

		stmt := getAsInstanceOf[ast.ExpressionStatement](t, p.Statements[0])
		exp := getAsInstanceOf[ast.PrefixExpression](t, stmt.Expression)

		assert.Equal(t, test.operator, exp.Operator())
		assert.Equal(t, test.operand, exp.Operand.String())
	}
}

func TestInfixOperators(t *testing.T) {
	tests := []struct {
		input    string
		left     string
		operator string
		right    string
	}{
		{"1 + 1", "1", "+", "1"},
		{"1 - 1", "1", "-", "1"},
		{"1 * 1", "1", "*", "1"},
		{"1/1", "1", "/", "1"},
		{"2**2", "2", "**", "2"},

		{"1>1", "1", ">", "1"},
		{"1 < 1", "1", "<", "1"},
		{"1 >= 1", "1", ">=", "1"},
		{"1 <= 1", "1", "<=", "1"},
		{"1 != 3", "1", "!=", "3"},
		{"3 == 3", "3", "==", "3"},

		{"1 or 1", "1", "or", "1"},
		{"1 and 1", "1", "and", "1"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p, err := New(l).Parse()

		testValidProgram(t, p, err, 1)

		stmt := getAsInstanceOf[ast.ExpressionStatement](t, p.Statements[0])
		exp := getAsInstanceOf[ast.InfixExpression](t, stmt.Expression)

		assert.Equal(t, test.operator, exp.Operator())
		assert.Equal(t, test.left, exp.Left.String())
		assert.Equal(t, test.right, exp.Right.String())
	}
}

func TestPrecedence(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"-a * b", "((- a) * b)"},
		{"a + -b", "(a + (- b))"},
		{"not -a", "(not (- a))"},
		{"5 + 2 * 10", "(5 + (2 * 10))"},
		{"123 * 3 * 2 ** 3", "((123 * 3) * (2 ** 3))"},
		{"4 <= 5 != 5 >= 4", "((4 <= 5) != (5 >= 4))"},
		{"a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)"},
		{"3 + 4 * 5 == 3 * 1 + 4 * 5 or a == b", "(((3 + (4 * 5)) == ((3 * 1) + (4 * 5))) or (a == b))"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p, err := New(l).Parse()

		testValidProgram(t, p, err, 1)

		assert.Equal(t, test.expected, p.String())
	}
}

func TestBools(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"false", "false"},
		{"true", "true"},
		{"foobar != true", "(foobar != true)"},
		{"let a = false", "let a = false;"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p, err := New(l).Parse()

		testValidProgram(t, p, err, 1)

		assert.Equal(t, test.expected, p.String())
	}
}
