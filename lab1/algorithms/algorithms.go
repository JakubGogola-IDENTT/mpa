package algorithms

import (
	"fmt"
	"log"
)

// Init initializes sorting algorithm with given array size
func (a *Algorithm) Init(arraySize int) {
	a.ArraySize = arraySize
}

// Run runs algorithm of given type
func (a *Algorithm) Run(algType string) {
	arr := generateArray(a.ArraySize)
	a.Comps = 0

	switch algType {
	case MERGE:
		a.MergeSort(arr)
		break
	case QUICK:
		a.QuickSort(arr)
		break
	default:
		log.Panic("Invalid algorithm type")
	}

	fmt.Println(arr)
	fmt.Println(a.Comps)
}
