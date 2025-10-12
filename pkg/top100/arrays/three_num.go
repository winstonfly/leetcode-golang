package arrays

import (
	"fmt"
	"sort"
)

// NO.15
// 和为0的不同的三元组, 直接三层循环，超时
func threeSum1(nums []int) [][]int {
	var ans [][]int
	m := make(map[string]bool, 1)
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				var tmp []int
				tmp = append(tmp, num, nums[j], nums[k])
				sort.Ints(tmp)
				if _, ok := m[fmt.Sprint(tmp)]; ok {
					continue
				}
				if nums[j]+nums[k] == -num {
					m[fmt.Sprint(tmp)] = true
					ans = append(ans, tmp)
				}
			}
		}
	}

	return ans
}

// 还是会超时
func threeSum2(nums []int) [][]int {
	var ans [][]int
	m := make(map[string]bool, 1)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		num := nums[i]
		//转为两数之和
		ms := make(map[int]bool, 1)
		for j := i + 1; j < len(nums); j++ {

			var tmp []int
			if _, ok := ms[-num-nums[j]]; ok {
				tmp = append(tmp, num, nums[j], -num-nums[j])
				sort.Ints(tmp)
				if _, ok := m[fmt.Sprint(tmp)]; ok {
					continue
				} else {
					m[fmt.Sprint(tmp)] = true
					ans = append(ans, tmp)
				}
			} else {
				ms[nums[j]] = true
			}
		}

	}

	return ans
}

// 双指针解法
func threeSum(nums []int) [][]int {
	var ans [][]int
	n := len(nums)
	sort.Ints(nums)

	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -nums[i]
		//从左循环
		right := n - 1
		for left := i + 1; left < n; left++ {
			if left > i+1 && nums[left] == nums[left-1] {
				continue
			}

			for left < right && nums[left]+nums[right] > target {
				right--
			}

			if left == right {
				break
			}

			if nums[left]+nums[right] == target {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
			}
		}
	}

	return ans
}
