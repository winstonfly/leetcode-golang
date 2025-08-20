package main

import (
	"fmt"
)

// [2，-2，2，-2]  target = 0

func main() {
	//a := []int{1, 3, -4, 5, -6, 5}
	//target := -1
	//
	//res := addArray(a, target)
	//fmt.Println(res)

	a := []int{1, 3, 1, 3, 100}
	result := rob2(a)
	fmt.Println(result)

	a1 := []int{2, 3, 2}
	result1 := rob2(a1)
	fmt.Println(result1)

	a2 := []int{1, 2, 3, 1}
	result2 := rob2(a2)
	fmt.Println(result2)

	a3 := []int{2, 7, 9, 3, 1}
	result3 := rob2(a3)
	fmt.Println(result3)

	a4 := []int{1, 1, 1, 2}
	result4 := rob2(a4)
	fmt.Println(result4)

}

func addArray(candidates []int, target int) [][]int {

	var result [][]int

	for i := 0; i < len(candidates); i++ {
		sum := 0
		var index []int
		for j := i; j < len(candidates); j++ {
			sum += candidates[j]
			index = append(index, j)
			if sum == target {
				result = append(result, index)
			}
		}
	}

	return result
}

// 打家动舍
// 2, 1,1, 2
func rob(tokens []int) int {
	// dp[i] 代表偷到第i个家里的时候的值，是dp[i-2] + current, 或者是dp[i-1]
	// 两种，偷第i房间，或者偷i-1房间的
	dp := make([]int, len(tokens))

	// 1,2 ,3,1
	if len(tokens) == 0 {
		return 0
	}

	if len(tokens) == 1 {
		return tokens[0]
	}

	dp[0] = tokens[0]
	dp[1] = max(tokens[0], tokens[1])

	for i := 2; i < len(tokens); i++ {
		dp[i] = max(dp[i-2]+tokens[i], dp[i-1])
	}

	return dp[len(tokens)-1]
}

// 1， 5， 7， 9， 11， 1
// dp0 = 1, dp1 = 5, dp2 = 8, dp3 = 14, dp4 = 19, dp5 = 15
// 1, 2, 3, 1
// dp0 = 1, dp1 = 2, dp3 = 4, dp4 = 3

// 2, 3, 2

// 1,1,1,2
func rob2(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}
	//偷第一个时，不偷最后一个
	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}
	return max(r(nums[:n-1]), r(nums[1:n]))
}

func r(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[1], dp[0])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}

	return dp[len(nums)-1]
}
