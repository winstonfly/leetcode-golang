package tree

// NO.543
// 左右子树同高时最长，p, q一个遍历左子树，一个右子树
func diameterOfBinaryTree(root *TreeNode) int {

	maxLen := 0
	var depthFn func(node *TreeNode) int
	depthFn = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		return max(depthFn(node.Left), depthFn(node.Right)) + 1
	}

	var compute func(node *TreeNode)
	var queue []*TreeNode
	queue = append(queue, root)
	compute = func(node *TreeNode) {
		if node == nil {
			return
		}

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			if curr == nil {
				continue
			}
			leftDepth := depthFn(curr.Left)
			rightDepth := depthFn(curr.Right)
			maxLen = max(maxLen, leftDepth+rightDepth)
			queue = append(queue, curr.Left, curr.Right)
		}
	}

	compute(root)

	return maxLen
}
