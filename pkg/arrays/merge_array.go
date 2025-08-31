package arrays

import "slices"

//NO.56

func merge(intervals [][]int) [][]int {

	row := len(intervals)
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	ans := [][]int{intervals[0]}

	for i := 0; i < row; i++ {
		current := intervals[i]
		last := ans[len(ans)-1]

		if current[0] <= last[1] { //如果当前元素的start小于等于上一个元素的end，说明有交集
			last[1] = max(last[1], current[1])
		} else {
			ans = append(ans, current)
		}
	}
	return ans
}
