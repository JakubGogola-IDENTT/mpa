package main

import (
	"lab5/trees"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	t := &trees.Tree{}

	t.Generate(10)
}
