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
	// May remove this and add a require flags to the command itself
	Required bool

	//todo: not sure if to keep it
	// to define how the values would be parse
	ValuesDilimiter []string

	// Does this flag requires a value
	RequiresValue bool

	// Defines if the flag takes one value or multiple values
	SingleValue bool

	// Flag values parsed from the current command being run
	Values []string
}
