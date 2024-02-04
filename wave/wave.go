package wave

import (
	"github.com/12yanogden/spinners/spinner"
)

type Wave struct {
	spinner spinner.Spinner
}

func (w *Wave) Init(active *bool) {
	w.spinner = spinner.Spinner{}

	w.spinner.Init(
		active,
		[]string{"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▁", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		6,
		100)
}

func (w *Wave) Play() string {
	return w.spinner.Repeat()
}
