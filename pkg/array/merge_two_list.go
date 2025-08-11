package main

import "fmt"

type ListNode2 struct {
	Val  int
	Next *ListNode2
}

func merge_two_lists(list1 *ListNode2, list2 *ListNode2) *ListNode2 {
	result := new(ListNode2)
	head := result
	for {
		if list1 == nil {
			head.Next = list2
			break
		}

		if list2 == nil {
			head.Next = list1
			break
		}

		if list1.Val <= list2.Val {
			next := list1
			list1 = list1.Next
			next.Next = nil

			head.Next = next
			head = head.Next
		} else {
			next := list2
			list2 = list2.Next
			next.Next = nil
			head.Next = next
			head = head.Next
		}

	}

	return result.Next
}

func main() {
	// Example usage
	list1 := &ListNode2{Val: 1, Next: &ListNode2{Val: 2, Next: &ListNode2{Val: 4}}}
	list2 := &ListNode2{Val: 1, Next: &ListNode2{Val: 3, Next: &ListNode2{Val: 4}}}

	mergedList := merge_two_lists(list1, list2)

	// Print merged list
	fmt.Println(mergedList)
}
