package list

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		head := BuildListNode([]int{1, 2, 3, 4, 5})
		p := removeNthFromEnd(head, 2)
		t.Log(p)
	})

}
