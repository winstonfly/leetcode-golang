package tree

// NO.437
func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, sum)
	res += pathSum(root.Left, sum)
	res += pathSum(root.Right, sum)
	return res
}

// 求以某个节点开始的和为targetSum的路径
func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}

	if root.Val == targetSum {
		res++
	}

	res += rootSum(root.Left, targetSum-root.Val)
	res += rootSum(root.Right, targetSum-root.Val)
	return
}
