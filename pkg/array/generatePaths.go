package main

import (
	"fmt"
)

var (
	leftK  = "("
	rightK = ")"
)

func generateParenthesis(n int) []string {
	var result []string
	var backtrack func(left, right int, current string)
	backtrack = func(left, right int, current string) {
		if left == 0 && right == 0 {
			//都用完，找到有效的括号组合
			result = append(result, current)
			return
		}

		if left > 0 {
			//如果左括号还有剩余，可以添加左括号
			backtrack(left-1, right, current+"(")
		}

		if right > left {
			backtrack(left, right-1, current+")")
		}
	}

	backtrack(n, n, "")
	return result
}

func main() {
	n := 5
	result := generateParenthesis(n)
	fmt.Println(result)
}
