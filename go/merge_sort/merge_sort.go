package merge_sort

// MergeSort function to sort a slice of integers
func MergeSort(items []int) []int {
	// if the slice has only one element, return it
	if len(items) < 2 {
		return items
	}

	// create a slice to hold the sorted values to return
	ans := make([]int, len(items))

	// split the slice into two halves and sort each half recursively
	left := MergeSort(items[:len(items)/2])
	right := MergeSort(items[len(items)/2:])

	// merge the two sorted halves
	i, j := 0, 0
	for k := 0; k < len(items); k++ {
		// if the left half is exhausted, copy the remaining values from the right half.
		// if the right half is exhausted, copy the remaining values from the left half.
		// Otherwise, compare the values in the left and right halves and copy the smaller value to the temporary slice
		if i < len(left) && (j == len(right) || left[i] < right[j]) {
			ans[k] = left[i]
			i++
		} else {
			ans[k] = right[j]
			j++
		}
	}
	return ans
}
