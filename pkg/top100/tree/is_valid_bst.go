package tree

import "math"

// NO.98
func isValidBST(root *TreeNode) bool {
	var check func(node *TreeNode, lower, upper int) bool

	check = func(node *TreeNode, lower, upper int) bool {
		if node == nil {
			return true
		}

		if node.Val <= lower || node.Val >= upper {
			return false
		}

		return check(node.Left, lower, node.Val) && check(node.Right, node.Val, upper)
	}

	return check(root, math.MinInt64, math.MaxInt64)
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

func isValidBST2(root *TreeNode) bool {
	var inorder func(node *TreeNode)
	ans := true

	var pre *TreeNode
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if pre != nil && node.Val <= pre.Val {
			ans = false
			return
		}
		pre = node
		inorder(node.Right)
	}

	inorder(root)
	return ans
}
