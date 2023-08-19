package rnd

import (
	"testing"
)

func TestRandomInt(t *testing.T) {
	minimum := 1
	maximum := 10
	for i := 0; i < 10; i++ {
		result := RandomInt(minimum, maximum)
		if result < minimum || result > maximum {
			t.Errorf("RandomInt(%d, %d) = %d, want between %d and %d", minimum, maximum, result, minimum, maximum)
		}
	}
}

func BenchmarkRandomInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomInt(1, 10)
	}
}

func BenchmarkRandomString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomString(10, LowerAlpha)
	}
}

func BenchmarkRandomIntGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomIntGlobal(1, 10)
	}
}

func BenchmarkRandomIntNGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandomIntNGlobal(10)
	}
}
