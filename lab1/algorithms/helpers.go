package algorithms

import (
	crand "crypto/rand"
	"log"
	"math/big"
)

func bounds(arr []int) (int, int) {
	return 0, len(arr) - 1
}

func getRandomInt() int {
	result, err := crand.Int(crand.Reader, big.NewInt(200))

	if err != nil {
		log.Panic(err)
	}

	return int(result.Int64())
}

func generateArray(size int) (arr []int) {
	arr = make([]int, size)

	for i := 0; i < size; i++ {
		arr[i] = getRandomInt()
	}

	return arr
}
