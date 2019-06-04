package game

import (
	"syscall/js"
)

var bricks []Brick

func getBrickByCoordinate(context js.Value, x int, y int, lastBrick Brick) Brick {
	for i := 0; i < len(bricks); i++ {
		brick := bricks[i]
		minX := brick.X
		maxX := brick.X + brick.Width
		minY := brick.Y
		maxY := brick.Y + brick.Height
		if (x >= minX && x < maxX) && (y >= minY && y < maxY) {
			if brick.Level > 0 && !(lastBrick.X == brick.X && lastBrick.Y == brick.Y) {
				bricks[i].Level = bricks[i].Level - 1
				DrawBrick(bricks[i], context)
			}

			return bricks[i]
		}
	}

	return Brick{}
}

func ActionEngine(scene []Brick, context js.Value, canvas js.Value) {
	bricks = scene
	var lastBrick Brick
	handleClick := js.NewEventCallback(js.StopImmediatePropagation, func(e js.Value) {
		x := e.Get("clientX").Int()
		y := e.Get("clientY").Int()
		lastBrick = getBrickByCoordinate(context, x, y, lastBrick)
	})
	canvas.Call("addEventListener", "mousemove", handleClick)
}
