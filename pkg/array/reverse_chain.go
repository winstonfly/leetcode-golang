package main

import "fmt"

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

// 1 -> 2-> 3 -> 4 -> 5
func reverseList(head *ListNode1) *ListNode1 {
	for head != nil && head.Next != nil {
		current := head
		var pre *ListNode1

	}
	return head
}

func BuildListNode(nums []int) *ListNode1 {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode1{Val: nums[0]}
	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode1{Val: nums[i]}
		current = current.Next
	}

	return head
}

func main() {
	fmt.Println(reverseList(BuildListNode([]int{1, 2, 3, 4, 5})))
}
