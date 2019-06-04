package game

import (
	"math"
	"syscall/js"
)

func getLevelFromMap(levelMap LevelMap, x int, y int) int {
	for _, brick := range levelMap.Bricks {
		if brick.X == x && brick.Y== y {
			return brick.Level
		}
	}

	return 0
}

func DrawScene(canvas js.Value, levelMap LevelMap) {
	width := canvas.Get("width").Int()
	cellWidth := int(math.Round(float64(width) / float64(levelMap.Count)))
	height := canvas.Get("height").Int()
	cellHeight := int(math.Round(float64(height) / float64(levelMap.Count)))

	context := canvas.Call("getContext", "2d")
	bricks := make([]Brick, 0)

	for w := 0; w < 100; w++ {
		for h := 0; h < 100; h++ {
			x := w * cellWidth
			y := h * cellHeight
			level := getLevelFromMap(levelMap, w, h)

			brick := Brick{X: x, Y: y, Width: cellWidth, Height: cellHeight, Level: level}
			DrawBrick(brick, context)
			bricks = append(bricks, brick)
		}
	}

	ActionEngine(bricks, context, canvas)
}
