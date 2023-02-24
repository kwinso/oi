package main

import (
	"github.com/fatih/color"
	"oilang/internal/repl"
	"os"
)

func main() {
	color.Magenta("Welcome to the REPL of oi language. Feel free to play around!")

	repl.Start(os.Stdin, os.Stdout)
}
