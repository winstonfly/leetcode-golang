package arrays

func moveZeroes(nums []int) {
	var result []int
	var zeros []int

	for _, v := range nums {
		if v != 0 {
			result = append(result, v)
		} else {
			zeros = append(zeros, v)
		}
	}
	result = append(result, zeros...)
	copy(nums, result)
}
