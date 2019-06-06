package modals

import "syscall/js"

func MapCompleteModal (canvas js.Value) {
	windowWidth := js.Global().Get("window").Get("innerWidth").Float()
	windowHeight := js.Global().Get("window").Get("innerHeight").Float()
	left := windowWidth * 0.2
	top := windowHeight * 0.2
	width := windowWidth * 0.6
	height := windowHeight * 0.6

	canvas.Call("")
}
