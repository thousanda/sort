package main

import (
	"fmt"
	"github.com/thousanda/sort/go/merge_sort"
	"github.com/thousanda/sort/go/timer"
)

func main() {
	timer := timer.Timer{}

	// 計測スタート
	timer.Start()

	// 計測したい処理
	runMergeSort()

	// 計測終了
	timer.Stop()
	// 計測結果出力
	timer.Print()
}

func runMergeSort() {
	// create multiple slices of integers
	itemsList := [][]int{
		{5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5},
		{1, 3, 2, 5, 4},
		{5, 1, 4, 2, 3},
	}

	for _, items := range itemsList {
		// sort the slice
		sorted := merge_sort.MergeSort(items)

		// print the sorted slice
		fmt.Println(sorted)
	}
}
