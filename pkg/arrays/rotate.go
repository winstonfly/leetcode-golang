package arrays

// NO.189
func rotate(nums []int, k int) {
	n := len(nums)
	if k == 0 || k%n == 0 {
		return
	}

	move := k % n

	var tmp []int
	for i := n - move; i < n; i++ {
		tmp = append(tmp, nums[i])
	}

	for i := n - 1; i >= move; i-- {
		nums[i] = nums[i-move]
	}

	for i := 0; i < move; i++ {
		nums[i] = tmp[i]
	}

}
