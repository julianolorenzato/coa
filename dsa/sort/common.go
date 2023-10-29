package sort

import "math/rand"

func randomSlice(size int) []int {
	s := make([]int, size)

	for i := range s {
		s[i] = rand.Intn(100)
	}

	return s
}

func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
