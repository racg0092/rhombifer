package rhombifer

type Flag struct {
	// flag name
	Name string

	// short drescriptio
	Short string

	// long description
	Long string

	// short format
	ShortFormat string

	// Long Format
	LongFormat string

	Required bool

	ValuesDilimiter string
}

func (f Flag) ShortHand() string {
	return f.Name[0:1]
}
