package modals

import "syscall/js"

func MapCompleteModal (canvas js.Value, onNext func()) {
	windowWidth := js.Global().Get("window").Get("innerWidth").Float()
	windowHeight := js.Global().Get("window").Get("innerHeight").Float()
	left := windowWidth * 0.1
	top := windowHeight * 0.1
	width := windowWidth * 0.8
	height := windowHeight * 0.8

	context := canvas.Call("getContext", "2d")
	context.Set("fillStyle", "#333333")
	context.Set("shadowColor", "#222222")
	context.Set("shadowBlur", "10")
	context.Call("fillRect", left, top, width, height)
}
