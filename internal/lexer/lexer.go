package lexer

import (
	"log"
	"oilang/internal/token"
	"strings"
)

// TODO:
// 	- Print lexical errors

// Lexer is responsible for reading input string character by character and converting it to list of tokens
type Lexer struct {
	input          string // input string
	curLine        int
	ch             byte // last read character
	pos            int  // position of current character
	readPos        int  // position of next character to be read by Lexer
	lastNewlinePos int
}

// New Creates new lexer from input string
func New(input string) *Lexer {
	l := &Lexer{input: input}
	// Init lexer with the first character in the input
	l.readNext()

	return l
}

// Tokens that appear as single character
var singleTokens = map[byte]token.TokenType{
	'+': token.PLUS,
	'/': token.DIVIDE,
	',': token.COMMA,
	'{': token.LBRACE,
	'}': token.RBRACE,
	'(': token.LPAREN,
	')': token.RPAREN,
	'.': token.DOT,
	';': token.SEMICOLON,
}

// Tokens that change their type if appeared next to another character
var doubleTokens = map[byte]struct {
	next   byte
	single token.TokenType
	double token.TokenType
}{
	'=': {'=', token.ASSIGN, token.EQ},
	'!': {'=', token.NOT, token.NEQ},
	'>': {'=', token.GT, token.GTE},
	'<': {'=', token.LT, token.LTE},
	'*': {'*', token.MULTIPLY, token.POWER},
	'-': {'>', token.MINUS, token.PIPE_OP},
}

// NextToken Parses next token in the input string
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipToNonWhiteSpace()

	switch l.ch {
	case '\n':
		tok = l.createToken(token.NEWLINE, "\n")
		// Need to add one to properly determine column for tokens
		l.lastNewlinePos = l.pos + 1
		l.curLine += 1
	case '@':
		tok = l.createToken(token.PIPE_CTX, "@")
		lastPos := l.pos
		// Check if next token is a fn keyword
		if l.peekNext() == 'f' {
			l.readNext()
			if l.readIdentifier() == "fn" {
				tok.Type = token.PIPE_FN
				tok.Literal = "@fn"
				break
			}

			l.pos = lastPos
			l.readPos = lastPos + 1
		}
	case 0:
		tok = l.createToken(token.EOF, "")
	default:
		// This cases return because they search until next invalid character. When it encounters, it's under the l.pos
		if isStartingIdentChar(l.ch) {
			tok = l.createToken(token.IDENT, "")
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupTokenType(tok.Literal)

			return tok
		}
		if isDigit(l.ch) {
			return l.readNumber()
		}

		if v, ok := doubleTokens[l.ch]; ok {
			if l.peekNext() == v.next {
				l.skipChars(1)
				tok = l.createToken(v.double, string(l.ch)+string(v.next))
			} else {
				tok = l.createToken(v.single, string(l.ch))
			}
			break
		}
		if v, ok := singleTokens[l.ch]; ok {
			tok = l.createToken(v, string(l.ch))
			break
		}

		tok = l.createToken(token.ILLEGAL, string(l.ch))
		tok.Issue = "unexpected character"
	}

	// Prepare for the next token
	l.readNext()
	return tok
}

// Creates token with position of cursor
func (l *Lexer) createToken(t token.TokenType, lit string) token.Token {
	return token.Token{Type: t, Literal: lit, Line: l.curLine, Col: l.pos - l.lastNewlinePos}
}

// Reads next character into the lexer, setting position to last read position and incrementing reading position
func (l *Lexer) readNext() {
	l.ch = l.peekAt(l.readPos)
	l.pos = l.readPos
	l.readPos += 1
}

func (l *Lexer) skipChars(step int) {
	l.readPos += step
}

// Returns next character that should be read
func (l *Lexer) peekNext() byte {
	return l.peekAt(l.readPos)
}

// Returns character at specified position. Returns ascii "NULL" if reached EOF
func (l *Lexer) peekAt(pos int) byte {
	if pos < 0 {
		log.Fatal("attempted to read character with index less than 0")
	}

	if pos >= len(l.input) {
		return 0
	}

	return l.input[pos]
}

// Reads until finds any non-whitespace character
func (l *Lexer) skipToNonWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' {
		l.readNext()
	}
}

// Reads from current position to next char that cannot be used in identifier / keyword
func (l *Lexer) readIdentifier() string {
	start := l.pos

	for isGeneralIdentChar(l.ch) {
		l.readNext()
	}

	return l.input[start:l.pos]
}

// TODO:
//   - Todo: 0b, 0x and 0o numbers
//
// Reads number from current position, automatically determines type of the number (int or float)
func (l *Lexer) readNumber() token.Token {
	tok := l.createToken(token.INT, "")
	start := l.pos
	hasFrac := false

	for isDigit(l.ch) || l.ch == '.' {
		// Allow underscores for number separation
		if l.peekNext() == '_' {
			if !isDigit(l.peekAt(l.readPos + 1)) {
				l.readNext()
				break
			}
			l.readNext()
		}

		if l.ch == '.' {
			if hasFrac {
				tok = l.createToken(token.ILLEGAL, string(l.ch))
				tok.Issue = "unexpected fraction delimiter"
				return tok
			}

			tok.Type = token.FLOAT
			hasFrac = true
		}
		l.readNext()
	}

	tok.Literal = strings.ReplaceAll(l.input[start:l.pos], "_", "")

	return tok
}
