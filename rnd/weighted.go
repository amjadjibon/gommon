package rnd

func WeightedRandom[T any](elements []T, weights []int) T {
	total := 0
	index := 0
	for k, v := range weights {
		total += v
		if RandomIntN(total) < v {
			index = k
		}
	}
	return elements[index]
}
