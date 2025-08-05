package main

func main() {
	s := "cbababdc"
	longestPalindrome(s)
}

func longestPalindrome(s string) string {
	// dp[i][j]代表其中一个子串， 为true时代表回文串， false时否

	n := len(s)
	dp := make([][]bool, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
	}

	//所有长度为1的都是回文串
	maxLen := 1
	start := 0
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	if len(s) < 2 {
		return s
	}

	//从长度为2开始
	for L := 2; L <= n; L++ {
		//左边界
		for i := 0; i < n; i++ {
			//由长度和左边界，确定右边界
			j := L + i - 1

			if j >= n {
				break
			}

			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if j-i < 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}

			if dp[i][j] && j-i+1 > maxLen {
				maxLen = j - i + 1
				start = i
			}

		}
	}

	return s[start : maxLen+start]

}
