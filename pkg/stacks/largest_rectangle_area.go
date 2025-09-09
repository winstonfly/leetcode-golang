package stacks

// NO.84 单调栈
func largestRectangleArea(heights []int) int {
	n := len(heights)
	left, right := make([]int, n), make([]int, n)
	stack := []int{}
	for i := 0; i < n; i++ {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			left[i] = -1
		} else {
			left[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	stack = []int{}
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= heights[i] {
			stack = stack[:len(stack)-1]
		}

		if len(stack) == 0 {
			right[i] = n
		} else {
			right[i] = stack[len(stack)-1]
		}

		stack = append(stack, i)
	}

	var ans int
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*heights[i])
	}

	return ans
}
