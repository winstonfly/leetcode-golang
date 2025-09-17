package arrays

// NO.560
func subarraySum(nums []int, k int) int {
	var count int
	pre := make(map[int]int, 1)
	sum := 0
	pre[0] = 1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if _, ok := pre[sum-k]; ok {
			count += pre[sum-k]
		}

		pre[sum] += 1
	}

	return count
}
