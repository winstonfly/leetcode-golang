package tree

// NO.114
func flatten(root *TreeNode) {

	var list []*TreeNode
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		list = append(list, root)
		inorder(root.Left)
		inorder(root.Right)
	}

	inorder(root)

	for i := 1; i < len(list); i++ {
		prev, curr := list[i-1], list[i]
		prev.Left, prev.Right = nil, curr
	}
}
