package main

import (
	"fmt"
	"github.com/thousanda/sort/go/selection_sort/reader"
	"github.com/thousanda/sort/go/selection_sort/timer"
	"os"
)

func main() {
	// どちらかを実行する
	// calcRun()
	calcFromStdIn()
}

// calcRun 入力なしで試しに動かす
func calcRun() {
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
		SelectionSort(items)

		// print the sorted slice
		fmt.Println(items)
	}

	// 計測終了
	timer.Stop()
	// 計測結果出力
	timer.Print()
}

// 標準入力からデータを読み込んで実行する
func calcFromStdIn() {
	// 標準入力から読み込み
	f := os.Stdin
	r := reader.NewReader(f)
	_, slc, err := r.Read()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 計測とソート
	t := timer.Timer{}
	t.Start()
	SelectionSort(slc)
	t.Stop()

	// 結果出力
	fmt.Println(slc)
	t.Print()
}

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
