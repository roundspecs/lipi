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
	l.eatSpaces() // Skip whitespace characters
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
	case 0:
		tok = token.Token{Type: token.EOF, Literal: ""}
	default:
		if isDigit(l.ch) {
			tok = token.Token{Type: token.INT, Literal: l.readNumber()}
			return tok
		} else if isLetter(l.ch) {
			// Identifiers in Lipi can include Bengali letters, digits, and underscores
			// But, it cannot start with a digit
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else {
			tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch)}
		}
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


// readNumber reads a sequence of digits and returns it as a string.
func (l *Lexer) readNumber() string {
	start := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return string(l.input[start:l.position])
}

// readIdentifier reads a sequence of letters and digits (including underscores)
// and returns it as a string.
func (l *Lexer) readIdentifier() string {
	start := l.position
	for isLetter(l.ch) || isDigit(l.ch) {
		l.readChar()
	}
	return string(l.input[start:l.position])
}

// eatSpaces skips over any whitespace characters (spaces, newlines, tabs)
func (l *Lexer) eatSpaces() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' {
		l.readChar()
	}
}