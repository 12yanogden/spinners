package braille

import (
	"fmt"
	"testing"
	"time"
)

func TestBraille(t *testing.T) {
	active := true
	braille := Braille{}

	braille.Init(&active)

	go doTask(&active)

	for active {
		fmt.Printf("\r%s", braille.Play())
	}

	fmt.Println()
}

func doTask(active *bool) {
	time.Sleep(time.Duration(3) * time.Second)

	*active = false
}
