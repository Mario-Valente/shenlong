package utils

import (
	"math/rand"
	"strconv"
)

func CreateRadomName(name string) string {
	randomNumber := rand.Intn(1000)
	name = name + strconv.Itoa(randomNumber)
	return name
}
