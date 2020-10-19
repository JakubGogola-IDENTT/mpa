package main

import (
	"lab1/algorithms"
)

func main() {
	alg := algorithms.Algorithm{}

	for i := 100; i <= 100000; i += 100 {
		alg.Init(i)
		alg.Run(algorithms.QUICK)
	}
}
