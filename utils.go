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
