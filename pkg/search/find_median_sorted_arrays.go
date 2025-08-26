package search

// NO.4
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	i, j := 0, 0
	var values []int
	l := m + n
	for i < m && j < n {
		if nums1[i] <= nums2[j] {
			values = append(values, nums1[i])
			i++
		} else {
			values = append(values, nums2[j])
			j++
		}
	}

	if i < m {
		values = append(values, nums1[i:]...)
	}

	if j < n {
		values = append(values, nums2[j:]...)
	}

	if (m+n)%2 == 0 {
		return float64(values[l/2-1]+values[l/2]) / 2
	} else {
		return float64(values[l/2])
	}
}
