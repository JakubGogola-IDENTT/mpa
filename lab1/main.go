package main

import (
	"fmt"
	"lab1/algorithms"
)

func main() {
	array := []int{2, 1, 5, 23, 11, 13}
	fmt.Println(array)

	algorithms.QuickSort(array)

	fmt.Println(array)
}
