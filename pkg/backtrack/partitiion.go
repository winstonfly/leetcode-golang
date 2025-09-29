package backtrack

// NO.131
func partition(s string) [][]string {
	var ans [][]string

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	for length := 2; length <= len(s); length++ {
		for i := 0; i <= len(s)-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && (length == 2 || dp[i+1][j-1]) { // length == 2 等同于j < i + 2
				dp[i][j] = true
			}
		}
	}

	var backtrack func(i int, curr []string)
	backtrack = func(i int, curr []string) {
		if i == len(s) {
			ans = append(ans, append([]string{}, curr...))
			return
		}

		for j := i; j < len(s); j++ {
			if dp[i][j] {
				curr = append(curr, s[i:j+1])
				backtrack(j+1, curr)
				curr = curr[:len(curr)-1]
			}
		}
	}

	backtrack(0, []string{})

	return ans
}
