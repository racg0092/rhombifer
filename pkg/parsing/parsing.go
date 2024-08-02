// Package parsing inmplements utilities functions for manipulating and parsing the raw
// user input.
//
// This is the package used to parse flags and load values into the command execution
package parsing

import (
	"strings"

	"github.com/racg0092/rhombifer/pkg/models"
)

const (
	// Long Format for flag
	LongFlag = iota
	// Short Format for flag
	ShortFlag
)

// Looks up flags in the args provided if any is found a map of runable flags (foundFlags) is returned.
// The function will stop looking for flags if it reaches the end of slice of args.
func FlagsLookup(flags []*models.Flag, args ...string) (foundFlags []*models.Flag, err error) {
	if flags == nil || len(flags) <= 0 {
		return nil, ErrFlagsNilOrEmpty
	}
	parsedFlags, err := ExtractFlagsFromArgs(args...)
	if err != nil {
		return foundFlags, err
	}
	for k, v := range parsedFlags {
		var ftype int
		var flagIndicatorLen int
		if strings.HasPrefix(k, "--") {
			ftype = LongFlag
			flagIndicatorLen = 2
		} else if strings.HasPrefix(k, "-") {
			ftype = ShortFlag
			flagIndicatorLen = 1
		}
		fname := k[flagIndicatorLen:]
		flag := FindOne(flags, fname, ftype)
		if flag == nil {
			return foundFlags, ErrUnrecognizedFlag
		}
		err := flag.AddValues(v...)
		if err != nil {
			return foundFlags, err
		}
		foundFlags = append(foundFlags, flag)
	}
	return foundFlags, nil
}

// Looks up a flag with the value `v` and the type `t`.
// It returns the flag if found
func FindOne(flags []*models.Flag, v string, t int) *models.Flag {
	for _, f := range flags {
		short, long := f.GetNames()
		if t == ShortFlag {
			if short == v {
				return f
			}
		} else if t == LongFlag {
			if long == v {
				return f
			}
		}
	}
	return nil
}
