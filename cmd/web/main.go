package main

import (
	"encoding/json"
	"fmt"
	"stryk/core"
	"stryk/game"
	"stryk/modals"
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
	request.OnLoad(func(args []js.Value) {
		var levelMap core.LevelMap
		json.Unmarshal([]byte(request.GetResponseText()), &levelMap)

		mapChanger := make(chan int)
		levelMap.AddListener("Finish", mapChanger)

		canvas := getCanvas()
		bricks := game.DrawScene(canvas, levelMap)
		game.ActionEngine(levelMap, bricks, canvas)
		game.DrawHUD(canvas, levelMap)

		go func() {
			for {
				currentLevel := <-mapChanger
				modals.MapCompleteModal(canvas, func() {
					loadMap(currentLevel + 1)
				})
			}
		}()
	})
	request.Open("GET", mapName)
	request.Send()
}

func main() {
	loadMap(1)

	c := make(chan struct{})
	<-c
}
