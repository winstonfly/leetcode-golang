package dp

// NO.279
func numSquares(n int) int {
	//f[i]代表数为i时完全平方数的最小个数

	f := make([]int, n+1)
	f[0] = 0

	for i := 1; i <= n; i++ {
		f[i] = n + 1
		for j := 1; j*j <= i; j++ {
			f[i] = min(f[i], f[i-j*j]+1)
		}
	}

	if f[n] > n {
		return -1
	}
	return f[n]
}
