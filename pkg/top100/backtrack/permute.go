package backtrack

// NO.46
func permute(nums []int) [][]int {
	var ans [][]int
	used := make(map[int]bool)
	var backtrack func(path []int)
	backtrack = func(path []int) {
		if len(path) == len(nums) {
			ans = append(ans, append([]int{}, path...))
		}

		for i := 0; i < len(nums); i++ {
			if used[nums[i]] {
				continue
			}
			path = append(path, nums[i])
			used[nums[i]] = true
			backtrack(path)
			path = path[:len(path)-1]
			used[nums[i]] = false
		}
	}

	backtrack([]int{})

	return ans
}
