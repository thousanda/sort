package main

import (
	"fmt"
	"github.com/thousanda/sort/go/reader"
	"github.com/thousanda/sort/go/selection_sort"
	"github.com/thousanda/sort/go/timer"
	"os"
)

// 標準入力からデータを読み込んで実行する

func main() {
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
	selection_sort.SelectionSort(slc)
	t.Stop()

	// 結果出力
	fmt.Println(slc)
	t.Print()
}
