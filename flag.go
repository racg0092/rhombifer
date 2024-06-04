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

	//todo: not sure if to keep it
	// to define how the values would be parse
	ValuesDilimiter []string

	// if the flag required a value or not
	RequiresValue bool
}
