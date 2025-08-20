package tree

import "testing"

func TestInorderTraversal(t *testing.T) {

	t.Run("inorderTraversal", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		}

		result := inorderTraversal(root)
		expected := []int{4, 2, 5, 1, 3}
		if len(result) != len(expected) {
			t.Errorf("expected length %d, got %d", len(expected), len(result))
			return
		}

		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("expected %d at index %d, got %d", expected[i], i, result[i])
			}
		}
	})
}

func TestInbackTraversal(t *testing.T) {
	t.Run("inbackTraversal", func(t *testing.T) {
		root := &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 5,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		}

		result := inbackTraversal(root)
		expected := []int{4, 5, 2, 3, 1}
		if len(result) != len(expected) {
			t.Errorf("expected length %d, got %d", len(expected), len(result))
			return
		}

		for i := range result {
			if result[i] != expected[i] {
				t.Errorf("expected %d at index %d, got %d", expected[i], i, result[i])
			}
		}
	})
}
