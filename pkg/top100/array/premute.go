package main

import "fmt"

func permute(nums []int) [][]int {
	var result [][]int

	m := make(map[int]bool, 1)
	for _, v := range nums {
		m[v] = true
	}
	var backtrack func(index int, current []int)
	backtrack = func(index int, current []int) {

		if index == len(nums) {
			result = append(result, append([]int{}, current...))
			return
		}

		for k, v := range m {
			if !v {
				continue
			}
			current = append(current, k)
			m[k] = false // Mark as used
			backtrack(index+1, current)
			current = current[:len(current)-1]
			m[k] = v // Restore the state
		}

	}

	backtrack(0, []int{})
	return result
}

func main() {
	nums := []int{1, 2, 3}
	result := permute(nums)
	for _, combination := range result {
		fmt.Println(combination)
	}
}
