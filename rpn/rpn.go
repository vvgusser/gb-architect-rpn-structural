package rpn

import (
	"errors"
	"fmt"
	"gusser/rpn/collection"
	"gusser/rpn/syntax"
	"gusser/rpn/tokenizer"
	"strconv"
	"strings"
)

const MaxExprLength = 256

type TooLongExpressionError struct {
	MaxLen, Len int
}

func (e TooLongExpressionError) Error() string {
	return fmt.Sprintf("expression too long! allowed max length = %v, given = %v", e.MaxLen, e.Len)
}

// Rpn evaluate given infix expression and also convert it
// to RPN form
func Rpn(expr string) (string, float64, error) {
	expr = strings.TrimSpace(expr)

	if len(expr) >= MaxExprLength {
		return "", 0, TooLongExpressionError{MaxExprLength, len(expr)}
	}

	tokens, err := tokenizer.Tokenize(expr)

	if err != nil {
		return "", 0, err
	}

	rpn := shuntingYard(tokens)
	sol, err := solveRpnFormula(rpn)

	if err != nil {
		return "", 0, err
	}

	return strings.Join(rpn, " "), sol, nil
}

// shuntingYard use source infix token for rearrange it
// in RPN manner.
func shuntingYard(tokens []string) []string {
	var result []string
	stack := collection.InitStack()

	for _, s := range tokens {
		switch s {
		case "(":
			stack.Push(s)
		case ")":
			for stack.IsNonEmpty() {
				pop, _ := stack.Pop()
				if pop == "(" {
					break
				}
				result = append(result, pop.(string))
			}
		default:
			if op, isOp := syntax.GetOpt(s); isOp {
				for stack.IsNonEmpty() {
					top, _ := stack.Top()
					op2, isOp := syntax.GetOpt(top.(string))
					if !isOp || op.Precedence > op2.Precedence {
						break
					}
					_, _ = stack.Pop()
					result = append(result, top.(string))
				}
				stack.Push(s)
			} else {
				result = append(result, s)
			}
		}
	}

	for stack.IsNonEmpty() {
		pop, _ := stack.Pop()
		result = append(result, pop.(string))
	}

	return result
}

// solveRpnFormula evaluate slice of rpn tokens and return
// result or error
func solveRpnFormula(tokens []string) (float64, error) {
	stack := collection.InitStack()

	for _, tok := range tokens {
		if op, isOp := syntax.GetOpt(tok); isOp {
			op1, op2 := popOperands(stack)

			if !op2.IsSet && !op.Unary {
				return 0, errors.New("last operator is not unary, expression error")
			}

			ans, err := syntax.GetHandler(tok)(op2, op1)

			if err != nil {
				return 0, err
			}

			stack.Push(ans)
		} else {
			f, err := strconv.ParseFloat(tok, 64)
			if err != nil {
				return 0, err
			}
			stack.Push(f)
		}
	}

	el, _ := stack.Top()
	return el.(float64), nil
}

// popOperands utility function that extract two operands from stack
// when on stack remains only one operand this return second unset
// operand
func popOperands(s *collection.Stack) (syntax.Operand, syntax.Operand) {
	o1, _ := s.Pop()
	o2, err := s.Pop()

	oper1 := syntax.Operand{IsSet: true, Value: o1.(float64)}

	if err != nil {
		return oper1, syntax.Operand{}
	}

	return oper1, syntax.Operand{IsSet: true, Value: o2.(float64)}
}
