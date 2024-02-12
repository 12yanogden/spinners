package shutters

import (
	"github.com/12yanogden/spinners/spinner"
)

type Shutters struct {
	spinner spinner.Spinner
}

func New(active *bool) Shutters {
	return Shutters{
		spinner: spinner.New(
			active,
			`▉▊▋▌▍▎▏▎▍▌▋▊▉`,
			6,
			100,
		),
	}
}

func (s *Shutters) Play() string {
	return s.spinner.Repeat()
}
