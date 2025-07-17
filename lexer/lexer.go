package lexer

import "lipi/token"

type Lexer struct {
	input []rune    // Input string as a slice of runes
	position int    // Current position in the input string
	ch rune         // Current character being examined
}

func NewLexer(input string) *Lexer {
	runes := []rune(input)
	return &Lexer{
		input:    runes,
		position: 0,
		ch:       runes[0],
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.ch {
	case '=':
		tok = token.Token{Type: token.ASSIGN, Literal: string(l.ch)}
	case '+':
		tok = token.Token{Type: token.PLUS, Literal: string(l.ch)}
	case ';':
		tok = token.Token{Type: token.SEMICOLON, Literal: string(l.ch)}
	case ',':
		tok = token.Token{Type: token.COMMA, Literal: string(l.ch)}
	case '(':
		tok = token.Token{Type: token.L_PAREN, Literal: string(l.ch)}
	case ')':
		tok = token.Token{Type: token.R_PAREN, Literal: string(l.ch)}
	case '{':
		tok = token.Token{Type: token.L_BRACE, Literal: string(l.ch)}
	case '}':
		tok = token.Token{Type: token.R_BRACE, Literal: string(l.ch)}
	default:
		tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
	}
	l.readChar()
	return tok
}

// readChar advances the lexer to the next character in the input.
// It increments the position and updates the current character (ch).
// If the end of input is reached, ch is set to 0.
func (l *Lexer) readChar() {
	l.position++
	if l.position >= len(l.input) {
		l.ch = 0 // End of input
	} else {
		l.ch = l.input[l.position]
	}
}
