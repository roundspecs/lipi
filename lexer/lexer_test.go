package lexer

import (
	"lipi/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `ধর সংখ্যা১ = ৫;
ধর সংখ্যা২ = ১০;

ধর যোগ_করো = ফাঙ্কশন(ক, খ) {
    ক + খ;
};

ধর ফলাফল = যোগ_করো(সংখ্যা১, সংখ্যা২);`

	lexer := NewLexer(input)

	test_tokens := []token.Token{
		{Type: token.LET, Literal: "ধর"},
		{Type: token.IDENT, Literal: "সংখ্যা১"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "৫"},
		{Type: token.SEMICOLON, Literal: ";"},

		{Type: token.LET, Literal: "ধর"},
		{Type: token.IDENT, Literal: "সংখ্যা২"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.INT, Literal: "১০"},
		{Type: token.SEMICOLON, Literal: ";"},

		{Type: token.LET, Literal: "ধর"},
		{Type: token.IDENT, Literal: "যোগ_করো"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.FUNCTION, Literal: "ফাঙ্কশন"},
		{Type: token.L_PAREN, Literal: "("},
		{Type: token.IDENT, Literal: "ক"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENT, Literal: "খ"},
		{Type: token.R_PAREN, Literal: ")"},
		{Type: token.L_BRACE, Literal: "{"},
		{Type: token.IDENT, Literal: "ক"},
		{Type: token.PLUS, Literal: "+"},
		{Type: token.IDENT, Literal: "খ"},
		{Type: token.SEMICOLON, Literal: ";"},
		{Type: token.R_BRACE, Literal: "}"},
		{Type: token.SEMICOLON, Literal: ";"},

		{Type: token.LET, Literal: "ধর"},
		{Type: token.IDENT, Literal: "ফলাফল"},
		{Type: token.ASSIGN, Literal: "="},
		{Type: token.IDENT, Literal: "যোগ_করো"},
		{Type: token.L_PAREN, Literal: "("},
		{Type: token.IDENT, Literal: "সংখ্যা১"},
		{Type: token.COMMA, Literal: ","},
		{Type: token.IDENT, Literal: "সংখ্যা২"},
		{Type: token.R_PAREN, Literal: ")"},
		{Type: token.SEMICOLON, Literal: ";"},
		
		{Type: token.EOF, Literal: ""},
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