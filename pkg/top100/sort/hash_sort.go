package sort

// HashSort sorts an array using the hash sort algorithm.
// It uses a map to count the occurrences of each element and then reconstructs the sorted array
// based on the counts.
func HashSort(arr []int) []int {
	counts := make(map[int]int)
	maxCount := 0
	for _, num := range arr {
		maxCount = max(num, maxCount)
		counts[num]++
	}

	sorted := make([]int, 0, len(arr))
	for i := 0; i < maxCount; i++ {
		if count, exists := counts[i]; exists {
			for j := 0; j < count; j++ {
				sorted = append(sorted, i)
			}
		}
	}

	return sorted
}
