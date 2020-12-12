package main

import (
	"fmt"
	"lab4/bins"
)

func main() {
	b := bins.Bins{}

	b.SetBallsAndBins(200, 100)

	b.Assign()

	fmt.Println(b.GetBins())
}
