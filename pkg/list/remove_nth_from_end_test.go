package list

import (
	"fmt"
	"leetcode-golang/pkg/entity"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {

	head := entity.BuildListNode([]int{1, 2, 3, 4, 5})
	p := removeNthFromEnd(head, 2)
	fmt.Println(p)
}
