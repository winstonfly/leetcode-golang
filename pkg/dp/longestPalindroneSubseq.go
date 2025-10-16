package dp

// NO.516
func longestPalindromeSubseq(s string) int {
	// dp[i][j] 表示s[i:j]的最长回文子序列长度，则dp[i][j]  = dp[i+1][j-1] + 2, 这个是指在s[i] == s[j]的情况下，如果不相等dp[i][j] = max(dp[i+1][j], dp[i][j-1])
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for i := n - 1; i >= 0; i-- {
		j := i + 1
		for ; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	return dp[0][n-1]
}
