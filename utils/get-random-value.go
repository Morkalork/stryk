package utils

import (
	"math/rand"
)

func GetRandomValue(min int, max int) int {
	return rand.Intn(max - min) + min
}
