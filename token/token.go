package token

type Token struct {
	Type    string // Type of the token, e.g., IDENT, INT, etc.
	Literal string // The actual string of the token, e.g., "main", "123", etc.
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT   = "IDENT" // main, foo, bar, etc.
	INT     = "INT"   // 1234567890

	ASSIGN  = "="
	PLUS    = "+"

	COMMA   = ","
	SEMICOLON = ";"

	L_PAREN = "("
	R_PAREN = ")"

	L_BRACE = "{"
	R_BRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)