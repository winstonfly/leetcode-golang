package tree

// NO.105
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	//前序第一个节点是根节点, 中序遍历第一个节点为左子树
	index := 0

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			index = i
			break
		}
	}

	root := &TreeNode{Val: preorder[0]}
	if len(inorder[:index]) > 0 {
		root.Left = buildTree(preorder[1:], inorder[:index])
	}

	if len(inorder[index+1:]) > 0 {
		root.Right = buildTree(preorder[index+1:], inorder[index+1:])
	}

	return root
}
