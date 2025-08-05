package main

import "fmt"

func minCostClimbingStairs(cost []int) int {
	// cost的长度是n，则下票为0 到 n-1代表所有的楼梯， 顶部即是为n的楼梯的最小花费就是最终结果
	// 定义一个dp[n + 1]， 则求dp[n]就是最终最小花费, 从i-1花费cost[i-1]，从i-2花费cost[i-2]即到i
	n := len(cost)
	if n == 0 {
		return 0
	}

	if n == 1 {
		return cost[0]
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 0

	for i := 2; i <= n; i++ {
		dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}

	fmt.Println(dp)
	return dp[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
func main() {
	cost := []int{10, 15, 20}
	result := minCostClimbingStairs(cost)
	println(result) // Output: 15
}
