package hp

import (
	"container/heap"
	"sort"
)

// NO.295
type MedianFinder struct {
	//小于中位数堆是大顶堆，大于中位数的是小顶堆, minHeap是小顶堆, maxHeap也是小顶堆，但存储的是负数，则变成了大顶堆
	minHeap, maxHeap hp
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: hp{IntSlice: []int{}},
		maxHeap: hp{IntSlice: []int{}},
	}
}

// 存储负数，方便两个堆移动
func (mf *MedianFinder) AddNum(num int) {
	maxQ, minQ := &mf.maxHeap, &mf.minHeap

	//num与中位数比较，如果是小于中位数，则放入大顶堆，否则放入小顶堆，正常情况下小顶堆比大顶堆元素相等或者多一个
	if minQ.Len() == 0 || num < minQ.IntSlice[0] {
		//放入大顶堆
		heap.Push(maxQ, -num)
		if maxQ.Len() > minQ.Len() {
			heap.Push(minQ, heap.Pop(maxQ).(int)*-1)
		}
	} else {
		//放入小顶堆
		heap.Push(minQ, num)
		if minQ.Len() > maxQ.Len()+1 {
			heap.Push(maxQ, heap.Pop(minQ).(int)*-1)
		}
	}
}

func (mf *MedianFinder) FindMedian() float64 {
	maxQ, minQ := &mf.maxHeap, &mf.minHeap
	if minQ.Len() > maxQ.Len() {
		return float64(minQ.IntSlice[0])
	} else {
		return float64(minQ.IntSlice[0]-maxQ.IntSlice[0]) / 2
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
