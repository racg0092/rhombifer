package parsing

import "github.com/racg0092/rhombifer/pkg/errs"

// Parsing Errors
var (
	// Error happens when [FlagsLookup] recieves an empty or nil command flags list
	ErrFlagsNilOrEmpty = errs.NewError("Flags are either nil or empty")

	// When a short formst flag has an - and no follow up flag short format sample (- )
	ErrShortFormatHasNoFlagId = errs.NewError("Shorthand '-' has not follow up flag identifier")
)
