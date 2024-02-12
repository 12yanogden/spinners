package line

import (
	"fmt"
	"testing"
	"time"
)

func TestLine(t *testing.T) {
	active := true
	line := New(&active)

	go doTask(&active)

	for active {
		fmt.Printf("\r%s", line.Play())
	}

	fmt.Println()
}

func doTask(active *bool) {
	time.Sleep(time.Duration(3) * time.Second)

	*active = false
}
