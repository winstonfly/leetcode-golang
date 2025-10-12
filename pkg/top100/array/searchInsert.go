package main

import "fmt"

//题35
func searchInsert(nums []int, target int) int {

	if len(nums) == 1 {
		if target <= nums[0] {
			return 0
		} else {
			return 1
		}
	}

	low := 0
	high := len(nums) - 1
	middle := low + (high-low)/2 + 1 //防溢出

	if target == nums[middle] {
		return middle
	}

	if target > nums[middle] {
		return middle + searchInsert(nums[middle:], target)
	}

	if target < nums[middle] {
		return searchInsert(nums[:middle], target)
	}

	return middle

}

func searchInsert1(nums []int, target int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}

	return len(nums)
}

func main() {
	fmt.Println(searchInsert1([]int{1, 3, 5, 6}, 7))
}
