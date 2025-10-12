package arrays

import "sort"

// NO.56
func merge(intervals [][]int) [][]int {

	var ans [][]int
	if len(intervals) == 0 {
		return ans
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < len(intervals); i++ {
		row := intervals[i]
		if len(ans) == 0 || ans[len(ans)-1][1] < row[0] {
			ans = append(ans, row)
		} else {
			ans[len(ans)-1][1] = max(row[1], ans[len(ans)-1][1])
		}
	}

	return ans
}
