package tree

import "math"

// NO.124
// 核心点是：最大路径取决于当前节点的值与当前节点的最大贡献值
func maxPathSum(root *TreeNode) int {

	var maxGain func(node *TreeNode) int
	maxSum := math.MinInt64
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		price := node.Val + leftGain + rightGain
		maxSum = max(maxSum, price)

		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)

	return maxSum
}
