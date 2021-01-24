package main

import (
	"fmt"
	"lab5/trees"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	t := &trees.Tree{}

	t.Generate(10)
	fmt.Println(t.IsValid())
}
