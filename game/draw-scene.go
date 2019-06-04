package game

import (
	"math"
	"stryk/core"
	"syscall/js"
)

func getLevelFromMap(levelMap core.LevelMap, x int, y int) int {
	for _, brick := range levelMap.Bricks {
		if brick.X == x && brick.Y== y {
			return brick.Level
		}
	}

	return 0
}

func DrawScene(canvas js.Value, levelMap core.LevelMap) {
	width := canvas.Get("width").Int()
	cellWidth := int(math.Round(float64(width) / float64(levelMap.Count)))
	height := math.Floor(canvas.Get("height").Float() * 0.7)
	cellHeight := int(math.Round(float64(height) / float64(levelMap.Count)))

	context := canvas.Call("getContext", "2d")
	bricks := make([]core.Brick, 0)

	for w := 0; w < 100; w++ {
		for h := 0; h < 100; h++ {
			x := w * cellWidth
			y := h * cellHeight
			level := getLevelFromMap(levelMap, w, h)

			brick := core.Brick{X: x, Y: y, Width: cellWidth, Height: cellHeight, Level: level}
			DrawBrick(levelMap.Colors, brick, context)
			bricks = append(bricks, brick)
		}
	}

	ActionEngine(levelMap, bricks, context, canvas)
}
