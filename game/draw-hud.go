package game

import (
	"math"
	"stryk/core"
	"syscall/js"
)

func DrawHUD(canvas js.Value, levelMap core.LevelMap) {
	width := canvas.Get("width").Int()
	top := math.Floor(canvas.Get("height").Float() * 0.7)
	height := math.Floor(canvas.Get("height").Float() * 0.3)

	context := canvas.Call("getContext", "2d")
	context.Set("fillStyle", "#000000")
	context.Call("fillRect", 0, top, width, height)

	go DrawTimeBar(canvas)
}
