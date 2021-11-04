package tokenizer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorTextFormat(t *testing.T) {
	err := IllegalExpressionError{"illegal character", "256 - a", 6}
	mess := err.Error()

	expected := "256 - a\n      ^\nerror: illegal character"

	assert.Equal(t, expected, mess)
}

func TestValidateBracesCount(t *testing.T) {
	assert.Error(t, ValidateBracesCount("7 * (9 * 2))"))
}

func TestValidateBracesCountWhenBracesValid(t *testing.T) {
	assert.NoError(t, ValidateBracesCount("7 * (9 * (3 - 2))"))
}

func TestValidateExpressionEndsWithOperator(t *testing.T) {
	assert.Error(t, ValidateExpressionEndsWithOperator("7 * 2 + +"))
	assert.Error(t, ValidateExpressionEndsWithOperator("7 * 2 + "))
}

func TestValidateIllegalCharacters(t *testing.T) {
	assert.Error(t, ValidateIllegalCharacters("2 * 3 + a"))
	assert.Error(t, ValidateIllegalCharacters("2 * 3 + ,"))
	assert.Error(t, ValidateIllegalCharacters("2 * 3 + !"))
	assert.Error(t, ValidateIllegalCharacters("2 * 3 + ;"))
}

func TestValidateIllegalCharactersValidExpr(t *testing.T) {
	assert.NoError(t, ValidateIllegalCharacters("2 * 2 / 3"))
	assert.NoError(t, ValidateIllegalCharacters("1 * (3 / (9 + 17))"))
	assert.NoError(t, ValidateIllegalCharacters("8.5 / 2.5"))
}

func TestValidateStartFromAllowedCharacterError(t *testing.T) {
	assert.Error(t, ValidateStartFromAllowedCharacter("-+29"))
	assert.Error(t, ValidateStartFromAllowedCharacter(")32*28"))
	assert.Error(t, ValidateStartFromAllowedCharacter(",32*28"))
	assert.Error(t, ValidateStartFromAllowedCharacter("!32*28"))
	assert.Error(t, ValidateStartFromAllowedCharacter("*32*28"))
	assert.Error(t, ValidateStartFromAllowedCharacter("/32*28"))
	assert.Error(t, ValidateStartFromAllowedCharacter("--"))
	assert.Error(t, ValidateStartFromAllowedCharacter("-"))
	assert.Error(t, ValidateStartFromAllowedCharacter("/"))
}

func TestValidateStartFromAllowedCharacterLegal(t *testing.T) {
	assert.NoError(t, ValidateStartFromAllowedCharacter("-(92+36)*12"))
	assert.NoError(t, ValidateStartFromAllowedCharacter("-(92+36)*12"))
	assert.NoError(t, ValidateStartFromAllowedCharacter("(62*93)+87"))
	assert.NoError(t, ValidateStartFromAllowedCharacter("6.5+32"))
	assert.NoError(t, ValidateStartFromAllowedCharacter("+5+5"))
}

func TestValidateOpenBraces(t *testing.T) {
	assert.NoError(t, ValidateOpenBraces("(92 + 45) * 3"))
	assert.NoError(t, ValidateOpenBraces("(92 + 45)"))
	assert.NoError(t, ValidateOpenBraces("-(92 + 45)"))
	assert.NoError(t, ValidateOpenBraces("65 * (99 - 32)"))
	assert.NoError(t, ValidateOpenBraces("65 *(99 - 32)"))
	assert.NoError(t, ValidateOpenBraces("  (  67* 32) / 33"))
}

func TestValidateOpenBracesErrors(t *testing.T) {
	assert.Error(t, ValidateOpenBraces("98 (32+6)"))
}

func TestValidateOperatorsErrors(t *testing.T) {
	assert.Error(t, ValidateOperators("34 +/17"))
	assert.Error(t, ValidateOperators("34 +/+ 17"))
	assert.Error(t, ValidateOperators("34 +/+* 17"))
}

func TestValidateOperatorsValid(t *testing.T) {
	assert.NoError(t, ValidateOperators("32 + 17"))
	assert.NoError(t, ValidateOperators("32 + -17"))
	assert.NoError(t, ValidateOperators("+32 / -17"))
	assert.NoError(t, ValidateOperators("+32 * 17"))
	assert.NoError(t, ValidateOperators("+32 ++ 17"))
}
