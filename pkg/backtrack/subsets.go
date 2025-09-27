package backtrack

// NO.78
func subsets(nums []int) [][]int {
	var ans [][]int
	var backtrack func(start int, path []int)
	backtrack = func(start int, path []int) {
		if len(path) > len(nums) {
			return
		}
		ans = append(ans, append([]int{}, path...))

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i+1, path)
			path = path[:len(path)-1]
		}
	}

	backtrack(0, []int{})
	return ans
}
