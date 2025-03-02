// Packag text provides some utilities functions for formating text
//
// You can add color boldness etc..
package text

import "fmt"

const (
	BLACK   = "30"
	RED     = "31"
	GREEN   = "32"
	YELLOW  = "33"
	BLUE    = "34"
	MAGENTA = "35"
	CYAN    = "36"
	WHITE   = "37"
)

// Sets text to bold
func Bold(txt string) string {
	return "\033[1m" + txt + "\033[0m"
}

// Sets text to LightGray
func LightGray(txt string) string {
	return "\033[97m" + txt + "\033[0m"
}

// Sets text to red
func Red(text string) string {
	return "\033[" + RED + "m" + text + "\033[0m"
}

// Sets text to green
func Green(text string) string {
	return "\033[" + GREEN + "m" + text + "\033[0m"
}

// Sets text to blue
func Blue(text string) string {
	return "\033[" + BLUE + "m" + text + "\033[0m"
}

func SetForGroundColor(color, txt string) string {
	return "\033[" + color + "m" + txt + "\033[0m"
}

func ApplyHexColor(t, hex string) (string, error) {
	rgb, err := HEXToRGB(hex)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", rgb.RED, rgb.GREEN, rgb.BLUE, t), nil
}

func ApplyRGB(rgb RGB, s string) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", rgb.RED, rgb.GREEN, rgb.BLUE, s)
}
