package parsing

import (
	"fmt"
	"github.com/racg0092/rhombifer/pkg/models"
	"strings"
)

// Looks up flags in the args provided if any is found a map of runable flags (flr) is returned.
// The function will stop looking for flags if it reaches the end of slice of args or runs into
// a sub command
func FlagsLookup(flags map[string]models.Flag, args ...string) (flr map[string]models.Flag, err error) {
	if flags == nil || len(flags) <= 0 {
		return nil, fmt.Errorf("Flags is either nil or empty")
	}
	for _, a := range args {
		if strings.HasPrefix(a, "-") && !strings.HasPrefix(a, "--") {
			a = a[1:]
			for _, shorthand := range a {
				s := string(shorthand)
				_, found := flags[s]
				if !found {
					return nil, fmt.Errorf("Unrecognized flag %s", s)
				}
			}
		} else if strings.HasPrefix(a, "--") {
			a = a[2:]
			_, found := flags[a]
			if !found {
				return nil, fmt.Errorf("Unrecognized flag %s", a)
			}
		} else {
			break
		}
	}
	return flr, nil
}
