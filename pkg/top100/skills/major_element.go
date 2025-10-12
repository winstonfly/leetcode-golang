package skills

// NO.169
func majorityElement(nums []int) int {
	counts := make(map[int]int)
	major := nums[0]
	for _, num := range nums {
		counts[num]++
		if counts[num] > counts[major] {
			major = num
		}
	}
	return major
}
