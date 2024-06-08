package parsing

import (
	"fmt"
	"github.com/racg0092/rhombifer/pkg/models"
	"strings"
)

// Looks up flags in the args provided if any is found is bundle in a slice and return.
// The function will stop looking for flags if it reaches the end of slice of args or runs into
// a sub command
func FlagsLookup(flags map[string]models.Flag, args ...string) (map[string]models.Flag, error) {
	if flags == nil || len(flags) <= 0 {
		return flags, fmt.Errorf("Flags is either nil or empty")
	}
	foundFlags := make(map[string]models.Flag)
	for _, a := range args {
		if strings.HasPrefix(a, "-") {
			a = a[1:]
			for _, shorthand := range a {
				s := string(shorthand)
				fmt.Println(s)
			}
		} else if strings.HasPrefix(a, "--") {

		} else {
			break
		}
	}
	return foundFlags, nil
}
