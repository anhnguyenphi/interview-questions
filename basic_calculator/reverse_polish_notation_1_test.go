package basic_calculator

import (
	"strconv"
	"testing"

	"github.com/fannpa/interviewquestions/util"

	"github.com/magiconair/properties/assert"
)

// https://cp-algorithms.com/string/expression_parsing.html

type Op struct {
	Val     string
	IsUnary bool
}

func isDelim(s string) bool {
	return s == " "
}

func isUnaryOp(s string) bool {
	return s == "+" || s == "-"
}

func processOp(st *util.IntStack, op Op) {
	if op.IsUnary {
		r, _ := st.Pop()
		switch op.Val {
		case "+":
			st.Push(r)
		case "-":
			st.Push(-r)
		}
		return
	}

	right, _ := st.Pop()
	left, _ := st.Pop()
	switch op.Val {
	case "+":
		st.Push(right + left)
	case "-":
		st.Push(left - right)
	case "*":
		st.Push(right * left)
	case "/":
		st.Push(left / right)
	}
}

func opPriority(op Op) int {
	if op.IsUnary {
		return 3
	}
	if op.Val == "+" || op.Val == "-" {
		return 1
	}
	if op.Val == "*" || op.Val == "/" {
		return 2
	}
	return 0
}

func convertToNumber(s string) (int, error) {
	return strconv.Atoi(s)
}

func isOp(s string) bool {
	return s == "+" || s == "-" || s == "*" || s == "/"
}

func reversePolishEval(expr string) int {
	st := &util.IntStack{}
	op := &util.AnyStack{}
	mayBeUnary := true

	for i := 0; i < len(expr); i++ {
		c := string(expr[i])
		if isDelim(c) {
			continue
		}
		if c == "(" {
			op.Push(Op{Val: "(", IsUnary: false})
			continue
		}

		if c == ")" {
			for top, ok := op.Top(); ok && top.(Op).Val != "("; top, ok = op.Top() {
				processOp(st, top.(Op))
				op.Pop()
			}
			op.Pop()
			mayBeUnary = false
			continue
		}

		if isOp(c) {
			isUn := isUnaryOp(c) && mayBeUnary
			curOp := Op{Val: c, IsUnary: isUn}
			for o, _ := op.Top(); !op.IsEmpty() && ((isUn && opPriority(curOp) < opPriority(o.(Op))) || (!isUn && opPriority(curOp) <= opPriority(o.(Op)))); o, _ = op.Top() {
				processOp(st, o.(Op))
				op.Pop()
			}
			op.Push(Op{Val: c, IsUnary: isUn})
			mayBeUnary = true
			continue
		}

		num := 0
		for digit, err := convertToNumber(c); err == nil; digit, err = convertToNumber(c) {
			num = num*10 + digit
			i++
			if i == len(expr) {
				break
			}
			c = string(expr[i])
		}
		i--

		st.Push(num)
		mayBeUnary = false
	}

	for !op.IsEmpty() {
		o, _ := op.Pop()
		processOp(st, o.(Op))
	}

	result, _ := st.Top()
	return result
}

func TestCase1(t *testing.T) {
	assert.Equal(t, reversePolishEval("2 * 3 + 2 * (1 + 3)"), 14)
}

func TestCase2(t *testing.T) {
	assert.Equal(t, reversePolishEval("101 + 2 * 30 + (3 - 1) * (1 + 3)"), 169)
}

func TestCase3(t *testing.T) {
	assert.Equal(t, reversePolishEval("(+-1)*2 + 3"), 1)
}
