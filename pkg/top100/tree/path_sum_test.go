package tree

import "testing"

func TestPathSum(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		tree := BuildTreeNode([]int{10, 5, -3, 3, 2, -1, 11, 3, -2, -1, 1})
		target := 8
		ans := pathSum(tree, target)
		t.Log(ans)
	})
}
