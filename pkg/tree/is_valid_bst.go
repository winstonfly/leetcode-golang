package tree

import "math"

// NO.98
func isValidBST(root *TreeNode) bool {
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, lower, upper int) bool {
	if root == nil {
		return true
	}

	if root.Val <= lower || root.Val >= upper {
		return false
	}

	return helper(root.Left, lower, root.Val) && helper(root.Right, root.Val, upper)
}

// 向二叉搜索树中插入节点
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	// 如果树为空，创建新节点作为根
	if root == nil {
		return &TreeNode{Val: val}
	}

	// 根据BST特性插入节点
	if val < root.Val && val != -1 {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}

	return root
}
