package tree

import "testing"

func TestLowestCOmmonAncestor(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		tree := BuildTreeNode([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4})
		p, q := tree.Left, tree.Right

		ans := lowestCommonAncestor(tree, p, q)
		t.Log(ans)
	})
}
