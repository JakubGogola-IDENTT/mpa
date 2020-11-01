package permutations

import (
	"crypto/rand"
	"log"
	"math/big"
)

func nextInt(n int) int {
	next, err := rand.Int(rand.Reader, big.NewInt(int64(n)))

	if err != nil {
		log.Panic(err)
	}

	return int(next.Int64())
}
