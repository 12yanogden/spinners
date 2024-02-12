package wave

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestSpinner(t *testing.T) {
	active := true
	wave := New(&active)
	repetitions := 4
	expected := strings.Repeat(wave.Pattern(), repetitions)
	actual := ""

	wave.setWidth(1)
	wave.setFrequency(10)

	for range (repetitions / 2) * len(wave.Pattern()) {
		actual += wave.Play()
	}

	if expected != actual {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v\n", expected, actual)
	}
}

func TestWave(t *testing.T) {
	active := true
	wave := New(&active)

	go doTask(&active)

	for active {
		fmt.Printf("\r%s", wave.Play())
	}

	fmt.Println()
}

func doTask(active *bool) {
	time.Sleep(time.Duration(3) * time.Second)

	*active = false
}
