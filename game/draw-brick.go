package game

import (
	"stryk/core"
	"syscall/js"
)

func DrawBrick(colors []string, brick core.Brick, context js.Value) {
	color := colors[brick.Level]
	if brick.IsBomb {
		if brick.Level == 2 {
			color = "#000000"
		} else if brick.Level == 1 {
			color = "#858585"
		}
	}
	context.Set("fillStyle", color)
	context.Call("fillRect", brick.X, brick.Y, brick.Width, brick.Height)
}
