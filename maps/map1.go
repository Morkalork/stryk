package maps

import (
	"stryk/core"
	"stryk/utils"
)

func Map1() core.LevelMap {
	var bricks = make([]core.Brick, 0)
	count := 20
	padding := 6
	for col := padding; col < count-padding; col++ {
		for row := padding; row < count-padding; row++ {

			brick := core.Brick{
				Level: utils.GetRandomValue(0, 4),
				X:     col,
				Y:     row,
			}

			bricks = append(bricks, brick)
		}
	}

	levelMap := core.LevelMap{
		Colors: []string{"#E4FFE1", "#C7FFC2", "#AEFFA7", "#96FF8D"},
		Count:  count,
		Bricks: bricks,
		Seconds: 30,
	}
	return levelMap
}
