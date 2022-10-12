package Color

import (
	"github.com/mkarten/fmtColor"
	"github.com/mkarten/fmtColor/RGB"
	"image/color"
)

func PrintColor(c color.Color, values ...any) {
	r, g, b := fmtColor.ToRGB(c)
	RGB.PrintRGB(r, g, b, values...)
}

func PrintlnColor(c color.Color, values ...any) {
	r, g, b := fmtColor.ToRGB(c)
	RGB.PrintlnRGB(r, g, b, values...)
}

func PrintfColor(c color.Color, format string, values ...any) {
	r, g, b := fmtColor.ToRGB(c)
	RGB.PrintfRGB(r, g, b, format, values...)
}

func PrintflnColor(c color.Color, format string, values ...any) {
	r, g, b := fmtColor.ToRGB(c)
	RGB.PrintflnRGB(r, g, b, format, values...)
}

func SprintColor(c color.Color, values ...any) string {
	r, g, b := fmtColor.ToRGB(c)
	return RGB.SprintRGB(r, g, b, values...)
}

func SprintlnColor(c color.Color, values ...any) string {
	r, g, b := fmtColor.ToRGB(c)
	return RGB.SprintlnRGB(r, g, b, values...)
}

func SprintfColor(c color.Color, format string, values ...any) string {
	r, g, b := fmtColor.ToRGB(c)
	return RGB.SprintfRGB(r, g, b, format, values...)
}

func SprintflnColor(c color.Color, format string, values ...any) string {
	r, g, b := fmtColor.ToRGB(c)
	return RGB.SprintflnRGB(r, g, b, format, values...)
}
