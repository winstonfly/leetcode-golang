package skills

// NO.75
func sortColor(nums []int) {
	// 0, 1, 2
	// 三路快排
	l, r := 0, len(nums)-1
	for i := 0; i <= r; {
		if nums[i] == 0 {
			nums[i], nums[l] = nums[l], nums[i]
			l++
			i++
		} else if nums[i] == 2 {
			nums[i], nums[r] = nums[r], nums[i]
			r--
		} else {
			i++
		}
	}
}
