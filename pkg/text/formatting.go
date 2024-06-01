package rhombitext

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

func Bold(txt string) string {
	return "\033[1m" + txt + "\033[0m"
}

func LightGray(txt string) string {
	return "\033[97m" + txt + "\033[0m"
}

func Red(text string) string {
	return "\033[" + RED + "m" + text + "\033[0m"
}

func Green(text string) string {
	return "\033[" + GREEN + "m" + text + "\033[0m"
}

func Blue(text string) string {
	return "\033[" + BLUE + "m" + text + "\033[0m"
}

func SetForGroundColor(color, txt string) string {
	return "\033[" + color + "m" + txt + "\033[0m"
}
