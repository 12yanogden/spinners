package spinner

import "time"

type Spinner struct {
	Active    *bool
	Index     int
	Pattern   []string
	Width     int
	Frequency int
}

func (spinner *Spinner) Init(active *bool, pattern []string, width int, frequency int) {
	spinner.Active = active
	spinner.Index = 0
	spinner.Pattern = pattern
	spinner.Width = width
	spinner.Frequency = frequency
}

func (spinner *Spinner) Repeat() string {
	out := spinner.Pattern[(spinner.Index)%len(spinner.Pattern)]
	spinner.Index++

	remainingCount := spinner.Width - 1
	for i := 0; i < remainingCount; i++ {
		out += spinner.Pattern[(i+spinner.Index)%len(spinner.Pattern)]
	}

	time.Sleep(time.Duration(spinner.Frequency) * time.Millisecond)

	return out
}
