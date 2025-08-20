package tree

func isSymmetric(root *TreeNode) bool {
	var check func(p, q *TreeNode) bool
	check = func(p, q *TreeNode) bool {
		if p == nil && q == nil {
			return true
		}

		if p == nil || q == nil {
			return false
		}

		return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
	}

	return check(root, root)
}
