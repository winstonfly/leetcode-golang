package tree

import "math"

// NO.124
// 核心点是：最大路径取决于当前节点的值与当前节点的最大贡献值
func maxPathSum(root *TreeNode) int {
	var maxSum int = math.MinInt64
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftGain := max(0, maxGain(node.Left))
		rightGain := max(0, maxGain(node.Right))
		price := leftGain + rightGain + node.Val
		maxSum = max(maxSum, price)
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}
