package game

import (
	"math"
	"stryk/core"
	"syscall/js"
)

var callbackIds []int

func DrawBar(canvas js.Value, maxValue int, currentValue int, label string, order int) {
	width := canvas.Get("width").Float()
	barWidth := int(math.Floor(width * 0.8))
	left := int(math.Floor(width * 0.1))
	top := int(math.Floor(canvas.Get("height").Float() * 0.7))
	fullHeight := math.Floor(canvas.Get("height").Float() * 0.3)
	segmentHeight := int(math.Floor(fullHeight / 6))

	context := canvas.Call("getContext", "2d")

	context.Set("font", "24px serif")
	context.Set("fillStyle", "#ffffff")
	context.Call("fillText", label, left, top+(segmentHeight*order))

	barTop := int(float64(top) + (float64(segmentHeight) * (float64(order) + 0.5)))
	lineWidth := 2
	padding := lineWidth * 2
	diff := float64(maxValue - currentValue)
	barWidthPercentage := diff / float64(maxValue)
	currentBarWidth := int(float64(barWidth-(padding*2)) * barWidthPercentage)
	context.Set("strokeStyle", "#ffffff")
	context.Set("lineWidth", lineWidth)
	context.Call("strokeRect", left, barTop, barWidth, segmentHeight)
	// Remove old bar
	context.Set("fillStyle", "#000000")
	context.Call("fillRect", left+lineWidth, barTop+padding, barWidth-padding, segmentHeight-(padding*2))
	// Draw new bar
	context.Set("fillStyle", "#B6D7F4")
	context.Call("fillRect", left+padding, barTop+padding, currentBarWidth, segmentHeight-(padding*2))
}

func DrawTimeBar(canvas js.Value, seconds int, levelMap core.LevelMap) {
	if callbackIds == nil {
		callbackIds = make([]int, 0)
	} else {
		for _, callbackId := range callbackIds {
			js.Global().Call("clearTimeout", callbackId)
		}
	}

	maxTime := seconds * 10 // Just don't.. you know. Don't.
	counter := 0
	timeoutCallback := js.NewCallback(func(args []js.Value) {
		DrawBar(canvas, maxTime, counter, "Time", 1)
		counter++
	})

	for i := maxTime; i > 0; i-- {
		callbackId := js.Global().Call("setTimeout", timeoutCallback, 50*i).Int()
		callbackIds = append(callbackIds, callbackId)
	}
}

func DrawProgressBar(canvas js.Value, maxProgress int, currentProgress int) {
	DrawBar(canvas, maxProgress, currentProgress, "Progress", 4)
}
