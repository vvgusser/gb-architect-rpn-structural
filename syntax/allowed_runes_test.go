package syntax

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAllowedSymIllegals(t *testing.T) {
	illegal := []rune{
		'_', '=',
		'a', 'b',
		'c', 'z',
		';', '{',
		'}', ',',
	}

	for _, tok := range illegal {
		assert.False(t, IsAllowedSym(tok))
	}
}

func TestIsAllowedSymLegals(t *testing.T) {
	legals := []rune{
		'1', '2',
		'3', '4',
		'5', '6',
		'7', '8',
		'9', '0',
		'(', ')',
		'-', '+',
		'/', '*',
		' ',
	}

	for _, legal := range legals {
		assert.True(t, IsAllowedSym(legal))
	}
}
