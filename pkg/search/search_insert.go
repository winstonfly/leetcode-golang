package search

// 1,3, 5, 6插入7
// 1,3
func searchInsert(tokens []int, target int) int {

	var dfs func(nums []int, target int, start int)
	var pos int
	dfs = func(nums []int, target int, start int) {
		if len(nums) == 0 {
			return
		}

		if len(nums) == 1 {
			if nums[0] >= target {
				if start > 0 {
					pos = start - 1
				} else {
					pos = start
				}
				return
			}
			pos = start + 1
			return
		}

		n := len(nums)
		middle := n / 2

		if nums[middle] == target {
			pos = start + middle
			return
		}

		if nums[middle] > target {
			dfs(nums[:middle], target, start)
		} else {
			dfs(nums[middle:], target, middle+start)
		}
	}

	dfs(tokens, target, 0)
	return pos
}
