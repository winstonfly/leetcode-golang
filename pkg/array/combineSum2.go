package main

import (
	"fmt"
	"slices"
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
	slices.Sort(candidates)
	var result [][]int
	//回溯解法
	var backtrack func(arr []int, sum, index int)
	backtrack = func(arr []int, sum, index int) {
		//找到终止条件
		if sum == target {
			result = append(result, append([]int{}, arr...)) // 这里使用append([]int{}, arr...)是为了创建一个新的切片，避免后续修改影响到结果
			return
		}

		if sum > target || index >= len(candidates) {
			return
		}

		if len(candidates) == 0 {
			return
		}

		for i := index; i < len(candidates); i++ {
			if i > index && candidates[i] == candidates[i-1] {
				//跳过重复的元素
				continue
			}
			arr = append(arr, candidates[i])
			sum += candidates[i]
			backtrack(arr, sum, i+1)
			//回溯
			arr = arr[:len(arr)-1]
			sum -= candidates[i]
		}

	}

	backtrack([]int{}, 0, 0)
	return result
}
