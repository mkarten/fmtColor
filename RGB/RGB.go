package RGB

import (
	"fmt"
	"math/rand"
	"time"
)

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

func PrintRandColor(values ...any) {
	r, g, b := RandRGB()
	PrintRGB(r, g, b, values...)
}

func PrintlnRandColor(values ...any) {
	PrintRandColor(values...)
	fmt.Println()
}

func PrintRGB(r uint8, g uint8, b uint8, values ...any) {
	fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, fmt.Sprint(values...))
}

func PrintlnRGB(r uint8, g uint8, b uint8, values ...any) {
	PrintRGB(r, g, b, values...)
	fmt.Println()
}

func PrintfRGB(r uint8, g uint8, b uint8, toFormat string, values ...any) {
	fmt.Printf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, fmt.Sprintf(toFormat, values...))
}

func PrintflnRGB(r uint8, g uint8, b uint8, toFormat string, values ...any) {
	PrintfRGB(r, g, b, toFormat, values...)
	fmt.Println()
}

func SprintRGB(r uint8, g uint8, b uint8, values ...any) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, fmt.Sprint(values...))
}

func SprintlnRGB(r uint8, g uint8, b uint8, values ...any) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, fmt.Sprintln(values...))
}

func SprintfRGB(r uint8, g uint8, b uint8, toFormat string, values ...any) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m", r, g, b, fmt.Sprintf(toFormat, values...))
}

func SprintflnRGB(r uint8, g uint8, b uint8, toFormat string, values ...any) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s\033[0m\n", r, g, b, fmt.Sprintf(toFormat, values...))
}
