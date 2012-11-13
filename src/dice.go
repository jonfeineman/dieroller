package dieroller

import "math/rand"

func d6() int {
	return rand.Intn(6) + 1
}

func d6e() int {
	var roll int = d6()
	if roll == 6{
		return roll + d6e()
	}
	return roll
}