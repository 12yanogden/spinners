package fader

import (
	"github.com/12yanogden/spinners/spinner"
)

type Fader struct {
	spinner spinner.Spinner
}

func New(active *bool) Fader {
	return Fader{
		spinner: spinner.New(
			active,
			`   ░▒▓██████▓▒░   `,
			1,
			150,
		),
	}
}

func (f *Fader) Play() string {
	return f.spinner.Repeat()
}
