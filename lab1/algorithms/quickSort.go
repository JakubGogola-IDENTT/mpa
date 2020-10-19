package algorithms

func (a *Algorithm) partition(arr []int) int {
	lo, hi := bounds(arr)

	pivot := arr[(hi+lo)/2]

	i := lo - 1
	j := hi + 1

	for {
		for ok := true; ok; ok = arr[i] < pivot {
			i++
			a.Comps++
		}

		for ok := true; ok; ok = arr[j] > pivot {
			j--
			a.Comps++
		}

		if i >= j {
			break
		}

		arr[i], arr[j] = arr[j], arr[i]
	}

	return j
}

// QuickSort implmentation using Hoare partitioning method
func (a *Algorithm) QuickSort(arr []int) {
	if lo, hi := bounds(arr); lo >= hi {
		return
	}

	p := a.partition(arr)
	a.QuickSort(arr[:p])
	a.QuickSort(arr[p+1:])
}
