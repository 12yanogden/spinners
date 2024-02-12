package spinner

import (
	"time"
	"unicode/utf8"
)

type Spinner struct {
	Active    *bool
	Index     int
	Pattern   string
	Width     int
	Frequency int
}

// Return a new instance of Spinner
func New(active *bool, pattern string, width int, frequency int) Spinner {
	return Spinner{
		Active:    active,
		Index:     0,
		Pattern:   pattern,
		Width:     width,
		Frequency: frequency,
	}
}

// Return the next iteration of the Spinner print window
func (s *Spinner) Repeat() string {
	index := s.Index
	out := ""

	// Assemble the print string
	for range s.Width {
		next := s.next(&index)
		out += next
	}

	s.next(&s.Index)

	// Wait a moment
	time.Sleep(time.Duration(s.Frequency) * time.Millisecond)

	// Return the print string
	return out
}

func (s *Spinner) next(index *int) string {
	runeValue, runeWidth := utf8.DecodeRuneInString(s.Pattern[*index:])

	*index = (*index + runeWidth) % len(s.Pattern)

	return string(runeValue)
}
