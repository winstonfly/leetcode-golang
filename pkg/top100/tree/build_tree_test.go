package tree

import "testing"

func TestBuildTree(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		preorder := []int{3, 9, 20, 15, 7}
		inorder := []int{9, 3, 15, 20, 7}
		tree := buildTree(preorder, inorder)
		t.Log(tree)
	})

	t.Run("test2", func(t *testing.T) {
		preorder := []int{1, 2}
		inorder := []int{1, 2}
		tree := buildTree(preorder, inorder)
		t.Log(tree)
	})
}
