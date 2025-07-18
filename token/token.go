package token

type Token struct {
	Type    string // Type of the token, e.g., IDENT, INT, etc.
	Literal string // The actual string of the token, e.g., "main", "123", etc.
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"

	IDENT     = "IDENT" // main, foo, bar, etc.
	INT       = "INT"   // 1234567890

	ASSIGN    = "ASSIGN"
	PLUS      = "PLUS"
	MINUS     = "MINUS"
	ASTERISK  = "ASTERISK"
	SLASH     = "SLASH"
	BANG     	= "BANG"

	LT        = "LT"
	GT        = "GT"

	COMMA     = "COMMA"
	SEMICOLON = "SEMICOLON"

	L_PAREN   = "L_PAREN"
	R_PAREN   = "R_PAREN"

	L_BRACE   = "L_BRACE"
	R_BRACE   = "R_BRACE"

	EQ				= "EQ"
	NEQ       = "NEQ"

	FUNCTION  = "FUNCTION"
	LET       = "LET"
	IF        = "IF"
	ELSE      = "ELSE"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	RETURN    = "RETURN"
)

var keywords = map[string]string{
	"ধর": LET,
	"ফাঙ্কশন": FUNCTION,
	"যদি": IF,
	"নাহলে": ELSE,
	"সত্য": TRUE,
	"মিথ্যা": FALSE,
	"ফেরাও": RETURN,
}

// LookupIdent checks if the identifier is a keyword and returns its type.
// If not, it returns IDENT.
func LookupIdent(ident string) string {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}