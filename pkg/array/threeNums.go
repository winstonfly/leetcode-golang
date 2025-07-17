package main

import (
	"fmt"
	"slices"
)

/**
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。





示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
示例 2：

输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。
示例 3：

输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。


提示：

3 <= nums.length <= 3000
-105 <= nums[i] <= 105

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
//使用双指针解决
//func threeSum(nums []int) [][]int {
//
//	var result [][]int
//	nums = sort.Sort(nums)
//
//	for i := 0; i < len(nums); i++ {
//		if nums[i] > 0 {
//			break
//		}
//
//		if i > 0 && (nums[i] == nums[i-1]) {
//			continue
//		}
//
//		L := i + 1
//		R := len(nums) - 1
//
//		for {
//			if L < R {
//				sum := nums[i] + nums[L] + nums[R]
//				if sum == 0 {
//					result = append(result, []int{nums[i], nums[L], nums[R]})
//
//					for {
//						if L < R && nums[L] == nums[L+1] {
//							L++
//							continue
//						}
//						break
//					}
//
//					for {
//						if L < R && nums[R] == nums[R-1] {
//							R--
//							continue
//						}
//						break
//					}
//
//					L++
//					R--
//				}
//
//				if sum < 0 {
//					L++
//				}
//
//				if sum > 0 {
//					R--
//				}
//			} else {
//				break
//			}
//
//		}
//	}
//
//	return result
//}

func threeSum1(nums []int) [][]int {
	n := len(nums)
	var result [][]int
	slices.Sort(nums)
	m := make(map[string]bool)
	for i := 0; i < n; i++ {
		target := -nums[i]
		left, right := i+1, n-1
		for left < right {
			if nums[left]+nums[right] == target {
				s := fmt.Sprintf("%s%s%s", nums[i], nums[left], nums[right])
				if _, ok := m[s]; !ok {
					result = append(result, []int{nums[i], nums[left], nums[right]})
					m[s] = true
				}
				left++
			}
			if nums[left]+nums[right] > target {
				right--
			} else {
				left++
			}
		}
	}

	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum1(nums)
	fmt.Println(result)

}
