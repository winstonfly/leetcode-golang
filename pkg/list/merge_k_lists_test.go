package list

import "testing"

func TestMergeKLists(t *testing.T) {
	//t.Run("mergeKLists", func(t *testing.T) {
	//	tests := []*ListNode{
	//		BuildListNode([]int{1, 4, 5}),
	//		BuildListNode([]int{1, 3, 4}),
	//		BuildListNode([]int{2, 6}),
	//	}
	//
	//	result := mergeKLists(tests)
	//	for result != nil {
	//		t.Log(result.Val)
	//		result = result.Next
	//	}
	//})

	t.Run("mergeKLists with empty list", func(t *testing.T) {
		tests := []*ListNode{
			BuildListNode([]int{2}),
			BuildListNode([]int{}),
			BuildListNode([]int{-1}),
		}
		result := mergeKLists(tests)
		if result != nil {
			t.Log(result.Val)
			result = result.Next
		}
	})

}
