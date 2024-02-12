package line

import (
	"github.com/12yanogden/spinners/spinner"
)

type Line struct {
	spinner spinner.Spinner
}

func New(active *bool) Line {
	return Line{
		spinner: spinner.New(
			active,
			`⎽⎼―⎻⎺ `,
			1,
			100,
		),
	}
}

func (l *Line) Play() string {
	return l.spinner.Repeat()
}
