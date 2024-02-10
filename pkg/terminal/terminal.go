package terminal

import "fmt"

func HideCursor() {
	fmt.Print("\033[?25l")
}

func ShowCursor() {
	fmt.Print("\033[?25h")
}


func ClearConsole() {
	fmt.Print("\033[2J\033[H")
}

func ResetCursor() {
	fmt.Print("\033[H")
}
