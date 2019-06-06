package core

type LevelMap struct {
	Id              int      `json:id`
	Count           int      `json:count`
	Colors          []string `json:colors`
	Bricks          []Brick  `json:bricks`
	Seconds         int      `json:seconds`
	Finisher        chan int `json:"-"`
}
