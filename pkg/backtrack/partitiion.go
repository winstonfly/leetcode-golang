package backtrack

// NO.131
func partition(s string) [][]string {
	var ans [][]string

	//单个字符都是回文串

	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}

	//1. 如果s[i]==s[j]且s[i+1:j-1]是回文串，那么s[i:j]也是回文串
	//2. 如果s[i]==s[j]且j-i<2，那么s[i:j]也是回文串
	for length := 2; length <= len(s); length++ {
		for i := 0; i <= len(s)-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && (j-i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
			}
		}
	}

	var dfs func(i int, curr []string)
	dfs = func(i int, curr []string) {
		if i == len(s) {
			ans = append(ans, append([]string{}, curr...))
			return
		}

		for j := i; j < len(s); j++ {
			if dp[i][j] {
				curr = append(curr, s[i:j+1])
				dfs(j+1, curr)
				curr = curr[:len(curr)-1]
			}
		}
	}

	dfs(0, []string{})
	return ans
}
