package parser

import (
	"github.com/stretchr/testify/assert"
	"oilang/internal/ast"
	"oilang/internal/lexer"
	"reflect"
	"testing"
)

func testValidProgram(t *testing.T, program *ast.Program, err *ParsingError, statementsLen int) {
	assert.Nil(t, err)
	assert.NotNil(t, program)
	assert.Len(t, program.Statements, statementsLen)
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

	for i, test := range tests {
		let := getAsInstanceOf[ast.LetStatement](t, p.Statements[i])
		assert.Equal(t, "let", let.Token.Literal)
		assert.Equal(t, test, let.Name.Value)
		assert.Equal(t, test, let.Name.Token.Literal)
	}
}

func TestInvalidStatements(t *testing.T) {
	tests := []string{"let x 5", "let = 10;", "let 10_123.12"}

	for _, test := range tests {
		l := lexer.New(test)
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

	for i, test := range tests {
		stmt := getAsInstanceOf[ast.ExpressionStatement](t, p.Statements[i])
		id := getAsInstanceOf[ast.IntegerLiteral](t, stmt.Expression)

		assert.Equal(t, test.lit, id.Token.Literal)
		assert.Equal(t, test.val, id.Value)
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

	for i, test := range tests {
		stmt := getAsInstanceOf[ast.ExpressionStatement](t, p.Statements[i])
		float := getAsInstanceOf[ast.FloatLiteral](t, stmt.Expression)

		assert.Equal(t, test.lit, float.Token.Literal)
		assert.Equal(t, test.val, float.Value)
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
		// Automatic precedence
		{"-a * b", "((- a) * b)"},
		{"a + -b", "(a + (- b))"},
		{"not -a", "(not (- a))"},
		{"5 + 2 * 10", "(5 + (2 * 10))"},
		{"123 * 3 * 2 ** 3", "((123 * 3) * (2 ** 3))"},
		{"4 <= 5 != 5 >= 4", "((4 <= 5) != (5 >= 4))"},
		{"a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)"},
		{"3 + 4 * 5 == 3 * 1 + 4 * 5 or a == b", "(((3 + (4 * 5)) == ((3 * 1) + (4 * 5))) or (a == b))"},

		// Forced precedence
		{"1 + (2 + 3) + 4", "((1 + (2 + 3)) + 4)"},
		{"(5 + 5) * 2", "((5 + 5) * 2)"},
		{"2 / (5 + 5)", "(2 / (5 + 5))"},
		{"-(5 + 5)", "(- (5 + 5))"},
		{"!(true == true)", "(! (true == true))"},
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

func TestIfExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"if x > y { true } else { false}", "if (x > y) { true } else { false }"},
		{"if x ** 2 > 128 { 128 }", "if ((x ** 2) > 128) { 128 }"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p, err := New(l).Parse()

		testValidProgram(t, p, err, 1)

		assert.Equal(t, test.expected, p.String())
	}
}

func TestFunctions(t *testing.T) {
	tests := []struct {
		input        string
		name         string
		paramsAmount int
		isStageFn    bool
	}{
		{`fn (x, y) { return x + y }`, "", 2, false},
		{`fn () { return; }`, "", 0, false},
		{`fn (x) { return 123; }`, "", 1, false},
		{`@fn (x, y) { return x + y }`, "", 2, true},
		{`fn func(x, a, b, y) { true }`, "func", 4, false},
		{`@fn stage(x, a, z) { true }`, "stage", 3, true},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p, err := New(l).Parse()

		testValidProgram(t, p, err, 1)
		stmt := getAsInstanceOf[ast.ExpressionStatement](t, p.Statements[0])
		fn := getAsInstanceOf[ast.FunctionLiteral](t, stmt.Expression)

		assert.Lenf(t, fn.Parameters, test.paramsAmount, "parameters amount do not match")
		if test.name != "" {
			assert.NotNil(t, fn.Name)
			assert.Equal(t, test.name, fn.Name.Value)
		} else {
			assert.Nil(t, fn.Name)
		}

		assert.Equal(t, test.isStageFn, fn.IsPipelineStage)
	}

}

func TestBadIfSyntax(t *testing.T) {
	tests := []struct {
		input string
		error string
	}{
		{"if { true } else { false}", "unexpected token"},
		{"if true  128 }", "expected { for main if branch"},
		{"if true  { 128 ", "expected } at the end of block"},
		{"if true  { 128 } else ", "expected { for else branch"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p, err := New(l).Parse()

		assert.Nil(t, p)
		assert.NotNil(t, err)
		assert.Equal(t, test.error, err.Message)
	}
}

func TestExpressionCall(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"1(a, b)", "1(a, b)"},
		{"a()", "a()"},
		{"a + b(12)", "(a + b(12))"},
		{"(c + d)(a, b, 12)", "(c + d)(a, b, 12)"},
		{"@fn (x, y) { return x > y }(1, 2)", "@fn (x, y) { return (x > y); }(1, 2)"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p, err := New(l).Parse()
		testValidProgram(t, p, err, 1)

		assert.Equal(t, test.expected, p.String())
	}
}

func TestBadExpressionCallSyntax(t *testing.T) {
	tests := []struct {
		input string
		error string
	}{
		{"1 a, b)", "unexpected token"},
		{"a(q", "expected )"},
		{"a + b(,)", "unexpected token"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p, err := New(l).Parse()

		assert.Nil(t, p)
		assert.NotNil(t, err)
		assert.Equal(t, test.error, err.Message)
	}
}

func TestBadFNSyntax(t *testing.T) {
	tests := []struct {
		input string
		error string
	}{
		{"fn 1()", "expected function name or ("},
		{"@fn name()  128 }", "expected { at the start of function body"},
		{"fn ()  { 128 ", "expected } at the end of block"},
		{"fn (213)  { 128 ", "expected parameter identifier"},
	}

	for _, test := range tests {
		l := lexer.New(test.input)
		p, err := New(l).Parse()

		assert.Nil(t, p)
		assert.NotNil(t, err)
		assert.Equal(t, test.error, err.Message)
	}
}
