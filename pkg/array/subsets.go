package main

import "fmt"

func subsets(nums []int) [][]int {

	var result [][]int

	var backtrack func(start int, index int, current []int)

	backtrack = func(start int, index int, current []int) {

		//if index <= len(nums) {
		//	result = append(result, append([]int{}, current...))
		//	return
		//}
		result = append(result, append([]int{}, current...))

		for i := start; i < len(nums); i++ {
			current = append(current, nums[i])
			backtrack(i+1, index+1, current)
			current = current[:len(current)-1]
		}
	}

	backtrack(0, 0, []int{})
	return result
}

func main() {
	nums := []int{1, 2, 3}
	result := subsets(nums)
	fmt.Println(result)
}
