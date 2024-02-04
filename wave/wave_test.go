package wave

import (
	"fmt"
	"testing"
	"time"
)

func TestWave(t *testing.T) {
	active := true
	wave := Wave{}

	wave.Init(&active)

	go doTask(&active)

	for active {
		fmt.Printf("\r%s", wave.Play())
	}

	fmt.Println()
}

func doTask(active *bool) {
	time.Sleep(time.Duration(10) * time.Second)

	*active = false
}
