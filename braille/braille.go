package braille

import (
	"github.com/12yanogden/spinners/spinner"
)

type Braille struct {
	spinner spinner.Spinner
}

func (b *Braille) Init(active *bool) {
	b.spinner = spinner.Spinner{}

	b.spinner.Init(active, []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}, 1, 100)
}

func (b *Braille) Play() string {
	return b.spinner.Repeat()
}
