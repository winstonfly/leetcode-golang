### 滑动窗口最大值

#### 题目
![滑动窗口最大值](../../../images/max_slide_window.png)
### 思考
- 题目要求在一个数组中，找到每个滑动窗口的最大值， 这个滑动窗口的大小是k，直观印象可以转为一个大顶堆，维护k个元素的堆，堆顶就是最大值， 但是堆的插入和删除都是O(logk)的复杂度，整体复杂度是O(nlogk)，不是最优解，查看官解，加入了一个单调队列，
队列里维护了当前窗口的最大值，队列里存的是数组的下标，队列头部是当前窗口的最大值下标，队列尾部是当前窗口的最小值下标，队列里的元素是单调递减的， 这样每次滑动窗口时，只需要比较新加入的元素和队列尾部元素的大小，如果新加入的元素大于等于队列尾部元素，就把队列尾部元素弹出，直到新加入的元素小于队列尾部元素或者队列为空，然后把新加入的元素下标加入到队列尾部，这样就保证了队列里元素是单调递减的，同时还要判断队列头部元素是否在当前窗口内，如果不在，就把队列头部元素弹出，这样每次滑动窗口时，队列头部元素就是当前窗口的最大值。

### 代码
- 方法一： 大顶堆
```go

var a []int

type hp struct {
	sort.IntSlice
}

func (h hp) Less(i, j int) bool { return a[h.IntSlice[i]] > a[h.IntSlice[j]] } //最大堆

func (h *hp) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() any {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

func maxSlidingWindow1(nums []int, k int) []int {
	var ans []int
	h := &hp{make([]int, k)}
	a = nums

	for i := 0; i < k; i++ {
		h.IntSlice[i] = i
	}
	heap.Init(h)

	ans = append(ans, nums[h.IntSlice[0]])
	for i := k; i < len(nums); i++ {
		heap.Push(h, i)
		for h.IntSlice[0] <= i-k {
			heap.Pop(h)
		}

		ans = append(ans, nums[h.IntSlice[0]])
	}

	return ans
}

```

- 方法二： 单调队列
```go   
func maxSlidingWindow(nums []int, k int) []int {

	var ans []int

	q := []int{}
	push := func(i int) {
		for len(q) > 0 && nums[i] >= nums[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, iß)
	}

	for i := 0; i < k; i++ {
		push(i)
	}

	ans = append(ans, nums[q[0]])

	for i := k; i < len(nums); i++ {
		push(i)
		if q[0] <= i-k {
			q = q[1:]
		}

		ans = append(ans, nums[q[0]])
	}
	return ans
}
```