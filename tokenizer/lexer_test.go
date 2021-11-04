package tokenizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsingEmptyExpr(t *testing.T) {
	empty := []string{"", " ", "  "}

	for _, exp := range empty {
		_, err := Tokenize(exp)

		assert.Error(t, err)
	}
}

func TestParseStrangeOperators(t *testing.T) {
	expected := []string{
		"2", "+", "3", "-", "-5",
	}

	tokens, err := Tokenize("2+3--5")

	assert.NoError(t, err)
	assert.Equal(t, expected, tokens)
}

func TestTokenizationWithBraces(t *testing.T) {
	expected := []string{
		"2", "+", "(", "9", "-", "5", ")", "*", "1",
	}

	tokens, err := Tokenize("2 + (9 - 5) * 1")

	assert.NoError(t, err)
	assert.Equal(t, expected, tokens)
}
