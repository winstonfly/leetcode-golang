package tree

import (
	"container/heap"
	"sort"
)

// NO.295
type MedianFinder struct {
	queMin, queMax hp
}

func Constructor() MedianFinder {
	return MedianFinder{}
}

func (mf *MedianFinder) AddNum(num int) {
	minQ, maxQ := &mf.queMin, &mf.queMax
	if minQ.Len() == 0 || num <= -minQ.IntSlice[0] {
		//将值放放小顶堆
		heap.Push(minQ, -num)
		if maxQ.Len()+1 < minQ.Len() {
			//将小顶堆中大于中位数的移到大顶堆
			heap.Push(maxQ, -heap.Pop(minQ).(int))
		}
	} else {
		heap.Push(maxQ, num)
		if maxQ.Len() > minQ.Len() {
			heap.Push(minQ, -heap.Pop(maxQ).(int))
		}
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	minQ, maxQ := &mf.queMin, &mf.queMax
	if minQ.Len() > maxQ.Len() {
		return float64(-minQ.IntSlice[0])
	} else {
		return float64(maxQ.IntSlice[0]-minQ.IntSlice[0]) / 2
	}
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
