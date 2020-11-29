package delaunay

import (
	crand "crypto/rand"
	"log"
	"math/big"
	"math/rand"
)

func getRandomFloat64() float64 {
	mult, err := crand.Int(crand.Reader, big.NewInt(100))

	if err != nil {
		log.Panic(err)
	}
	next := rand.Float64() * float64(mult.Int64())

	if mult.Int64()%2 == 1 {
		return -next
	}

	return next
}

func getRandomCoords() (float64, float64) {
	return getRandomFloat64(), getRandomFloat64()
}
