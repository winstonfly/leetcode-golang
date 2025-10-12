package arrays

// 暴力解法
func dailyTemperature2(temperatures []int) []int {
	var result []int
	m := make(map[int]int, 1)
	for i := 0; i < len(temperatures); i++ {
		if i >= 1 {
			if v, ok := m[i-1]; ok && temperatures[i] == temperatures[i-1] {
				// 如果前一天已经有结果，直接使用
				if v == 0 {
					result = append(result, 0) // 如果前一天没有更高的温度，添加0
					m[i] = 0                   // 记录当前温度没有更高的天数差
				} else {
					result = append(result, v-1)
					m[i] = v - 1 // 记录当前温度的天数差
				}
				continue
			}
		}

		found := false
		for j := i + 1; j < len(temperatures); j++ {
			if temperatures[j] > temperatures[i] {
				result = append(result, j-i)
				m[i] = j - i // 记录当前温度的天数差
				found = true
				break
			}
		}

		if !found {
			result = append(result, 0) // 如果没有找到更高的温度，添加0
			m[i] = 0                   // 记录当前温度没有更高的天数差
		}
	}

	return result
}

// 单调栈解法
func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, len(temperatures))
	var stack []int

	for i := 0; i < len(temperatures); i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1] // 弹出栈顶元素
			ans[prevIndex] = i - prevIndex
		}

		stack = append(stack, i)
	}

	return ans
}
