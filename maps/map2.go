package maps

import (
	"stryk/core"
	"stryk/utils"
)

func Map2() core.LevelMap {
	var bricks = make([]core.Brick, 0)
	count := 50
	padding := 10
	for col := padding; col < count-padding; col++ {
		for row := padding; row < count-padding; row++ {

			if col%2 == 0 && row%2 == 0 {
				brick := core.Brick{
					Level: utils.GetRandomValue(0, 4),
					X:     col,
					Y:     row,
				}

				bricks = append(bricks, brick)
			}
		}
	}

	levelMap := core.LevelMap{
		Colors: []string{"#F0F8FF", "#DBEEFF", "#BFE1FF", "#9ED2FF"},
		Count:  count,
		Bricks: bricks,
	}
	return levelMap
}
