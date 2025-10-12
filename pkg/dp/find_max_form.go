package dp

// NO.474
func findMaxForm(strs []string, m int, n int) int {
	// dp[i][j] 表示能组成有i个0和j个1的情况下，子集的数量, dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones] + 1)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		zeros, ones := 0, 0
		for _, c := range str {
			if c == '0' {
				zeros++
			} else {
				ones++
			}
		}

		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones]+1)
			}
		}

	}

	return dp[m][n]
}
