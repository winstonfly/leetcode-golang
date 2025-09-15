package dp

// NO.1143
func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)

	// dp[i][j]代表text[0..i]与text[0-j]的最长子串
	dp := make([][]int, m+1)
	for i, _ := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		c1 := text1[i]
		for j := 0; j < n; j++ {
			c2 := text2[j]
			if c1 == c2 {
				//如果两个字符相等，则为dp[i-1][j-1] + 1
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				//如果不相等，则应该考虑text1(0..n-1), text2(0-j)与text1
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[m][n]
}
