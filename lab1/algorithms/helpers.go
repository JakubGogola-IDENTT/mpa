package algorithms

import (
	crand "crypto/rand"
	"log"
	"math/big"
	"math/rand"
)

func bounds(arr []int) (int, int) {
	return 0, len(arr) - 1
}

func getRandomInt() (n int) {
	result, err := crand.Int(crand.Reader, big.NewInt(100))

	if err != nil {
		log.Panic(err)
	}

	n = int(result.Int64())

	if rand.Int()%2 == 1 {
		return -n
	}

	return n
}

func generateArray(size int) (arr []int) {
	arr = make([]int, size)

	for i := 0; i < size; i++ {
		arr[i] = getRandomInt()
	}

	return arr
}
