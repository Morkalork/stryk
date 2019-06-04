package utils

import (
	"math/rand"
	"time"
)

func GetRandomValue(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}
