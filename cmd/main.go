package main

import (
	"fmt"
	"github.com/fatih/color"
	"oilang/internal/repl"
	"os"
)

const BANNER = `░░░░░░░█████╗░██╗░░░░░░░█████╗░██╗░░░░░░
░░░░░░██╔══██╗██║░░░░░░██╔══██╗██║░░░░░░
░░░░░░██║░░██║██║█████╗██║░░██║██║░░░░░░
░░░░░░██║░░██║██║╚════╝██║░░██║██║░░░░░░
██╗██╗╚█████╔╝██║░░░░░░╚█████╔╝██║██╗██╗
╚═╝╚═╝░╚════╝░╚═╝░░░░░░░╚════╝░╚═╝╚═╝╚═╝
███████╗██████╗░░██╗░░░░░░░██╗██╗███╗░░██╗
██╔════╝██╔══██╗░██║░░██╗░░██║██║████╗░██║
█████╗░░██████╔╝░╚██╗████╗██╔╝██║██╔██╗██║
██╔══╝░░██╔══██╗░░████╔═████║░██║██║╚████║
███████╗██║░░██║░░╚██╔╝░╚██╔╝░██║██║░╚███║
╚══════╝╚═╝░░╚═╝░░░╚═╝░░░╚═╝░░╚═╝╚═╝░░╚══╝`

func main() {
	color.Magenta(BANNER)
	color.White("Welcome to the REPL of oi language.\nFeel free to play around!")
	fmt.Println()

	repl.Start(os.Stdin, os.Stdout)
}
