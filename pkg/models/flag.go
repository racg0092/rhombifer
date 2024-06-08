package models

type Flag struct {
	// flag name
	Name string

	//Shot Desctiption
	Short string

	// Long Description
	Long string

	// Short Format for flag if any
	ShortFormat string

	// Is this a require flag
	Required bool

	//todo: not sure if to keep it
	// to define how the values would be parse
	ValuesDilimiter []string

	// Does this flag requires a value
	RequiresValue bool
}
