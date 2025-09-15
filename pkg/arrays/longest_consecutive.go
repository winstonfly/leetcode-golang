package arrays

// NO.128
func longestConsecutive(nums []int) int {
	m := make(map[int]bool, 1)
	for _, v := range nums {
		m[v] = true
	}

	ans := 0
	n := len(nums)
	for i, _ := range m {
		if !m[i-1] {
			tmp := 0
			for j := 0; j < n; j++ {
				if m[i+j] {
					tmp++
				} else {
					i = i + j - 1
					break
				}
			}
			ans = max(ans, tmp)
		}
	}

	return ans
}
