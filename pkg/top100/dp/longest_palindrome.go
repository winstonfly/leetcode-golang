package dp

// NO.5
func longestPalindrome(s string) string {
	var ans string
	//dp[i][j]表示是否为回文串
	n := len(s)
	dp := make([][]bool, n)
	//长度为1
	for i := range dp {
		dp[i] = make([]bool, n)
		dp[i][i] = true
		ans = s[i : i+1]
	}

	//长度为2
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			ans = s[i : i+2]
		}
	}

	//>=3时
	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
				if length > len(ans) {
					ans = s[i : j+1]
				}
			}
		}
	}
	
	return ans
}

// NO.131
func partition(s string) [][]string {
	var ans [][]string
	n := len(s)
	if n < 2 {
		ans = append(ans, []string{s})
		return ans
	}

	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}

	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
		}
	}

	for length := 3; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] && dp[i+1][j-1] {
				dp[i][j] = true
			}
		}
	}

	splits := []string{}
	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string{}, splits...))
			return
		}

		for j := i; j < n; j++ {
			if dp[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}

	dfs(0)

	return ans
}
