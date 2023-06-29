package main

import (
	"fmt"
	"github.com/thousanda/sort/go/selection_sort"
	"github.com/thousanda/sort/go/util"
	"os"
)

const ShouldPrintSorted = false

func main() {
	/* 引数のファイルからソート対象を読み込む */
	if len(os.Args) < 2 {
		fmt.Println("ファイル名ください")
		return
	}

	n, arr, err := util.InputFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)

	/* ソート実行と計測 */
	timer := util.Timer{}

	// 計測スタート
	timer.Start()

	// 計測したい処理
	selection_sort.SelectionSort(arr)

	// 計測終了
	timer.Stop()

	// ソート結果出力
	if ShouldPrintSorted {
		fmt.Println(arr)
	}

	// 計測結果出力
	timer.Print()
}
