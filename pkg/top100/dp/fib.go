package dp

import "fmt"

// 动态规划
// dp[i] 代表第i个斐波那契数
// dp[i] = dp[i-1] + dp[i-2] i >= 2
// dp[0] = 0
// dp[1] = 1
func fib(n int) int {
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	if n < 2 {
		return n
	}

	prev, curr := 0, 1
	for i := 2; i < n+1; i++ {
		next := prev + curr
		prev = curr
		curr = next
		dp[i] = curr
		fmt.Printf("i: %d, prev: %d, curr: %d, next: %d\n", i, prev, curr, next)
	}

	fmt.Println(dp)
	return curr
}
