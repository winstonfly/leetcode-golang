package dp

// NO.139
func wordBreak(s string, wordDict []string) bool {

	// f[i] 代表能否被前i个字符组成
	f := make([]bool, len(s)+1)
	f[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			word := s[j:i]
			f[i] = f[j] && contains(wordDict, word)
			if f[i] {
				break
			}
		}
	}

	return f[len(s)]
}

func contains(wordDict []string, word string) bool {
	for _, v := range wordDict {
		if v == word {
			return true
		}
	}

	return false
}
