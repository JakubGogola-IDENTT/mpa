package main

import (
	"fmt"
	"lab2/permutations"
)

func main() {
	p := permutations.Permutation{}

	p.SetSize(10)
	p.Permute()

	fmt.Println(p.Perm)
	fmt.Println(p.Cycles())
}
