package sort

func sortColors(nums []int) {
	//使用insertSort
	for i := 1; i < len(nums); i++ {
		j := i
		value := nums[j]
		for ; j > 0 && value < nums[j-1]; j-- {
			// j - 1后移一位
			nums[j] = nums[j-1]
		}

		//insert
		nums[j] = value
	}
}
