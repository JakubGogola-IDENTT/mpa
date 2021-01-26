package main

import (
	anal "lab5/analytics"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	a := &anal.Analytics{}

	a.Init()
	a.TestForDifferentN()
	a.TestUniformDistribution()
	a.TestProtecionNumber()
}
