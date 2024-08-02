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
	var check func(f *models.Flag, v string) *models.Flag
	if t == LongFlag {
		check = func(f *models.Flag, v string) *models.Flag {
			if f.Name == v {
				return f
			}
			return nil
		}
	} else if t == ShortFlag {
		check = func(f *models.Flag, v string) *models.Flag {
			if f.ShortFormat == v {
				return f
			}
			return nil
		}
	}
	for _, f := range flags {
		found := check(f, v)
		if found != nil {
			return found
		}
	}
	return nil
}
