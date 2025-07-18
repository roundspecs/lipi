package repl

import (
	"bufio"
	"fmt"
	"io"
	"os/user"
	"strings"

	"lipi/lexer"
	"lipi/token"
)

const PROMPT = "lipi> "

// REPL represents the Read-Eval-Print Loop.
func Start(r io.Reader, w io.Writer) {
	user, err := user.Current()
	if err != nil {
		fmt.Fprintf(w, "Error getting current user: %v\n", err)
		return
	}

	fmt.Fprintf(w, "Welcome to Lipi REPL, %s!\n", user.Username)
	fmt.Fprintf(w, "Type 'exit' to quit.\n")

	scanner := bufio.NewScanner(r)

	for {
		fmt.Fprint(w, PROMPT)
		if !scanner.Scan() {
			break // Exit on EOF
		}

		line := scanner.Text()
		if strings.TrimSpace(line) == "exit" {
			break // Exit on 'exit' command
		}

		l := lexer.NewLexer(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(w, "%+v\n", tok)
		}
	}

	fmt.Fprintln(w, "Exiting REPL. Goodbye!")
}