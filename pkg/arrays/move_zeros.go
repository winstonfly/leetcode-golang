package arrays

// NO.283
// 常规则两层循环解法
func moveZeroes1(nums []int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] != 0 {
					nums[i], nums[j] = nums[j], nums[i]
					break
				}
			}
		}
	}
}

func moveZeroes(nums []int) {
	left, right := 0, 1

	for right < len(nums) {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}

		right++
	}
}
