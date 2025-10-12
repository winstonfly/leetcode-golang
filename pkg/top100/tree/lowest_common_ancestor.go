package tree

// NO.236
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var ans *TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {

		if node == nil {
			return
		}

		if pathAncestor(node, p) && pathAncestor(node, q) {
			ans = node
		}

		dfs(node.Left)
		dfs(node.Right)
	}

	dfs(root)
	return ans
}

func pathAncestor(node, target *TreeNode) bool {
	if node == nil {
		return false
	}

	if node == target {
		return true
	}

	return pathAncestor(node.Left, target) || pathAncestor(node.Right, target)
}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	if right != nil {
		return right
	}

	return nil
}
