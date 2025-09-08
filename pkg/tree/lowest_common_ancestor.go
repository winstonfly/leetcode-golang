package tree

// NO.236
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	var ans *TreeNode
	if pathAncestor(root, p) && pathAncestor(root, q) {
		ans = root
	}

	if tmp := lowestCommonAncestor(root.Left, p, q); tmp != nil {
		ans = tmp
	}
	if tmp := lowestCommonAncestor(root.Right, p, q); tmp != nil {
		ans = tmp
	}
	return ans
}

func pathAncestor(root *TreeNode, p *TreeNode) bool {
	if root == nil {
		return false
	}

	if root == p {
		return true
	}

	return pathAncestor(root.Left, p) || pathAncestor(root.Right, p)
}
