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
			isBomb := false
			level := utils.GetRandomValue(0, 4)
			if utils.GetRandomValue(0, 100) >= 98 {
				isBomb = true
				level = 2
			}
			brick := core.Brick{
				Level:  level,
				X:      col,
				Y:      row,
				IsBomb: isBomb,
			}

			bricks = append(bricks, brick)
		}
	}

	levelMap := core.LevelMap{
		Id: 1,
		Colors:  []string{"#F3FFF2","#E4FFE1", "#C7FFC2", "#AEFFA7", "#96FF8D"},
		Count:   count,
		Bricks:  bricks,
		Seconds: 30,
	}
	return levelMap
}
