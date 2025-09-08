package tree

import "testing"

func TestFlatten(t *testing.T) {
	t.Run("test1", func(t *testing.T) {

		tree := BuildTreeNode([]int{1, 2, 5, 3, 4, -1, 6})
		flatten(tree)
		t.Log(tree)
	})
}
