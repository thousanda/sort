package main

import (
	"fmt"
	"github.com/thousanda/sort/go/merge_sort"
	"github.com/thousanda/sort/go/util"
	"os"
)

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
	sorted := merge_sort.MergeSort(arr)

	// 計測終了
	timer.Stop()

	// ソート結果出力
	fmt.Println(sorted)

	// 計測結果出力
	timer.Print()

}
