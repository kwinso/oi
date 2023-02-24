package lexer

import (
	"github.com/stretchr/testify/assert"
	"oilang/internal/token"
	"testing"
)

func TestLexer(t *testing.T) {
	l := New(`let fn true false return if else @fn @fnot
hello hello_123 _name_ a.b
123 123.01 1_000 10_000.12
== != <= >= < >
! and or;
=+-*/**
(){}
->
`)
	tests := []struct {
		Type    token.TokenType
		Literal string
	}{
		{token.LET, "let"},
		{token.FN, "fn"},
		{token.TRUE, "true"},
		{token.FALSE, "false"},
		{token.RETURN, "return"},
		{token.IF, "if"},
		{token.ELSE, "else"},
		{token.PIPE_FN, "@fn"},
		{token.PIPE_CTX, "@"},
		{token.IDENT, "fnot"},
		{token.NEWLINE, "\n"},

		{token.IDENT, "hello"},
		{token.IDENT, "hello_123"},
		{token.IDENT, "_name_"},
		{token.IDENT, "a"},
		{token.DOT, "."},
		{token.IDENT, "b"},
		{token.NEWLINE, "\n"},

		{token.INT, "123"},
		{token.FLOAT, "123.01"},
		{token.INT, "1000"},
		{token.FLOAT, "10000.12"},
		{token.NEWLINE, "\n"},

		// "some string with new\nlines"
		//{token.STRING, "some string with new\n lines"},

		{token.EQ, "=="},
		{token.NEQ, "!="},
		{token.LTE, "<="},
		{token.GTE, ">="},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.NEWLINE, "\n"},

		{token.NOT, "!"},
		{token.AND, "and"},
		{token.OR, "or"},
		{token.SEMICOLON, ";"},
		{token.NEWLINE, "\n"},

		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.MINUS, "-"},
		{token.MULTIPLY, "*"},
		{token.DIVIDE, "/"},
		{token.POWER, "**"},
		{token.NEWLINE, "\n"},

		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.NEWLINE, "\n"},

		{token.PIPE_OP, "->"},
	}

	for _, expected := range tests {
		tok := l.NextToken()

		assert.Equalf(t, expected.Type, tok.Type, "Token types did not match. Expected %s, got %s", expected.Type, tok.Type)
		assert.Equal(t, expected.Literal, tok.Literal, "Token literals did not match")
	}
}

func TestInvalid(t *testing.T) {
	l := New(`10
%|
1__10 10..1`)

	tests := []token.Token{
		{token.INT, "10", 0, 0, ""},
		{token.NEWLINE, "\n", 0, 2, ""},

		{token.ILLEGAL, "%", 1, 0, "unexpected character"},
		{token.ILLEGAL, "|", 1, 1, "unexpected character"},
		{token.NEWLINE, "\n", 1, 2, ""},

		{token.INT, "1", 2, 0, ""},
		{token.IDENT, "__10", 2, 1, ""},
		{token.ILLEGAL, ".", 2, 9, "unexpected fraction delimiter"},
	}

	for _, expected := range tests {
		tok := l.NextToken()

		assert.Equalf(t, expected, tok, "Tokens did not match. Expected %q, got %q", expected, tok)
	}
}
