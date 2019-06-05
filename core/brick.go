package core

type Brick struct {
	X int `json:x`
	Y int `json:y`
	Width int `json:width`
	Height int `json:height`
	Level int `json:level`
	IsBomb bool `json:isBomb`
}
