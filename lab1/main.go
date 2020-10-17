package main

import (
	"lab1/algorithms"
)

func main() {
	alg := algorithms.Algorithm{}
	alg.Init(10)
	alg.Run(algorithms.MERGE)
}
