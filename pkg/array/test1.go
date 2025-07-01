package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 0 {
		os.Exit(-1)
	}

	input := os.Args[1]
	s := input[1 : len(input)-1] // Remove leading and trailing brackets
	tokens := strings.Split(s, ",")
	fmt.Println(evalRPN(tokens))
}

func evalRPN(tokens []string) int {
	if len(tokens) < 1 {
		return -1
	}

	var stack []int
	for i := 0; i < len(tokens); i++ {
		v := tokens[i]
		value, err := strconv.Atoi(v)
		if err != nil {
			if len(stack) < 2 {
				return -1
			}

			length := len(stack)

			a := stack[length-1]
			b := stack[length-2]
			switch v {
			case "+":
				stack[length-2] = a + b
			case "-":
				stack[length-2] = a + b
			case "*":
				stack[length-2] = a * b
			case "/":
				if a == 0 {
					return -1
				} else {
					stack[length-2] = b / a
				}
			}

			stack = stack[:length-1]
		} else {
			stack = append(stack, value)
		}

	}

	return stack[0]
}
