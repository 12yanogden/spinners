package spinner

import (
	"strings"
	"testing"
)

func TestSpinner(t *testing.T) {
	active := true
	pattern := `|/-\`
	spinner := New(
		&active,
		pattern,
		1,
		100,
	)
	repetitions := 3
	expected := strings.Repeat(pattern, repetitions)
	actual := ""

	for range repetitions * len(pattern) {
		actual += spinner.Repeat()
	}

	if expected != actual {
		t.Fatalf("\nExpected:\t%#v\nActual:\t\t%#v\n", expected, actual)
	}
}
