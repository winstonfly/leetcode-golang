package sort

// insertSort 插入排序
// 原理: 将待排序的数组分为两个区间，未排序区间和已排序区间， 从未排序区间里拿出数据， 放到已排序区间里，保证数据有序
// 插入排序时间复杂度是O(n2)， 不需要额外的内存，是原地的排序算法
func insertSort(a []int) []int {
	for i := 1; i < len(a); i++ {
		j := i
		value := a[j]
		for ; j > 0 && value < a[j-1]; j-- {
			a[j] = a[j-1] //移动位置
		}
		a[j] = value //将待插入的值放到正确的位置
	}

	return a
}
