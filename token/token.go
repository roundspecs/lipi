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

var keywords = map[string]string{
	"ধর": LET,
	"ফাঙ্কশন": FUNCTION,
}

// LookupIdent checks if the identifier is a keyword and returns its type.
// If not, it returns IDENT.
func LookupIdent(ident string) string {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}