package scanner

import (
	"fmt"
	"testing"
	"time"
)

func TestScanner(t *testing.T) {
	active := true
	scanner := New(&active)

	go doTask(&active)

	for active {
		fmt.Printf("\r%s", scanner.Play())
	}

	fmt.Println()
}

func doTask(active *bool) {
	time.Sleep(time.Duration(3) * time.Second)

	*active = false
}
