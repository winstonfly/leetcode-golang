package backtrack

// NO.22
func generateParenthesis(n int) []string {
	var ans []string
	var backtrack func(left, right int, current string)
	backtrack = func(left, right int, current string) {
		if left == 0 && right == 0 {
			ans = append(ans, current)
			return
		}

		if left > 0 {
			backtrack(left-1, right, current+"(")
		}

		if right > left {
			backtrack(left, right-1, current+")")
		}

	}
	backtrack(n, n, "")
	return ans
}
