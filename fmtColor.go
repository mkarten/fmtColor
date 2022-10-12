package fmtColor

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"
)

func ToRGB(c color.Color) (uint8, uint8, uint8) {
	r, g, b, _ := c.RGBA()
	return uint8(r), uint8(g), uint8(b)
}

func RandRGB(seed ...int64) (uint8, uint8, uint8) {
	if len(seed) > 0 {
		rand.Seed(seed[0])
	} else {
		rand.Seed(time.Now().UnixNano())
	}
	r := uint8(rand.Intn(255))
	g := uint8(rand.Intn(255))
	b := uint8(rand.Intn(255))
	return r, g, b
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

func PrintRandColor(values ...any) {
	r, g, b := RandRGB()
	PrintRGB(r, g, b, values...)
}

func PrintlnRandColor(values ...any) {
	PrintRandColor(values...)
	fmt.Println()
}
