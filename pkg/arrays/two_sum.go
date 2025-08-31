package arrays

func twoSum(nums []int, target int) []int {
	var ans []int
	m := make(map[int]int, 1)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok {
			ans = append(ans, v)
			ans = append(ans, i)
		} else {
			m[nums[i]] = i
		}
	}

	return ans
}
