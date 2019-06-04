package game

type LevelMap struct {
	Count int `json:count`
	Bricks []Brick `json:bricks`
}
