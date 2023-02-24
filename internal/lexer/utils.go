package lexer

// Checks if byte can be used for starting identifier / keyword
func isStartingIdentChar(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Identifiers could contain digits, but not at the start
func isGeneralIdentChar(ch byte) bool {
	return isStartingIdentChar(ch) || isDigit(ch)
}

// Checks if byte is an integer char (0 - 9)
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
