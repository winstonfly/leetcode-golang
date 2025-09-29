package stacks

// NO.739
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	var stack []int
	for i := 0; i < len(temperatures); i++ {

		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}

		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temperatures[i] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans[prevIndex] = i - prevIndex

		}

		stack = append(stack, i)
	}

	return ans
}
