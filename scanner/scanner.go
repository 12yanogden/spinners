package scanner

import (
	"github.com/12yanogden/spinners/spinner"
)

type Scanner struct {
	spinner spinner.Spinner
}

func New(active *bool) Scanner {
	return Scanner{
		spinner: spinner.New(
			active,
			`⎽⎼―⎻⎺⎻―⎼`,
			1,
			100,
		),
	}
}

func (s *Scanner) Play() string {
	return s.spinner.Repeat()
}
