package fmtColor

import (
	"fmt"
	"image/color"
)

func ToRGB(c color.Color) (uint8, uint8, uint8) {
	r, g, b, _ := c.RGBA()
	return uint8(r), uint8(g), uint8(b)
}

func PrintRGB(r uint8, g uint8, b uint8, values ...any) {
	fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, fmt.Sprint(values...))
}
func PrintColor(c color.Color, values ...any) {
	r, g, b := ToRGB(c)
	PrintRGB(r, g, b, values...)
}

func PrintlnRGB(r uint8, g uint8, b uint8, values ...any) {
	fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m\n", r, g, b, fmt.Sprint(values...))
}

func PrintlnColor(c color.Color, values ...any) {
	r, g, b := ToRGB(c)
	PrintlnRGB(r, g, b, values...)
}
