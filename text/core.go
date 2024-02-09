package rhombitext

func Bold(txt string) string {
	return "\033[1m" + txt + "\033[0m"
}

func LightGray(txt string) string {
	return "\033[97m" + txt + "\033[0m"
}
