package selection_sort

func SelectionSort(items []int64) {
	// iterate over the slice
	for i := 0; i < len(items); i++ {
		// find the minimum value in the unsorted slice
		min := i
		for j := i + 1; j < len(items); j++ {
			if items[j] < items[min] {
				min = j
			}
		}

		// swap the minimum value with the first value in the unsorted slice
		items[i], items[min] = items[min], items[i]
	}
}
