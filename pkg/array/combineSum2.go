package main

import (
	"fmt"
)

func main() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	result := combinationSum2(candidates, target)
	for _, combination := range result {
		fmt.Println(combination)
	}
}

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 || target < 0 {
		return nil
	}

	result := [][]int{}
	var backtrack func(start int, current []int, remaining int)

	backtrack = func(start int, current []int, remaining int) {
		if remaining == 0 {
			result = append(result, append([]int{}, current...))
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > remaining {
				continue
			}
			current = append(current, candidates[i])
			backtrack(i, current, remaining-candidates[i])
			current = current[:len(current)-1]
		}
	}

	backtrack(0, []int{}, target)
	return result
}

// 给定一个候选人编号的集合 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。
// candidates 中的每个数字在每个组合中只能使用 一次 。
func combinationSum2(candidates []int, target int) [][]int {

	var (
		result [][]int
	)
	var backtrack func(tokens []int, path []int, target int)
	backtrack = func(tokens []int, path []int, target int) {
		if target == 0 {
			result = append(result, append([]int(nil), path...))
			return
		}

		if len(tokens) == 0 || target < 0 {
			return
		}

		for i := 0; i < len(tokens); i++ {
			if tokens[i] > target {
				i++
				continue
			}

			if i < len(tokens) {
				path = append(path, tokens[i])
				backtrack(tokens[i:], path, target-tokens[i])
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(candidates, []int{}, target)
	return result
}
