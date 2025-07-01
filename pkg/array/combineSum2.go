package main

//func main() {
//	candidates := []int{2, 3, 6, 7}
//	target := 7
//	result := combinationSum(candidates, target)
//	for _, combination := range result {
//		fmt.Println(combination)
//	}
//}

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
