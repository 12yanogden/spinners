package wave

import (
	"github.com/12yanogden/spinners/spinner"
)

type Wave struct {
	spinner spinner.Spinner
}

func New(active *bool) Wave {
	return Wave{
		spinner: spinner.New(
			active,
			`▁▂▃▄▅▆▇█▇▆▅▄▃▁              `,
			6,
			100,
		),
	}
}

func (w *Wave) Play() string {
	return w.spinner.Repeat()
}

func (w *Wave) Pattern() string {
	return w.spinner.Pattern
}

func (w *Wave) Width() int {
	return w.spinner.Width
}

func (w *Wave) setWidth(width int) {
	w.spinner.Width = width
}

func (w *Wave) setFrequency(frequency int) {
	w.spinner.Frequency = frequency
}
