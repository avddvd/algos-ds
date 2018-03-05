package stack

import (
	"strconv"
	"strings"
)

const (
	closeParens = ")}]"
	openParens  = "({["
	digits      = "0123456789"
	ops         = "+-*/"
)

func matches(open, close string) bool {
	if strings.Index(openParens, open) == strings.Index(closeParens, close) {
		return true
	}
	return false
}

func ParChecker(symbols string) bool {
	balanced := true
	symStack := new(Stack)
	i, n := 0, len(symbols)
	for balanced != false && i < n {
		symbol := string(symbols[i])
		if strings.Contains(openParens, symbol) {
			symStack.Push(symbol)
		} else if strings.Contains(closeParens, symbol) {
			if symStack.IsEmpty() {
				balanced = false
			} else {
				openSym := symStack.Pop().GetData().(string)
				if !matches(openSym, symbol) {
					balanced = false
				}
			}
		}
		i += 1
	}
	if !symStack.IsEmpty() {
		balanced = false
	}
	return balanced
}

func BaseConverter(decimal, base int) string {
	baseValues := "0123456789ABCDEF"
	result := ""
	remStack := new(Stack)
	for decimal > 0 {
		rem := decimal % base
		remStack.Push(string(baseValues[rem]))
		decimal = decimal / base
	}
	for !remStack.IsEmpty() {
		result += remStack.Pop().GetData().(string)
	}
	return result
}

func BaseConverterRecursive(decimal, base int) string {
	charStr := "0123456789ABCDEF"
	if decimal < base {
		return string(charStr[decimal])
	} else {
		digit := decimal % base
		decimal = decimal / base
		return BaseConverterRecursive(decimal, base) + string(charStr[digit])
	}
}

func InfixToPostfix(expr string) string {
	result := ""
	var top string
	opsMap := map[string]int{
		"(": 0,
		")": 0,
		"+": 1,
		"-": 1,
		"*": 2,
		"/": 2,
	}
	opStack := new(Stack)
	for _, sym := range expr {
		symbol := string(sym)
		if strings.Contains(digits, symbol) {
			result += symbol
		} else if symbol == "(" {
			opStack.Push(symbol)
		} else if symbol == ")" {
			top = opStack.Pop().GetData().(string)
			for top != "(" {
				result += top
				top = opStack.Pop().GetData().(string)
			}
		} else if strings.Contains(ops, symbol) {
			for !opStack.IsEmpty() &&
				opsMap[symbol] <= opsMap[opStack.Peek().(string)] {
				result += opStack.Pop().GetData().(string)
			}
			opStack.Push(symbol)
		}
	}
	for !opStack.IsEmpty() {
		result += opStack.Pop().GetData().(string)
	}
	return result
}

func PostfixEval(expr string) string {
	operandStack := new(Stack)
	var operand string
	for _, sym := range expr {
		symbol := string(sym)
		if strings.Contains(digits, symbol) {
			operandStack.Push(symbol)
		} else if strings.Contains(ops, symbol) {
			operand = operandStack.Pop().GetData().(string)
			op2, _ := strconv.Atoi(operand)
			operand = operandStack.Pop().GetData().(string)
			op1, _ := strconv.Atoi(operand)
			res := doMath(op1, op2, symbol)
			operandStack.Push(res)
		}
	}
	return operandStack.Pop().GetData().(string)
}

func doMath(p1, p2 int, symbol string) string {
	sum := 0
	switch symbol {
	case "+":
		sum = p1 + p2
	case "-":
		sum = p1 - p2
	case "*":
		sum = p1 * p2
	case "/":
		sum = p1 / p2
	}
	return strconv.Itoa(sum)
}
