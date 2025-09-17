package str

// NO.3
func lengthOfLongestSubstring(s string) int {
	var ans int

	m := make(map[byte]int, 1)
	left, right := 0, 0
	for right < len(s) {
		if v, ok := m[s[right]]; ok {
			ans = max(ans, right-left)
			left = max(left, v+1)
		}
		m[s[right]] = right
		right++
	}

	ans = max(ans, right-left)
	return ans
}
