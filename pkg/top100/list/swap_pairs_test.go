package list

import "testing"

func TestSwapPairs(t *testing.T) {
	head := BuildListNode([]int{1, 2, 3, 4})
	p := swapPairs(head)
	for p != nil {
		t.Log(p.Val)
		p = p.Next
	}
}
