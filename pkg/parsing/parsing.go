package parsing

import (
	"fmt"
	"github.com/racg0092/rhombifer/pkg/models"
	"strings"
)

const (
	LongFlag = iota
	ShortFlag
)

// Looks up flags in the args provided if any is found a map of runable flags (flr) is returned.
// The function will stop looking for flags if it reaches the end of slice of args or runs into
// a sub command
func FlagsLookup(flags []models.Flag, args ...string) (flr []models.Flag, err error) {
	if flags == nil || len(flags) <= 0 {
		return nil, fmt.Errorf("Flags is either nil or empty")
	}
	flr = make([]models.Flag, 0)
	// i need to change this to a index loop instead of a iterator
	for i := 0; i < len(args); i++ {
		a := args[i]
		if strings.HasPrefix(a, "-") && !strings.HasPrefix(a, "--") {
			a = a[1:]
			fmt.Println(a)
			for _, shorthand := range a {
				s := string(shorthand)
				flag := FindOne(flags, s, ShortFlag)
				if flag.Name == "" {
					return nil, fmt.Errorf("Unrecognized flag %s", s)
				}
				LookUpFlagValues(&flag, i, args[i:]...)
				flr = append(flr, flag)
			}
		} else if strings.HasPrefix(a, "--") {
			a = a[2:]
			flag := FindOne(flags, a, LongFlag)
			if flag.Name == "" {
				return nil, fmt.Errorf("Unrecognized flag %s", a)
			}
			// check for value and extract if need be
			flr = append(flr, flag)
		} else {
			break
		}
	}
	return flr, nil
}

// Looks up a flag with the value `v` and the type `t`.
// It returns the flag if found
func FindOne(flags []models.Flag, v string, t int) models.Flag {
	var check func(f models.Flag, v string) models.Flag
	if t == LongFlag {
		check = func(f models.Flag, v string) models.Flag {
			if f.Name == v {
				return f
			}
			return models.Flag{}
		}
	} else if t == ShortFlag {
		check = func(f models.Flag, v string) models.Flag {
			if f.ShortFormat == v {
				return f
			}
			return models.Flag{}
		}
	}
	for _, f := range flags {
		found := check(f, v)
		if found.Name != "" {
			return found
		}
	}
	return models.Flag{}
}

// Forward look up of values realted to a flag
func LookUpFlagValues(flag *models.Flag, index int, args ...string) {
	flag.Name = "change to r"
}
