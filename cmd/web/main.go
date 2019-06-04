package main

import (
	"encoding/json"
	"fmt"
	"stryk/core"
	"stryk/game"
	"stryk/xhr"
	"syscall/js"
)

func getCanvas() js.Value {
	windowWidth := js.Global().Get("window").Get("innerWidth")
	windowHeight := js.Global().Get("window").Get("innerHeight")
	canvas := js.Global().Get("document").Call("getElementById", "scene")
	canvas.Set("width", windowWidth)
	canvas.Set("height", windowHeight)

	return canvas
}

func loadMap(level int) {
	mapName := fmt.Sprintf("assets/map-%d.json", level)
	request := xhr.New()
	println("Starting!")
	request.OnLoad(func(args []js.Value) {
		println("DONE!")
		var levelMap core.LevelMap
		json.Unmarshal([]byte(request.GetResponseText()), &levelMap)
		canvas := getCanvas()
		game.DrawScene(canvas, levelMap)
		game.DrawHUD(canvas, levelMap)
	})
	request.Open("GET", mapName)
	request.Send()
}

func main() {
	loadMap(1)

	c := make(chan struct{})
	<-c
}