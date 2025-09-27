package backtrack

// NO.22
func generateParenthesis(n int) []string {
	var ans []string

	var backtrack func(left, right int, curr string)
	backtrack = func(left, right int, curr string) {

		if left == n && right == n {
			ans = append(ans, curr)
			return
		}

		if left < n {
			backtrack(left+1, right, curr+"(")
		}

		if right < left {
			backtrack(left, right+1, curr+")")
		}
	}

	backtrack(0, 0, "")

	return ans
}
