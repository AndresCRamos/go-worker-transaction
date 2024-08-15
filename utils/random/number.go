package random

import "math/rand"

func RandInt(max, min int) int {
	delta := max - min + 1

	return rand.Intn(delta) + min
}
