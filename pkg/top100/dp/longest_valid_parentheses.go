package dp

// NO.32
func longestValidParentheses(s string) int {
	var ans int
	// dp[i]表示以i结尾的最长有效括号的长度
	dp := make([]int, len(s))

	//分两种情况，第一种是()结尾的，即如果i == ")" && s[i-1] == "("，则dp[i] = dp[i-2] + 2
	//第二种是))结尾的，如果i==")" && s[i-1] == ")"，则需要看s[i-dp[i-1]-1]是否为"("，如果是，则dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
	for i := 0; i < len(s); i++ {
		if s[i] == ')' {
			if i > 0 && s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			}

			if i > 0 && s[i-1] == ')' && i-dp[i-1]-1 >= 0 && s[i-dp[i-1]-1] == '(' {
				dp[i] = dp[i-1] + 2
				if i-dp[i-1]-2 >= 0 {
					dp[i] += dp[i-dp[i-1]-2]
				}
			}
		}
	}

	for _, v := range dp {
		ans = max(ans, v)
	}

	return ans
}
