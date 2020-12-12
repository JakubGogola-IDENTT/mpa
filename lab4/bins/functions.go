package bins

import (
	"crypto/rand"
	"log"
	"math/big"
)

func nextInt(bound int) int {
	value, err := rand.Int(rand.Reader, big.NewInt(int64(bound)))

	if err != nil {
		log.Fatal(err)
	}

	return int(value.Int64())
}
