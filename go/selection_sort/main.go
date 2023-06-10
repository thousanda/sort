package main

import "fmt"

// simple main function to test selection sort
func main() {
	// create multiple slices of integers
	itemsList := [][]int{
		{5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5},
		{1, 3, 2, 5, 4},
		{5, 1, 4, 2, 3},
	}

	for _, items := range itemsList {
		// sort the slice
		SelectionSort(items)

		// print the sorted slice
		fmt.Println(items)
	}
}

func SelectionSort(items []int) {
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
