package braille

import (
	"github.com/12yanogden/spinners/spinner"
)

type Braille struct {
	spinner spinner.Spinner
}

func New(active *bool) Braille {
	return Braille{
		spinner: spinner.New(
			active,
			`⣾⣽⣻⢿⡿⣟⣯⣷`,
			1,
			100,
		),
	}
}

func (b *Braille) Play() string {
	return b.spinner.Repeat()
}
