package main

import (
	"math/rand"
)

func main() {
	var seed = 1234456789
	rand.NewSource(int64(seed))
}
