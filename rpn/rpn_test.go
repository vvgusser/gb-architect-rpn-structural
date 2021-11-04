package rpn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRpnValidExpressions(t *testing.T) {
	type Test struct {
		expr string
		rpn  string
		ans  float64
	}

	table := []Test{
		{"(1 + 2) * 4 + 3", "1 2 + 4 * 3 +", 15},
		{"9 + 8", "9 8 +", 17},
		{"-(8 / 2) * 4", "8 2 / 4 * -", -16},
		{"98 / 32 + 16 * (78 - 78)", "98 32 / 16 78 78 - * +", 3.0625},
		{"0 * (0)", "0 0 *", 0},
		{"32 / 89 * 8", "32 89 / 8 *", 2.8764044943820224},
		{"0.5 + 1.5", "0.5 1.5 +", 2},
		{"9 * (8 / (4 / 2))", "9 8 4 2 / / *", 36},
	}

	for _, d := range table {
		rpn, ans, err := Rpn(d.expr)

		assert.NoError(t, err)
		assert.Equal(t, d.rpn, rpn)
		assert.Equal(t, d.ans, ans)
	}
}

func TestRpnInvalidExpressions(t *testing.T) {
	expr := []string{
		"",
		"--9",
		"98+-78+",
		"/(8709)+8",
		"32(78/2)+8",
		"3+++9",
		"-+9",
		"-",
		"+",
		"/",
		"*",
		"a-15+c",
		"  +17**9",
		"-17+9-  ",
		"328 / 9 **        88",
		" ",
		"   ",
	}

	for _, ex := range expr {
		_, _, err := Rpn(ex)

		assert.Error(t, err)
	}
}

func TestRpnDivisionByZero(t *testing.T) {
	_, _, err := Rpn("78 / (10 - 9 - 1)")

	assert.Error(t, err)
}

func TestRpnVeryLongExpr(t *testing.T) {
	expr := ""

	for i := 0; i <= MaxExprLength; i++ {
		expr += "98 + 98 "
	}

	_, _, err := Rpn(expr)

	assert.Error(t, err)
}

func TestOctalNumberExpressions(t *testing.T) {
	type Case struct {
		expr string
		ans  float64
	}

	table := []Case{
		{"06 + 06", 12},
		{"20 + 09", 29},
		{"00030 + 030 + 30", 90},
	}

	for _, row := range table {
		_, ans, err := Rpn(row.expr)

		assert.NoError(t, err)
		assert.Equal(t, row.ans, ans)
	}
}
