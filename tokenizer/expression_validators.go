package tokenizer

import (
	"fmt"
	"gusser/rpn/syntax"
	"strings"
	"unicode"
)

type IllegalExpressionError struct {
	Message    string
	Expression string
	Column     int
}

// Error format IllegalExpressionError with suitable pointer and
// message for given error
func (r IllegalExpressionError) Error() string {
	text := r.Expression
	pointer := strings.Repeat(" ", r.Column) + "^"
	errorText := "error: " + r.Message

	return strings.Join([]string{text, pointer, errorText}, "\n")
}

// ValidateIllegalCharacters check given expression to have illegal
// characters. Illegal characters is any symbol except space, operator
// or symbols from allowed runes map
func ValidateIllegalCharacters(expr string) error {
	for pos, r := range expr {
		if !syntax.IsAllowedSym(r) {
			return IllegalExpressionError{fmt.Sprintf("illegal character '%v'", string(r)), expr, pos}
		}
	}
	return nil
}

// ValidateBracesCount check that given expression has correct number of
// open and closed braces
func ValidateBracesCount(expr string) error {
	closed, open := 0, 0

	for _, r := range expr {
		switch r {
		case '(':
			open++
		case ')':
			closed++
		}
	}

	if closed != open {
		return IllegalExpressionError{
			fmt.Sprintf("incorrect number of open braces and closed braces [%v:%v]", open, closed),
			expr, 0,
		}
	}

	return nil
}

// ValidateExpressionEndsWithOperator this validator checks that expression
// don't end with operator symbol
func ValidateExpressionEndsWithOperator(expr string) error {
	expr = strings.TrimSpace(expr)

	pos := len(expr) - 1

	if syntax.IsOp(string(expr[pos])) {
		return IllegalExpressionError{
			"expression can't end with operator",
			expr, pos,
		}
	}

	return nil
}

// ValidateStartFromAllowedCharacter check that expression starts from
// allowed characters, when start character is illegal this return error
func ValidateStartFromAllowedCharacter(expr string) error {
	expr = strings.TrimSpace(expr)

	first := rune(expr[0])

	if unicode.IsDigit(first) {
		return nil
	} else if op, IsOp := syntax.GetOpt(string(first)); IsOp {
		if !op.Unary {
			return IllegalExpressionError{"expression can starts only from unary syntax + or -", expr, 0}
		}

		if len(expr) < 2 {
			return IllegalExpressionError{"expression can't have only unary operator", expr, 0}
		}

		if next := rune(expr[1]); next != '(' && !unicode.IsDigit(next) {
			return IllegalExpressionError{"after unary operator can be only open brace or number", expr, 1}
		}
	} else if first != '(' {
		return IllegalExpressionError{"expression can start only from digit, ( or unary operator", expr, 0}
	}

	return nil
}

// ValidateOpenBraces this validator checks that before open brace always
// present operator or operator can not be present when brace in start of
// expression
func ValidateOpenBraces(expr string) error {
	expr = strings.TrimSpace(expr)

	prev := rune(0)

	for pos, r := range expr {
		if r == ' ' {
			continue
		}

		if r == '(' {
			if pos > 0 && !syntax.IsOp(string(prev)) {
				return IllegalExpressionError{
					"before open brace always must be operator except start of expression",
					expr, pos,
				}
			}

			if next, hasNext := NextNonSpaceTokenAfter(expr, pos+1); hasNext && next == ')' {
				return IllegalExpressionError{"empty braces not allowed", expr, pos}
			}
		}

		prev = r
	}

	return nil
}

// ValidateOperators check that expression doesn't have unexpected sequences
// of syntax like 36 -+/ 32 in this case allowed only 36 - +32
func ValidateOperators(expr string) error {
	expr = strings.TrimSpace(expr)

	prev := rune(0)
	for p, r := range expr {
		// ignore spaces because it's trash
		if r == ' ' {
			continue
		}

		if op, IsCurrOp := syntax.GetOpt(string(r)); IsCurrOp && syntax.IsOp(string(prev)) {
			if !op.Unary {
				return IllegalExpressionError{
					fmt.Sprintf("after '%v' can follow only '-', '+', '(' or number", string(prev)), expr, p,
				}
			}

			if next, hasNext := NextNonSpaceTokenAfter(expr, p+1); hasNext && syntax.IsOp(string(next)) {
				return IllegalExpressionError{"three or more syntax sequence", expr, p}
			}
		}

		prev = r
	}

	return nil
}

// NextNonSpaceTokenAfter return next rune after given position
// that is not whitespace and return that rune with flag true
// what means that token present otherwise this return 0 in rune
// and false in hasNext status, that means that any symbol expect
// whitespace not present in given string
func NextNonSpaceTokenAfter(str string, pos int) (rune, bool) {
	for i := pos; i < len(str); i++ {
		if tok := rune(str[i]); tok != ' ' {
			return tok, true
		}
	}
	return 0, false
}
