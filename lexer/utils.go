package lexer

// isDigit checks if a rune is a Bengali digit (০-৯).
func isDigit(ch rune) bool {
	return ch >= '০' && ch <= '৯'
}

// isLetter checks if a rune is a Bengali letter or underscore.
// The Unicode Standard: https://www.unicode.org/charts/PDF/U0980.pdf
func isLetter(ch rune) bool {
	return (ch >= 'ঁ' && ch <= 'ঃ') || (ch >= 'অ' && ch <= 'ঋ') ||
				 (ch >= 'এ' && ch <= 'য়') || ch == '_'
}