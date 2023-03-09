package lab2

import (
	"fmt"
	"strings"
)

func PrefixToPostfix(input string) (string, error) {
	stack := []string{}
	tokens := strings.Fields(input)
	tokenWasNumber := false
	for i := len(tokens) - 1; i >= 0; i-- {
		token := tokens[i]
		switch {
		case isOperator(token):
			if len(stack) < 2 {
				return "", fmt.Errorf("invalid expression: %s", input)
			}
			postfix := ""
			number1 := stack[len(stack)-1]
			number2 := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			if tokenWasNumber && len(stack) == 0 {
				postfix = number2 + " " + number1 + " " + token
			} else {
				postfix = number1 + " " + number2 + " " + token
			}
			stack = append(stack, postfix)
			tokenWasNumber = false
		case isNumber(token):
			stack = append(stack, token)
			tokenWasNumber = true
		default:
			return "", fmt.Errorf("invalid token: %s", token)
		}
	}
	if len(stack) != 1 {
		return "", fmt.Errorf("invalid expression: %s", input)
	}
	return stack[0], nil
}

// isOperator returns true if the given token is an operator.
func isOperator(token string) bool {
	return token == "+" || token == "-" || token == "/" || token == "*" || token == "^"
}

// isNumber returns true if the token is a number.
func isNumber(token string) bool {
	return !isOperator(token)
}
