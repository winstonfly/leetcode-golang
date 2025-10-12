package tree

// NO.114
func flatten1(root *TreeNode) {

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

func flatten(root *TreeNode) {
	var inorder func(root *TreeNode)
	var ans []*TreeNode

	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		ans = append(ans, root)
		inorder(root.Left)
		inorder(root.Right)
	}

	inorder(root)

	for i := 0; i < len(ans)-1; i++ {
		ans[i].Left = nil
		ans[i].Right = ans[i+1]
	}
}
