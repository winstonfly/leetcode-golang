package main

import "fmt"

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

// 1 -> 2-> 3 -> 4 -> 5
func reverseList(head *ListNode1) *ListNode1 {
	//创建哨兵节点
	dummy := &ListNode1{Next: head}
	var prev *ListNode1
	for dummy.Next != nil {
		//获取一个节点
		current := dummy.Next
		dummy.Next = current.Next //删除这个节点
		current.Next = prev
		prev = current //将当前节点放到前面
	}

	return prev
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
	l := reverseList(BuildListNode([]int{1, 1, 2, 1}))
	fmt.Println(l)
}
