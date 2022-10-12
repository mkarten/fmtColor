package Terminal

import "fmt"

func ChangeFontColor(r uint8, g uint8, b uint8) {
	fmt.Sprintf("\033[38;2;%d;%d;%dm", r, g, b)
}

func ChangeBackgroundColor(r uint8, g uint8, b uint8) {
	fmt.Sprintf("\033[48;2;%d;%d;%dm", r, g, b)
}

func ResetColor() {
	fmt.Sprintf("\033[0m")
}
