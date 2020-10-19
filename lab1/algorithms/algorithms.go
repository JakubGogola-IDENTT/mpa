package algorithms

import (
	"log"
)

// Init initializes sorting algorithm with given array size
func (a *Algorithm) Init(arraySize int) {
	a.ArraySize = arraySize
}

// Run runs algorithm of given type
func (a *Algorithm) Run(algType string) {
	a.Array = generateArray(a.ArraySize)
	a.Comps = 0

	switch algType {
	case MERGE:
		a.MergeSort(a.Array)
		break
	case QUICK:
		a.QuickSort(a.Array)
		break
	default:
		log.Panic("Invalid algorithm type")
	}
}
