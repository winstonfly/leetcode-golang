package greed

// NO.45
func jump(nums []int) int {
	lastPos := len(nums) - 1
	steps := 0

	for lastPos > 0 {
		for i := 0; i < lastPos; i++ {
			if nums[i]+i >= lastPos {
				lastPos = i
				steps++
				break
			}
		}
	}

	return steps
}
