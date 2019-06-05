package core

type LevelMap struct {
	Count   int      `json:count`
	Colors  []string `json:colors`
	Bricks  []Brick  `json:bricks`
	Seconds int      `json:seconds`
}
