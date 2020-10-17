package algorithms

func bounds(array []int) (int, int) {
	return 0, len(array) - 1
}

// Hoare's partitioning
func partition(array []int) int {
	lo, hi := bounds(array)

	pivot := array[(hi+lo)/2]

	i := lo - 1
	j := hi + 1

	for {
		for ok := true; ok; ok = array[i] < pivot {
			i++
		}

		for ok := true; ok; ok = array[j] > pivot {
			j--
		}

		if i >= j {
			break
		}

		array[i], array[j] = array[j], array[i]
	}

	return j
}

// QuickSort implmentation using Hoare partitioning method
func QuickSort(array []int) {
	if lo, hi := bounds(array); lo < hi {
		p := partition(array)
		QuickSort(array[:p])
		QuickSort(array[p+1:])
	}
}
