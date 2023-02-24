package token

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
	//"@":   PIPE_CTX,
}

// TODO: Add track of lines and pos
type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Col     int
	Issue   string
}

// LookupTokenType Matches supplied string with the possible keywords and returns TokenType if matched, and IDENT if not
func LookupTokenType(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
