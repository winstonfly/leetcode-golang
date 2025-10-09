package dp

// NO.322
func coinChange(coins []int, amount int) int {
	f := make([]int, amount+1)
	f[0] = 0

	for i := 1; i <= amount; i++ {
		f[i] = amount + 1
		for _, c := range coins {
			if i >= c {
				f[i] = min(f[i], f[i-c]+1)
			}
		}

	}

	if f[amount] > amount {
		return -1
	}

	return f[amount]

}
