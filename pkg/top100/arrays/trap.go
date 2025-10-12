package arrays

// NO.42
func trap(height []int) int {

	n := len(height)
	//从左向右遍历， 取第i节点及期左侧最大值, 从右向左遍历，取i节点及期右侧最大值, 则公式为: dp[i] = min(leftMax[i], rightMax[i]) -height[i]
	leftMax, rightMax := make([]int, n), make([]int, n)
	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]
	for i := 1; i <= n-1; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	var ans int
	for i := 0; i < n; i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}

	return ans
}
