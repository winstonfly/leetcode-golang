package skills

func findDuplicate(nums []int) int {
	// 287. 寻找重复数
	// 1. 哈希表
	// 2. 二分查找
	// 3. 快慢指针
	slow, fast := nums[0], nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}
