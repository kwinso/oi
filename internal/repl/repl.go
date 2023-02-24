package repl

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"io"
	"oilang/internal/lexer"
	"oilang/internal/token"
)

const PROMPT = "oi: "

// TODO:
// - Allow to use newlines
// - Add commands: .help, .exit, .export
// - Exit with Ctrl + D
// - Stop execution with Ctrl + C

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		_, _ = color.New(color.FgMagenta).Print(PROMPT)

		input := scanner.Scan()
		if !input {
			return
		}

		l := lexer.New(scanner.Text())

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			_, _ = fmt.Fprintf(out, "%v: %v\n", tok.Type, tok.Literal)
		}
	}
}
