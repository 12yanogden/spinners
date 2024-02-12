package shutters

import (
	"fmt"
	"testing"
	"time"
)

func TestShutters(t *testing.T) {
	active := true
	shutters := New(&active)

	go doTask(&active)

	for active {
		fmt.Printf("\r%s", shutters.Play())
	}

	fmt.Println()
}

func doTask(active *bool) {
	time.Sleep(time.Duration(3) * time.Second)

	*active = false
}
