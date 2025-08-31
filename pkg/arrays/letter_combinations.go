package arrays

var m = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {

	var ans []string

	var backtrack func(index int, path []byte)

	backtrack = func(index int, path []byte) {
		if digits == "" {
			return
		}
		if len(path) == len(digits) {
			ans = append(ans, string(path))
			return
		}
		letters := m[digits[index]]
		if letters == "" {
			return
		}
		for j := 0; j < len(letters); j++ {
			path = append(path, letters[j])
			backtrack(index+1, path)
			path = path[:len(path)-1]
		}
	}

	backtrack(0, []byte{})

	return ans
}
