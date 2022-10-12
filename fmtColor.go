package fmtColor

import "image/color"

func ToRGB(c color.Color) (uint8, uint8, uint8) {
	r, g, b, _ := c.RGBA()
	return uint8(r), uint8(g), uint8(b)
}
