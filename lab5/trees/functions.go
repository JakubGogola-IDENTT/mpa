package trees

import (
	"math/rand"
)

func randomArray(n int) (array []int) {
	array = make([]int, 2*n)

	val := 1

	for i := range array {
		array[i] = val
		val *= -1
	}

	rand.Shuffle(
		len(array),
		func(i, j int) {
			array[i], array[j] = array[j], array[i]
		},
	)

	return array
}
