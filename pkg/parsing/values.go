package parsing

import (
	"regexp"
	"strings"
)

// Cheks if the input string is composed of only raw values an no flags.
// If it is it will return true otherwise false. It returns an error if something fails
func IsRawValuesOnly(args ...string) (bool, error) {
	joined := strings.Join(args, " ")
	reg, err := regexp.Compile(`--`)
	if err != nil {
		return false, err
	}
	matched := reg.Match([]byte(joined))
	return !matched, nil
}
