package rhombifer

import (
	"fmt"
	"strings"

	"github.com/racg0092/rhombifer/pkg/models"
)

//todo: If custom flag definition is implemented this function will need to be change

// Identifies if the argument being passed in is a flag. Returns true if it is and false if it isn't
func IsFirstArgFlag(arg string) bool {
	if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
		return true
	}
	return false
}

// Extract expecified quantity of values from flag
func ExtractFlagValues(flag *models.Flag, quantity int) ([]string, error) {
	vals := make([]string, 0)
	if flag == nil {
		return vals, ErroFlagUndefined
	}

	if len(flag.Values) == 0 {
		ErroFlagHasNoValues.AppendMessage("flag name: " + flag.Name)
		return vals, ErroFlagHasNoValues
	}

	if quantity == 0 {
		return vals, fmt.Errorf("ExtractFlagValues func must have a quatity parameter equal or greater than 1")
	}

	for count := 0; count < quantity; count++ {
		vals = append(vals, flag.Values[count])
	}

	return vals, nil
}

// Get found flags for the current executed command
func GetFlags() (*[]*models.Flag, error) {
	if ff == nil {
		return nil, ErroFoundFlagsIsNil
	}
	return ff, nil
}

// Check found flags in the current executed command and returns all flags specified in the aliases
func FindFlags(aliases ...string) ([]*models.Flag, error) {
	var flags []*models.Flag
	if ff == nil {
		return flags, ErroFoundFlagsIsNil
	}

	if len(aliases) == 0 {
		return flags, fmt.Errorf("no aliases provided")
	}

	flags = make([]*models.Flag, 0)

	for _, alias := range aliases {
		if strings.HasPrefix(alias, "--") {
			alias = alias[2:]
		}
		if strings.HasPrefix(alias, "-") {
			alias = alias[1:]
		}
		for _, f := range *ff {
			if f.Name == alias || f.ShortFormat == alias {
				flags = append(flags, f)
			}
		}
	}
	return flags, nil
}

// Checks found flags and returns the first flag that matches any of the aliases provided. Once a flag has been
// matched the rest of the aliases are no searched, use [FindFlags] if that is your intent
func FindFlag(aliases ...string) (*models.Flag, error) {
	if ff == nil {
		return nil, ErroFoundFlagsIsNil
	}
	var flag *models.Flag
floop:
	for _, f := range *ff {
		for _, alias := range aliases {
			if strings.HasPrefix(alias, "--") {
				alias = alias[2:]
			}
			if strings.HasPrefix(alias, "-") {
				alias = alias[1:]
			}
			if alias == f.Name || alias == f.ShortFormat {
				flag = f
				break floop
			}
		}
	}
	if flag == nil {
		return nil, ErroFlagNotFound
	}
	return flag, nil
}
