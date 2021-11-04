package syntax

import (
	"errors"
)

// Operator structure which represent mathematical
// operator which this program can evaluate
//
// When operator is unary, then evaluator can pass
// as second argument Operand which has isSet flag
// to false, it's mean that evaluator has one operand
// and one unary operator
type Operator struct {
	Unary      bool
	Precedence int
}

type Operand struct {
	IsSet bool
	Value float64
}

var operators = map[string]Operator{
	"+": {true, 2},
	"-": {true, 2},
	"*": {false, 3},
	"/": {false, 3},
}

// plus operator sum two operands, this is also
// unary operator and can handle only one operand
// when second has IsSet = false
func plus(a, b Operand) (float64, error) {
	if !a.IsSet {
		return b.Value, nil
	}
	return a.Value + b.Value, nil
}

// minus operator subtract two operands, this is
// also unary operator and can handle only one
// operand when second has IsSet = false in this
// case we negate one accepted operand
func minus(a, b Operand) (float64, error) {
	if !a.IsSet {
		return -b.Value, nil
	}
	return a.Value - b.Value, nil
}

// divide operator operate over both operands, this
// may return error when evaluated divider is zero
func divide(a, b Operand) (float64, error) {
	if b.Value == 0 {
		return 0, errors.New("division by zero")
	}
	return a.Value / b.Value, nil
}

// multiply simple binary operator which multiply
// first operand to second
func multiply(a, b Operand) (float64, error) {
	return a.Value * b.Value, nil
}

// handlers map of handling functions for all operators
var handlers = map[string]func(a, b Operand) (float64, error){
	"+": plus,
	"-": minus,
	"/": divide,
	"*": multiply,
}

// GetOpt return operator meta information and boolean
// flag which means that operator present or absent.
func GetOpt(s string) (op Operator, present bool) {
	op, present = operators[s]
	return
}

// GetHandler find handler for given operator and return
// than handler and present flag, if handler exist this
// function return that handler and present flag is true
// otherwise handler will be nil and flag is false
func GetHandler(s string) (handler func(a, b Operand) (float64, error)) {
	handler, present := handlers[s]

	if !present {
		panic("critical error. no handler for operator '" + s + "'")
	}

	return
}

// IsOp return true when given string has reference to
// operator in operators map
func IsOp(s string) (present bool) {
	_, present = GetOpt(s)
	return
}
