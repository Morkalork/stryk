package game

import "syscall/js"

func DrawBrick(brick Brick, context js.Value) {
	context.Set("fillStyle", Colors[brick.Level])
	context.Call("fillRect", brick.X, brick.Y, brick.Width, brick.Height)
}
