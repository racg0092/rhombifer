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
	ValuesDilimiter string

	// when true child command has visibility to the flag
	Inherited bool
}
