package main

import (
	"fmt"
	"github.com/thousanda/sort/go/selection_sort"
	"github.com/thousanda/sort/go/timer"
)

// 入力なしで試しに動かす

func main() {
	// create multiple slices of integers
	itemsList := [][]int64{
		{5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5},
		{1, 3, 2, 5, 4},
		{5, 1, 4, 2, 3},
	}

	timer := timer.Timer{}

	// 計測スタート
	timer.Start()

	// 計測したい処理
	for _, items := range itemsList {
		// sort the slice
		selection_sort.SelectionSort(items)

		// print the sorted slice
		fmt.Println(items)
	}

	// 計測終了
	timer.Stop()
	// 計測結果出力
	timer.Print()
}
