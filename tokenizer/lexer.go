package tokenizer

import (
	"errors"
	"gusser/rpn/syntax"
	"strings"
	"unicode"
)

var validators = []func(expr string) error{
	ValidateIllegalCharacters,
	ValidateBracesCount,
	ValidateStartFromAllowedCharacter,
	ValidateExpressionEndsWithOperator,
	ValidateOpenBraces,
	ValidateOperators,
}

// Tokenize accept expression string and return tokens from
// that string, this also validate illegal characters and
// return error if it exists.
func Tokenize(expr string) ([]string, error) {
	expr = strings.TrimSpace(expr)

	if len(expr) == 0 {
		return nil, errors.New("empty expressions not acceptable")
	}

	if err := validateExpression(expr); err != nil {
		return nil, err
	}

	return tokenizeValidExpression(expr), nil
}

// validateExpression check source expression string
// for illegal characters. allowed characters is
//
// - digits
// - operators
// - allowedRunes
//
// Any other symbol is illegal
func validateExpression(expr string) error {
	for _, validate := range validators {
		if err := validate(expr); err != nil {
			return err
		}
	}
	return nil
}

// tokenizeValidExpression this is internal function that tokenize
// clear expression to tokens, this correctly handle braces, numbers
// real numbers and operators
func tokenizeValidExpression(expr string) []string {
	// delete all whitespaces because it's trash
	expr = strings.ReplaceAll(expr, " ", "")

	var tokens []string

	i := 0
	N := len(expr)

	for i < N {
		s := rune(expr[i])

		if s == '(' || s == ')' {
			tokens = append(tokens, string(s))
			i++
		} else {
			if _, isOp := syntax.GetOpt(string(s)); isOp {
				if len(tokens) > 0 && !isLastTokenOperator(tokens) {
					tokens = append(tokens, string(s))
					i++
					continue
				}
			}

			j := i
			for j < N {
				l := rune(expr[j])
				if j > i && !unicode.IsDigit(l) && l != '.' {
					break
				}
				j++
			}
			tokens = append(tokens, expr[i:j])
			i = j
		}
	}

	return tokens
}

// isLastTokenOperator checks that in tokens slice last token
// exists in operators map
func isLastTokenOperator(tokens []string) bool {
	if len(tokens) == 0 {
		return false
	}
	return syntax.IsOp(tokens[len(tokens)-1])
}
