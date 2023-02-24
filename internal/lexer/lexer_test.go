package lexer

import (
	"oilang/internal/token"
	"reflect"
	"testing"
)

func TestLexer(t *testing.T) {
	l := New(`let fn true false return if else @fn @fnot
hello hello_123 _name_ a.b
123 123.01 1_000 10_000.12
== != <= >= < >
! and or 
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

	for i, expected := range tests {
		tok := l.NextToken()

		if tok.Type != expected.Type {
			t.Fatalf("Test [%d] didn't match: Wrong token type. Expected=%q, but got=%q",
				i, expected.Type.String(), tok.Type.String())
		}
		if tok.Literal != expected.Literal {
			t.Fatalf("Test [%d] didn't match: Wrong wrong literal value. Expected=%q, but got=%q",
				i, expected.Type.String(), tok.Type.String())
		}
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

	for i, expected := range tests {
		tok := l.NextToken()

		if !reflect.DeepEqual(tok, expected) {
			t.Fatalf("Test [%d] didn't match. Got %v instead of %v", i, tok, expected)
		}
	}
}
