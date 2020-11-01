package main

import (
	"fmt"
	"lab2/permutations"
)

func main() {
	p := permutations.Permutation{}

	for i := 100; i <= 10000; i += 100 {
		p.SetSize(i)

		fmt.Println(i)

		for j := 0; j < 500; j++ {
			p.Permute()
			p.Properties()
		}
	}
}
