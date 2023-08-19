package rnd

import (
	"math/rand"
	"time"
)

var (
	source     = rand.NewSource(time.Now().UnixNano())
	randSource = rand.New(source)
)

// RandomIntN returns a rnd integer between [0, n)
func RandomIntN(n int) int {
	return randSource.Intn(n)
}

// RandomIntNGlobal returns a rnd integer between [0, n)
func RandomIntNGlobal(n int) int {
	return rand.Intn(n)
}

// RandomInt returns a rnd integer between min and max
func RandomInt(minimum, maximum int) int {
	return minimum + RandomIntN(maximum-minimum)
}
