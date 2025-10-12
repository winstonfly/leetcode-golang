package tree

// NO.543
func diameterOfBinaryTree(root *TreeNode) int {

	if root == nil {
		return 0
	}

	var ans int
	m := make(map[*TreeNode]int)
	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		if v, ok := m[node]; ok {
			return v
		}

		leftDepth := depth(node.Left)
		rightDepth := depth(node.Right)
		m[node] = max(leftDepth, rightDepth) + 1
		return max(leftDepth, rightDepth) + 1
	}

	var queue []*TreeNode
	queue = append(queue, root)
	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		if q.Left != nil {
			queue = append(queue, q.Left)
		}

		if q.Right != nil {
			queue = append(queue, q.Right)
		}

		leftDepth := depth(q.Left)
		rightDepth := depth(q.Right)
		ans = max(ans, leftDepth+rightDepth)
	}

	return ans
}
