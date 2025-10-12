package main

import "fmt"

func main() {
	nums1 := []int{2, 0}
	m := 1
	num2 := []int{1}
	n := 1

	copy(nums1[m:m+n], num2[:n])
	sort1(nums1)
	fmt.Println(nums1)
}

func merge1(nums1 []int, m int, nums2 []int, n int) {

	var num []int
	num = append(num, nums1[0:m]...)
	num = append(num, nums2[0:n]...)

	sort1(num)
	fmt.Println(num)

}

func sort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-1-i; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

func sort1(a []int) {
	for i := 0; i < len(a); i++ {
		tmp := a[i]
		for j := i; j < len(a); j++ {
			if a[i] > a[j] {
				tmp = a[j]
				a[j] = a[i]
				a[i] = tmp
			}
		}

	}
}

func sortedSquares(nums []int) []int {
	n := len(nums)
	i, j, k := 0, n-1, n-1

	arr := make([]int, n)
	for i <= j {
		a, b := nums[i]*nums[i], nums[j]*nums[j]
		if a > b {
			arr[k] = a
			i++
		} else {
			arr[k] = b
			j--
		}

		k--
	}

	return arr
}
