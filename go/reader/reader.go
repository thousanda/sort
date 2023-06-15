package reader

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Reader struct {
	sc *bufio.Scanner
}

func NewReader(f *os.File) *Reader {
	scanner := bufio.NewScanner(f)
	return &Reader{
		sc: scanner,
	}
}

func (r *Reader) Read() (int64, []int64, error) {
	// Read N
	if !r.sc.Scan() {
		if r.sc.Err() != nil {
			return 0, nil, r.sc.Err()
		}
		// EOF
		return 0, nil, fmt.Errorf("unexpected EOF while reading N")
	}
	nStr := strings.TrimSpace(r.sc.Text())
	if nStr == "" {
		return 0, nil, fmt.Errorf("no input for N")
	}
	n, err := strconv.ParseInt(nStr, 10, 64)
	if err != nil {
		return 0, nil, err
	}

	// Read the list of integers
	if !r.sc.Scan() {
		if r.sc.Err() != nil {
			return 0, nil, r.sc.Err()
		}
		// EOF
		return 0, nil, fmt.Errorf("unexpected EOF while reading integers")
	}
	integers := r.sc.Text()

	// Split the integers by space and convert them to int
	intSlice := make([]int64, n)
	integerStrings := strings.Split(integers, " ")
	for i, s := range integerStrings {
		intSlice[i], err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			return 0, nil, err
		}
	}

	return n, intSlice, nil
}
