package util

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func InputFile(filename string) (int, []int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, nil, fmt.Errorf("ファイルオープン失敗: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// 1行目を読み込み
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, nil, fmt.Errorf("nを整数にパースできなかったよ: %w", err)
	}

	// 2行目を読み込み
	scanner.Scan()
	strArr := strings.Fields(scanner.Text())
	arr := make([]int64, len(strArr))
	for i, v := range strArr {
		arr[i], err = strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, nil, fmt.Errorf("数列の中にint64にパースできないやつがあったよ, %s: %w", strArr[i], err)
		}
	}

	return n, arr, nil
}
