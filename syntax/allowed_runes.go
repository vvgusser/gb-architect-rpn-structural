package syntax

import "unicode"

// allowedRunes is map of allowed non operator symbols
// expect numbers
var allowedRunes = map[rune]bool{
	'(': true,
	')': true,
	// for allow real numbers
	'.': true,
}

// IsAllowedSym return true when given rune is allowed
// for expression otherwise false
func IsAllowedSym(r rune) bool {
	return unicode.IsSpace(r) ||
		IsOp(string(r)) ||
		unicode.IsDigit(r) ||
		isAllowedRune(r)
}

// IsAllowedRune return true when given rune exist in allowedRunes
// otherwise this return false
func isAllowedRune(r rune) (isAllowed bool) {
	_, isAllowed = allowedRunes[r]
	return
}
