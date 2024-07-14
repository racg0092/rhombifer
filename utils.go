package rhombifer

import "strings"

//todo: If custom flag definition is implemented this function will need to be change

// Identifies if the argument being passed in is a flag. Returns true if it is and false if it isn't
func IsFirstArgFlag(arg string) bool {
	if strings.HasPrefix(arg, "-") || strings.HasPrefix(arg, "--") {
		return true
	}
	return false
}
