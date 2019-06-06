package game

import (
	"math"
	"stryk/core"
	"syscall/js"
)

func getBrick(levelMap core.LevelMap, x int, y int) core.Brick {
	for _, brick := range levelMap.Bricks {
		if brick.X == x && brick.Y == y {
			return brick
		}
	}

	return core.Brick{}
}

func DrawScene(canvas js.Value, levelMap core.LevelMap) []core.Brick {
	width := canvas.Get("width").Int()
	cellWidth := int(math.Round(float64(width) / float64(levelMap.Count)))
	height := math.Floor(canvas.Get("height").Float() * 0.7)
	cellHeight := int(math.Round(float64(height) / float64(levelMap.Count)))

	context := canvas.Call("getContext", "2d")
	context.Set("fillStyle", levelMap.Colors[0])
	context.Call("fillRect", 0, 0, width, height)

	bricks := make([]core.Brick, 0)

	for w := 0; w < levelMap.Count; w++ {
		for h := 0; h < levelMap.Count; h++ {
			x := w * cellWidth
			y := h * cellHeight
			matchingBrick := getBrick(levelMap, w, h)

			brick := core.Brick{
				X:      x,
				Y:      y,
				Width:  cellWidth,
				Height: cellHeight,

				Level:  matchingBrick.Level,
				IsBomb: matchingBrick.IsBomb,
			}

			DrawBrick(levelMap.Colors, brick, context)
			bricks = append(bricks, brick)
		}
	}

	return bricks
}
