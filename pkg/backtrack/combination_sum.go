package backtrack

// NO.39
func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int

	var backtrack func(start int, path []int, sum int)

	backtrack = func(start int, path []int, sum int) {
		if sum > target {
			return
		}
		if sum == target {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			sum += candidates[i]
			backtrack(i, path, sum)
			path = path[:len(path)-1]
			sum -= candidates[i]
		}
	}

	backtrack(0, []int{}, 0)

	return ans
}
