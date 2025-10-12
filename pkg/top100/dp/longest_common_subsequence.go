package dp

// NO.1143
func longestCommonSubsequence(text1 string, text2 string) int {
	//dp[i][j]表示text1[0:i]和text2[0:j]的最长公共子序列长度,分两种情况,第一种是text1[i]==text2[j],那么dp[i][j]=dp[i-1][j-1]+1

	m, n := len(text1), len(text2)

	//将dp数组初始化为m+1行n+1列,多出的一行一列用来表示空字符串, 这样可以避免边界条件的判断
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[m][n]
}
