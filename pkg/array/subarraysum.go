package main

import "fmt"

func subarraySum(nums []int, k int) int {
	var result [][]int
	for i := 0; i < len(nums); i++ {
		sum := 0
		var index []int
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			index = append(index, nums[j])
			if sum == k {
				result = append(result, index)
			}
		}
	}
	return len(result)

}

func subArrayMax(nums []int) int {
	result := -1 << 31
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > result {
				result = sum
			}
		}
	}

	return result
}

func main() {
	nums := []int{-3, -2, -1}
	result := subArrayMax(nums)
	fmt.Println(result)
}
