package arrays

// NO.41  不符合题解， 题目要求用常量级的空间
func firstMissingPositive(nums []int) int {

	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}

	for i := 1; i <= len(nums)+1; i++ {
		if !m[i] {
			return i
		}
	}

	return 0
}
