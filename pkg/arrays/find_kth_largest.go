package arrays

import "container/heap"

// NO.215  使用堆

type maxHp struct {
	a []int
}

func (h maxHp) Len() int {
	return len(h.a)
}

func (h maxHp) Swap(i, j int) {
	h.a[i], h.a[j] = h.a[j], h.a[i]
}

func (h maxHp) Less(i, j int) bool {
	return h.a[i] > h.a[j]
}

func (h *maxHp) Push(x interface{}) {
	(*h).a = append((*h).a, x.(int))
}

func (h *maxHp) Pop() interface{} {
	a := (*h).a
	v := a[len(a)-1]
	(*h).a = a[:len(a)-1]
	return v
}

// NO.215
func findKthLargest(nums []int, k int) int {
	q := &maxHp{
		a: make([]int, len(nums)),
	}
	for i := 0; i < len(nums); i++ {
		q.a[i] = nums[i]
	}

	heap.Init(q)

	for i := 0; i < k-1; i++ {
		heap.Pop(q)
	}
	return q.a[0]
}
