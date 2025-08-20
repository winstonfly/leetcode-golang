package tree

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		tests := []int{1, 2, 3, 4, 5}
		root := BuildTreeNode(tests)
		v := diameterOfBinaryTree(root)
		t.Log(v)
	})

	t.Run("test2", func(t *testing.T) {
		tests := []int{4, -7, -3, -1, -1, -9, -3, 9, -7, -4, -1, 6, -1, -6, -6, -1, -1, 0, 6, 5, -1, 9, -1, -1, -1, -4, -1, -1, -1, -2}
		root := BuildTreeNode(tests)
		v := diameterOfBinaryTree(root)
		t.Log(v)
	})
}
