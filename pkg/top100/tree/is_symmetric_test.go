package tree

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {

	//t.Run("Test isSymmetric-0", func(t *testing.T) {
	//	tests := []int{1, 2, 2, 3, 4, 4, 3}
	//	node := BuildTreeNode(tests)
	//	if !isSymmetric(node) {
	//		t.Errorf("Expected tree to be symmetric, but it is not")
	//	} else {
	//		t.Logf("Tree is symmetric as expected")
	//	}
	//})

	//t.Run("Test isSymmetric", func(t *testing.T) {
	//	tests := []int{1, 2, 2, -1, 3, -1, 3}
	//	node := BuildTreeNode(tests)
	//	if !isSymmetric(node) {
	//		t.Errorf("Expected tree to be symmetric, but it is not")
	//	} else {
	//	}
	//})

	t.Run("Test isSymmetric", func(t *testing.T) {
		tests := []int{2, 3, 3, 4, -1, -1, 4, -1, 5, 5, -1, -1, 6, 6, -1, 7, 8, 8, 7, 9, 0, 0, 1, 1, 0, 0, 9}
		node := BuildTreeNode(tests)
		if isSymmetric(node) {
			t.Errorf("Expected tree to be symmetric, but it is not")
		}
	})
}

func TestIsSymmetric2(t *testing.T) {
	start := 2
	fmt.Println(1 << start)
}
