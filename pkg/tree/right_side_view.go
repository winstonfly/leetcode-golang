package tree

// NO.199
func rightSideView(root *TreeNode) []int {

	var ans []int
	if root == nil {
		return ans
	}

	start := []*TreeNode{root}
	ans = append(ans, root.Val)
	for len(start) > 0 {
		var queue []*TreeNode

		for _, node := range start {
			if node != nil {
				queue = append(queue, node.Left, node.Right)
			}
		}

		if len(queue) > 0 {
			for i := len(queue) - 1; i >= 0; i-- {
				node := queue[i]
				if node != nil {
					ans = append(ans, node.Val)
					break
				}
			}
		}

		start = queue
	}

	return ans
}
