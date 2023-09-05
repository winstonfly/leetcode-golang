package sort

// InsertSort 插入排序
// 原理: 将待排序的数组分为两个区间，未排序区间和已排序区间， 从未排序区间里拿出数据， 放到已排序区间里，保证数据有序
// 插入排序时间复杂度是O(n2)， 不需要额外的内存，是原地的排序算法
func InsertSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	// 从未排序区间选择元素
	for i := 1; i < len(a); i++ {
		value := a[i]
		j := i - 1
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j]
			} else {
				break
			}
		}

		a[j+1] = value
	}

	return a
}

// BubbleSort 冒泡排序
func BubbleSort(a []int, n int) []int {
	if n <= 1 {
		return a
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[i] > a[j] {
				tmp := a[j]
				a[j] = a[i]
				a[i] = tmp
			}
		}
	}
	return a
}
