package fader

import (
	"fmt"
	"testing"
	"time"
)

func TestFader(t *testing.T) {
	active := true
	fader := New(&active)

	go doTask(&active)

	for active {
		fmt.Printf("\r%s", fader.Play())
	}

	fmt.Println()
}

func doTask(active *bool) {
	time.Sleep(time.Duration(10) * time.Second)

	*active = false
}
