// Package parsing inmplements utilities functions for manipulating and parsing the raw
// user input.
//
// This is the package used to parse flags and load values into the command execution
package parsing

import (
	"fmt"
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
	for i := 0; i < len(args); i++ {
		a := args[i]
		if strings.HasPrefix(a, "-") && !strings.HasPrefix(a, "--") {
			if len(a) == 1 {
				return nil, ErrShortFormatHasNoFlagId
			}
			idx, err := parseShortHand(a[1:], flags, args[i+1:], &foundFlags, i)
			if err != nil {
				return nil, err
			}
			i = idx
		} else if strings.HasPrefix(a, "--") {
			a = a[2:]
			idx, err := parseLongHand(a, flags, args[i+1:], &foundFlags, i)
			if err != nil {
				return nil, err
			}
			i = idx
		} else {
			break
		}
	}
	return foundFlags, nil
}

// Handles parsing for a shothand flags. It takes in `shortHands` which the flag it can be **-a** or **-abc**.
// It takes in the `flags` registered in the command if any, the `args` which is the raw input string after the flag
// has been parsed out. A pointer to `foundFlags` which will create a slice of pointers to the `flags` only
// for the found ones. The `index` represenst the position in the `args` input it used yo move the pointer in the loop to avoid parsing values as
// flags if they have already been parsed and assign to a flag
//
// This function should be called by [FlagsLookup] only. However below is simple example of usage
//
//	flags := cmd.Flags
//	args := "-rgb test"
//	shortHands := args[0]
//	var foundFlags []*models.Flag
//	var i int // this would usually be the index from a loop
//	// before running parseShortHand you would need to identifie if the flag is in short format
//	idx, err := parseShortHand(shortHands[1:], flags, args[1:], &foundFlags, i)
//	if err != nil {
//	  return err
//	}
//	i = idx // updates position of index for the loop iteration
func parseShortHand(
	shortHands string,
	flags []*models.Flag,
	args []string,
	foundFlags *[]*models.Flag,
	index int) (int, error) {
	for _, shorthand := range shortHands {
		s := string(shorthand)
		flag := FindOne(flags, s, ShortFlag)
		if flag == nil {
			return -1, fmt.Errorf("Unrecognized flag %s", s)
		}
		if len(shortHands) == 1 {
			idx, err := LookUpFlagValues(flag, index, args...)
			if err != nil {
				return idx, err
			}
			index = idx
		}
		*foundFlags = append(*foundFlags, flag)
	}
	return index, nil
}

// Hanldes parsing a longFormat flag
func parseLongHand(
	longFormat string,
	flags []*models.Flag,
	args []string,
	foundFlags *[]*models.Flag,
	index int,
) (int, error) {
	flag := FindOne(flags, longFormat, LongFlag)
	if flag == nil {
		return -1, fmt.Errorf("Unrecognized flag %s", longFormat)
	}
	index, err := LookUpFlagValues(flag, index, args...)
	if err != nil {
		return index, err
	}
	*foundFlags = append(*foundFlags, flag)
	return index, nil
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

// todo: needs testing
// Forward look up of values realted to a flag
func LookUpFlagValues(flag *models.Flag, index int, args ...string) (int, error) {
	if flag.SingleValue {
		if len(args) > 1 {
			if !strings.HasPrefix(args[1], "--") || !strings.HasPrefix(args[1], "-") {
				ErrFlagOnlyAccepstOneValue.AppendMessage("flag name: " + flag.Name)
				return index, ErrFlagOnlyAccepstOneValue
			}
		}
		index = index + 1
		flag.Values = append(flag.Values, args[0])
		return index, nil
	}
	for _, val := range args {
		if strings.HasPrefix(val, "-") || strings.HasPrefix(val, "--") {
			break
		}
		flag.Values = append(flag.Values, val)
		index = index + 1
	}

	return index, nil
}
