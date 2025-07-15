package lexer

import (
	"lipi/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `=+,;(){}`

	lexer := NewLexer(input)

	test_tokens := []token.Token{
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.L_PAREN, Literal: "("},
		{Type: token.R_PAREN, Literal: ")"},
		{Type: token.L_BRACE, Literal: "{"},
		{Type: token.R_BRACE, Literal: "}"},
	}

	for i, test_token := range test_tokens {
		next_token := lexer.NextToken()
		if next_token.Type != test_token.Type {
			t.Fatalf(
				"tests[%d] - wrong type. expected=%q, got=%q", i, test_token.Type,
				next_token.Type,
			)
		}
		if next_token.Literal != test_token.Literal {
			t.Fatalf(
				"tests[%d] - wrong literal. expected=%q, got=%q", i, test_token.Literal,
				next_token.Literal,
			)
		}
	}
}