package rnd

import (
	"testing"
)

func TestWeightedRandom(t *testing.T) {
	weights := []int{3, 5, 2}           // Example weights
	elements := []string{"a", "b", "c"} // Example elements

	data := make(map[string]int)

	// Generate 100 random elements
	for i := 0; i < 100; i++ {
		data[WeightedRandom(elements, weights)]++
	}

	// Print the results
	for k, v := range data {
		t.Log(k, v, float64(v)/100000)
	}
}

func BenchmarkWeightedRandom(b *testing.B) {
	weights := []int{3, 5, 2}           // Example weights
	elements := []string{"a", "b", "c"} // Example elements

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		WeightedRandom(elements, weights)
	}
}
