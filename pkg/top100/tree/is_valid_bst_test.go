package tree

import "testing"

func TestIsValidBst(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		nums := []int{5, 4, 6, -1, -1, 3, 7}
		var tree *TreeNode
		for _, v := range nums {
			tree = insertIntoBST(tree, v)
		}
		v := isValidBST(tree)
		if !v {
			t.Error("expected true")
		}
	})
}
