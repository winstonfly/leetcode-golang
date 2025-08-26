package dp

// NO.139
func wordBreak(s string, wordDict []string) bool {

	wordDictSet := make(map[string]bool)
	for _, v := range wordDict {
		wordDictSet[v] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	// dp[i] 表示字符串在第i个之前的是否能使用wordDict中单词拼出， 则递推公司为则dp[i-1]就是指前i个拼成与否与字典中是否包含第i-1到到i个字符的值
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
