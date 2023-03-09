package lab2

import (
	"fmt"
	"strings"
)

func PrefixToPostfix(input string) (string, error) {
	stack := []string{}
	tokens := strings.Split(input, " ")
	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
		switch {
		case isOperator(token):
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid prefix expression")
			}
			number1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			number2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			postfix := number1 + " " + number2 + " " + token
			stack = append(stack, postfix)
		default:
			stack = append(stack, token)
		}
	}
	if len(stack) != 1 {
		return "", fmt.Errorf("invalid prefix expression")
	}
	return stack[0], nil
}

// isOperator returns true if the given token is an operator.
func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "/" || token == "*" || token == "^"
}
