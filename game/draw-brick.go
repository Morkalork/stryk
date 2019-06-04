package game

import (
	"stryk/core"
	"syscall/js"
)

func DrawBrick(colors []string, brick core.Brick, context js.Value) {
	context.Set("fillStyle", colors[brick.Level])
	context.Call("fillRect", brick.X, brick.Y, brick.Width, brick.Height)
}
