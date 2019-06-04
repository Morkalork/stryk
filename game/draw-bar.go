package game

import (
	"math"
	"syscall/js"
	"time"
)

func DrawBar(canvas js.Value, maxValue int, currentValue int, label string, order int) {
	width := canvas.Get("width").Float()
	left := width * 0.1
	top := int(math.Floor(canvas.Get("height").Float() * 0.7))
	fullHeight := math.Floor(canvas.Get("height").Float() * 0.3)
	segmentHeight := int(math.Floor(fullHeight / 6))

	context := canvas.Call("getContext", "2d")

	context.Set("font", "24px serif")
	context.Call("fillStyle", "#ffffff")
	context.Call("fillText", label, left, top + (segmentHeight * order))
}

func DrawTimeBar(canvas js.Value) {
	for i:=30; i > 0; i-- {
		DrawBar(canvas, 30, i, "Time", 1)
		time.Sleep(time.Second)
	}
}
