package rnd

import (
	"math/rand"
	"time"
)

var (
	source     = rand.NewSource(time.Now().UnixNano())
	randSource = rand.New(source)
)

const (
	LowerAlpha = "abcdefghijklmnopqrstuvwxyz"
	UpperAlpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers    = "0123456789"
)

// RandomInt returns a rnd integer between min and max
func RandomInt(minimum, maximum int) int {
	return minimum + randSource.Intn(maximum-minimum)
}

func RandomIntGlobalSeed(minimum, maximum int) int {
	return minimum + rand.Intn(maximum-minimum)
}

// RandomString returns a rnd string of length from chars
func RandomString(length int, chars string) string {
	var result string
	for i := 0; i < length; i++ {
		result += string(chars[RandomInt(0, len(chars))])
	}
	return result
}

// RandomLowerAlpha returns a rnd string of length from LowerAlpha
func RandomLowerAlpha(length int) string {
	return RandomString(length, LowerAlpha)
}

// RandomUpperAlpha returns a rnd string of length from UpperAlpha
func RandomUpperAlpha(length int) string {
	return RandomString(length, UpperAlpha)
}

// RandomAlpha returns a rnd string of length from LowerAlpha and UpperAlpha
func RandomAlpha(length int) string {
	return RandomString(length, LowerAlpha+UpperAlpha)
}

// RandomLowerAlphaNumeric returns a rnd string of length from LowerAlpha and Numbers
func RandomLowerAlphaNumeric(length int) string {
	return RandomString(length, LowerAlpha+Numbers)
}

// RandomUpperAlphaNumeric returns a rnd string of length from UpperAlpha and Numbers
func RandomUpperAlphaNumeric(length int) string {
	return RandomString(length, UpperAlpha+Numbers)
}

// RandomAlphaNumeric returns a rnd string of length from LowerAlpha, UpperAlpha and Numbers
func RandomAlphaNumeric(length int) string {
	return RandomString(length, LowerAlpha+UpperAlpha+Numbers)
}
