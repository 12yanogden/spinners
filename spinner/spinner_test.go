package spinner

import (
	"testing"
	"time"

	"github.com/12yanogden/intslices"
	"github.com/12yanogden/strslices"
)

func TestSpinner(t *testing.T) {
	var spinner Spinner
	pattern := []string{"|", "/", "-", `\`}
	repetitions := 3
	active := true
	expected := strslices.Repeat(pattern, repetitions)
	actual := []string{}

	spinner.Init(&active, pattern, 1, 100)

	for range intslices.Seq(1, len(pattern)*repetitions) {
		actual = append(actual, spinner.Repeat())
	}

	if !strslices.Equals(expected, actual) {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v\n", expected, actual)
	}
}

func sleepSeconds(seconds int) {
	time.Sleep(time.Duration(seconds) * time.Second)
}
