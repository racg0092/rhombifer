package parsing

import (
	"strings"
)

// Takes in the raw user input args and extracts the flags and their values
func ExtractFlagsFromArgs(args ...string) (map[string][]string, error) {
	flags := make(map[string][]string, 0)
	var currentflag string
	for _, i := range args {
		if strings.HasPrefix(i, "--") {
			a := i[2:]
			currentflag = "--" + a
			if _, exists := flags[currentflag]; !exists {
				flags[currentflag] = make([]string, 0)
			}
			continue
		} else if strings.HasPrefix(i, "-") {
			a := i[1:]
			if len(a) == 1 {
				currentflag = "-" + a
				if _, exists := flags[currentflag]; !exists {
					flags[currentflag] = make([]string, 0)
				}
			} else {
				for _, c := range a {
					currentflag = ""
					if _, exists := flags["-"+string(c)]; !exists {
						flags["-"+string(c)] = make([]string, 0)
					}
				}
			}
			continue

		}

		if "" == currentflag {
			return make(map[string][]string), ErrMFIWithValues
		}

		f, exists := flags[currentflag]
		if !exists {
			return make(map[string][]string), ErrFlagsExtractionIsNil
		}

		f = append(f, i)
		flags[currentflag] = f
	}
	return flags, nil
}
