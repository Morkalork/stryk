package game

import (
	"stryk/core"
	"syscall/js"
)

var bricks []core.Brick

func getBrickByCoordinate(levelMap core.LevelMap, canvas js.Value, context js.Value, x int, y int, lastBrick core.Brick) core.Brick {
	for i := 0; i < len(bricks); i++ {
		brick := bricks[i]
		minX := brick.X
		maxX := brick.X + brick.Width
		minY := brick.Y
		maxY := brick.Y + brick.Height
		if (x >= minX && x < maxX) && (y >= minY && y < maxY) {
			if brick.Level > 0 && !(lastBrick.X == brick.X && lastBrick.Y == brick.Y) {
				bricks[i].Level = bricks[i].Level - 1

				DrawBrick(levelMap.Colors, bricks[i], context)
				totalActiveBricks := CountActiveBricks(levelMap.Bricks)
				DrawProgressBar(canvas, totalActiveBricks, totalActiveBricks-CountActiveBricks(bricks))
			}

			return bricks[i]
		}
	}

	return core.Brick{}
}

func CountActiveBricks(bricks []core.Brick) int {
	count := 0
	for i := 0; i < len(bricks); i++ {
		if bricks[i].Level > 0 {
			count++
		}
	}

	return count
}

func ActionEngine(levelMap core.LevelMap, scene []core.Brick, context js.Value, canvas js.Value) {
	bricks = scene
	var lastBrick core.Brick
	mouseMove := js.NewEventCallback(js.StopImmediatePropagation, func(e js.Value) {
		x := e.Get("clientX").Int()
		y := e.Get("clientY").Int()
		lastBrick = getBrickByCoordinate(levelMap, canvas, context, x, y, lastBrick)
	})
	mouseDown := js.NewEventCallback(js.StopImmediatePropagation, func(e js.Value) {
		canvas.Call("addEventListener", "mousemove", mouseMove)
	})
	mouseUp := js.NewEventCallback(js.StopImmediatePropagation, func(e js.Value) {
		canvas.Call("removeEventListener", "mousemove", mouseMove)
	})
	canvas.Call("addEventListener", "mousedown", mouseDown)
	canvas.Call("addEventListener", "mouseup", mouseUp)
}

func ResetProgressBar(canvas js.Value, levelMap core.LevelMap) {
	DrawProgressBar(canvas, CountActiveBricks(levelMap.Bricks), 0)
}
