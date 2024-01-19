package line

import (
	"github.com/12yanogden/spinners/spinner"
)

type Line struct {
	spinner spinner.Spinner
}

func (l *Line) Init(active *bool) {
	l.spinner = spinner.Spinner{}

	l.spinner.Init(active, []string{"|", "/", "-", `\`}, 1, 100)
}

func (l *Line) Play() string {
	return l.spinner.Repeat()
}
