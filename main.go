package main

import (
	"lipi/repl"
	"os"
)

func main() {
	// Start the REPL (Read-Eval-Print Loop)
	repl.Start(os.Stdin, os.Stdout)
}