package sort

// BubbleSort sorts an array using the bubble sort algorithm.

func BubbleSort(arr []int) {
	n := len(arr)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // Swap the elements
			}
		}
	}

}
