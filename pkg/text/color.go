package text

import (
	"fmt"
	"strconv"
)

type RGB struct {
	RED   int
	GREEN int
	BLUE  int
}

func HEXToRGB(hex string) (RGB, error) {
	rgb := RGB{}
	if len(hex) < 6 {
		return rgb, fmt.Errorf("Invalid HEX color")
	}
	r, err := strconv.ParseInt(hex[:2], 16, 32)
	if err != nil {
		return rgb, err
	}
	g, err := strconv.ParseInt(hex[2:4], 16, 32)
	if err != nil {
		return rgb, err
	}
	b, err := strconv.ParseInt(hex[4:6], 16, 32)
	if err != nil {
		return rgb, err
	}
	rgb.RED = int(r)
	rgb.GREEN = int(g)
	rgb.BLUE = int(b)
	return rgb, nil
}
