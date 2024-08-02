package parsing

import "github.com/racg0092/rhombifer/pkg/errs"

// Parsing Errors
var (
	// Error happens when [FlagsLookup] recieves an empty or nil command flags list
	ErrFlagsNilOrEmpty = errs.NewError("flags are either nil or empty.")

	// When a short formst flag has an - and no follow up flag short format sample (- )
	ErrShortFormatHasNoFlagId = errs.NewError("shorthand '-' has not follow up flag identifier.")

	// When a flag that only accpets one value gets passed more than one value
	ErrFlagOnlyAccepstOneValue = errs.NewError("flag only accepts one value.")

	ErrMFIWithValues = errs.NewError("multi flag identifier with values. multiflag identifier ex(-abc) does not allow for values")

	ErrFlagsExtractionIsNil = errs.NewError("flag extraction is trying to access unitialized parsed flag")

	ErrUnrecognizedFlag = errs.NewError("unrecognized flag.")
)
