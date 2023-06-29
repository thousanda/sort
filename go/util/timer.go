package util

import (
	"fmt"
	"time"
)

type Timer struct {
	start time.Time
	stop  time.Time
}

func (t *Timer) Start() {
	t.start = time.Now()
}

func (t *Timer) Stop() {
	t.stop = time.Now()
}

func (t *Timer) Print() {
	elapsed := t.stop.Sub(t.start)
	fmt.Println("Execution time: ", elapsed)
}
