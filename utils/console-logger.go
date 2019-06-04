package utils

import "syscall/js"

func ConsoleLogger(text string) {
	js.Global().Get("console").Call("log", text)
}