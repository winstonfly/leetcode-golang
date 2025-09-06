package tree

// NO.230
func kthSmallest(root *TreeNode, k int) int {
	var ans []int
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		//中序遍历
		inorder(root.Left)
		ans = append(ans, root.Val)
		inorder(root.Right)
	}

	inorder(root)
	return ans[k-1]
}
