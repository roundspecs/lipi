package lexer

import "lipi/token"

type Lexer struct {
	// TODO: Add fields to manage the input string, current position, and any
	// other necessary state.

	// For now, we just define an empty struct.
}

func NewLexer(input string) *Lexer {
	// TODO: Initialize the lexer with the input string and set up any necessary
	// state.

	// For now, we just return a new Lexer instance without any state.
	return &Lexer{}
}

func (l *Lexer) NextToken() token.Token {
	// TODO: Implement the logic to read the next token from the input string.
	// This will involve skipping whitespace, identifying the type of the next token,
	// and returning it.
	
	// For now, it just returns an illegal token.
	return token.Token{Type: token.ILLEGAL, Literal: ""}
}
