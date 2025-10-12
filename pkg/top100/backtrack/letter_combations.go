package backtrack

// NO.17
func letterCombinations(digits string) []string {
	var ans []string
	if len(digits) == 0 {
		return ans
	}

	m := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var backtrack func(start int, curr string)
	backtrack = func(start int, curr string) {
		if len(curr) == len(digits) {
			ans = append(ans, curr)
			return
		}

		for i := start; i < len(digits); i++ {
			letters := m[digits[i]]
			for j := 0; j < len(letters); j++ {
				curr += string(letters[j])
				backtrack(i+1, curr)
				curr = curr[:len(curr)-1]
			}
		}
	}

	backtrack(0, "")
	return ans
}
