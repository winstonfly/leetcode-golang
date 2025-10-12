package main

import (
	"fmt"
	"leetcode-golang/pkg/sort"
)

func fourSum(nums []int, target int) [][]int {

	nums = sort.Sort(nums)

	var result [][]int
	m := make(map[string]bool)
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		newTarget := target - a
		res := threeSums(nums[i+1:], newTarget)
		if len(res) > 0 {
			for _, v := range res {
				key := fmt.Sprintf("%d%d%d%d", nums[i], v[0], v[1], v[2])
				if !m[key] {
					result = append(result, []int{nums[i], v[0], v[1], v[2]})
					m[key] = true
				}

			}
		}
	}

	return result
}

func threeSums(nums []int, target int) [][]int {

	var result [][]int
	for i := 0; i < len(nums); i++ {
		if i > 0 && (nums[i] == nums[i-1]) {
			continue
		}

		L := i + 1
		R := len(nums) - 1
		for {
			if L < R {
				sum := nums[i] + nums[L] + nums[R]
				if sum == target {
					result = append(result, []int{nums[i], nums[L], nums[R]})
					if nums[i] == nums[L] && nums[L] == nums[R] {
						break
					}

					L++
					R--
					continue
				}
				if sum < target {
					L++
				} else if sum > target {
					R--
				}
			} else {
				break
			}
		}
	}

	return result
}

func main() {
	nums := []int{2, 2, 2, 2, 2}
	result := fourSum(nums, 8)
	fmt.Println(result)
}
