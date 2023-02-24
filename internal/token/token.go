package token

import "fmt"

//go:generate stringer -type=TokenType
type TokenType int

const (
	ILLEGAL TokenType = iota
	EOF
	NEWLINE

	IDENT
	INT
	FLOAT
	TRUE
	FALSE
	STRING

	COMMA
	DOT
	SEMICOLON
	LPAREN
	RPAREN
	LBRACE
	RBRACE

	ASSIGN
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	POWER

	// Logical operators
	AND // &&
	OR  // ||
	NOT // !

	// Comparison operators
	EQ  // ==
	NEQ // !+
	LT  // <
	GT  // >
	LTE // <=
	GTE // >=

	LET
	FN
	RETURN
	IF
	ELSE

	// Piping
	PIPE_CTX // @
	PIPE_FN  // @fn
	PIPE_OP  // ->
)

var keywords = map[string]TokenType{
	"let":    LET,
	"fn":     FN,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"and":    AND,
	"or":     OR,
}

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Col     int
	Issue   string
}

func (t Token) String() string {
	return fmt.Sprintf("[%v:%v] %s %v", t.Line, t.Col, t.Type, t.Literal)
}

// LookupTokenType Matches supplied string with the possible keywords and returns TokenType if matched, and IDENT if not
func LookupTokenType(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
