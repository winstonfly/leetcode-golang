package skills

// NO.136
func singleNumber(nums []int) int {

	var value int
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = m[nums[i]] + 1
	}
	for k, v := range m {
		if v == 1 {
			value = k
			break
		}
	}

	return value
}
