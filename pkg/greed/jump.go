package greed

// NO.45
func jump(nums []int) int {
	lastPos := len(nums) - 1
	start := lastPos
	times := 0
	for lastPos > 0 && start >= 0 {
		if start+nums[start] >= lastPos {
			start--
			continue
		}
		//最左侧
		lastPos = start
		times++
	}

	return times
}
