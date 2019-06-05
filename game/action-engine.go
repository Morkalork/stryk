package game

import (
	"stryk/core"
	"syscall/js"
)

var bricks []core.Brick

func detonateBomb(brick core.Brick, levelMap core.LevelMap, canvas js.Value, x int, y int, lastBrick core.Brick) {
	width := brick.Width
	height := brick.Height
	lastBrick = updateBrickLevel(levelMap, canvas, x-width, y-height, lastBrick, false) // Top left
	lastBrick = updateBrickLevel(levelMap, canvas, x, y-height, lastBrick, false)       // Top
	lastBrick = updateBrickLevel(levelMap, canvas, x+width, y-height, lastBrick, false) // Top right

	lastBrick = updateBrickLevel(levelMap, canvas, x-width, y, lastBrick, false) // Left
	lastBrick = updateBrickLevel(levelMap, canvas, x+width, y, lastBrick, false) // Right

	lastBrick = updateBrickLevel(levelMap, canvas, x-width, y+height, lastBrick, false) // Bottom left
	lastBrick = updateBrickLevel(levelMap, canvas, x, y+1, lastBrick, false)            // Bottom
	lastBrick = updateBrickLevel(levelMap, canvas, x+width, y+height, lastBrick, false) // Bottom right
}

func updateBrickLevel(
	levelMap core.LevelMap,
	canvas js.Value,
	x int,
	y int,
	lastBrick core.Brick,
	canBomb bool,
) core.Brick {
	for i := 0; i < len(bricks); i++ {
		brick := bricks[i]
		minX := brick.X
		maxX := brick.X + brick.Width
		minY := brick.Y
		maxY := brick.Y + brick.Height
		if (x >= minX && x < maxX) && (y >= minY && y < maxY) {
			if brick.Level > 0 && !(lastBrick.X == brick.X && lastBrick.Y == brick.Y) {
				context := canvas.Call("getContext", "2d")
				if brick.IsBomb {
					if canBomb {
						detonateBomb(bricks[i], levelMap, canvas, brick.X, brick.Y, lastBrick)
					} else {
						continue // If we happen to read a bomb during a bombing raid
					}
				}

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
		lastBrick = updateBrickLevel(levelMap, canvas, x, y, lastBrick, true)
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
