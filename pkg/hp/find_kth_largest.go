package hp

import "container/heap"

// NO.215

type maxHeap struct {
	nums []int
}

func (n *maxHeap) Len() int {
	return len(n.nums)
}

func (n *maxHeap) Less(i, j int) bool {
	return n.nums[i] > n.nums[j]
}

func (n *maxHeap) Swap(i, j int) {
	n.nums[i], n.nums[j] = n.nums[j], n.nums[i]
}

func (n *maxHeap) Push(x any) {
	n.nums = append(n.nums, x.(int))
}

func (n *maxHeap) Pop() any {
	v := n.nums[len(n.nums)-1]
	n.nums = n.nums[:len(n.nums)-1]
	return v
}

func findKthLargest(nums []int, k int) int {

	h := &maxHeap{nums: []int{}}
	for _, v := range nums {
		heap.Push(h, v)
	}
	heap.Init(h)

	for i := 0; i < k-1; i++ {
		heap.Pop(h)
	}

	return h.nums[0]
}
