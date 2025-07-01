package main

import "fmt"

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	result := combinationSum(candidates, target)
	for _, combination := range result {
		fmt.Println(combination)
	}
}

func combinationSum1(candidates []int, target int) [][]int {
	if len(candidates) == 0 || target < 0 {
		return nil
	}

	var result [][]int
	var path []int

	var combin func(candidates []int, target int, path []int)

	combin = func(candidates []int, target int, path []int) {
		if target == 0 {
			result = append(result, append([]int(nil), path...)) // Append a copy of the path
			return
		}

		if len(candidates) == 0 || target < 0 {
			return
		}

		for i := 0; i < len(candidates); i++ {
			if candidates[i] > target {
				continue
			}

			// Include the current candidate and continue searching
			path = append(path, candidates[i])
			combin(candidates[i:], target-candidates[i], path)
			// Backtrack: remove the last candidate added
			path = path[:len(path)-1]
		}

	}

	combin(candidates[0:], target, path)

	return result
}
