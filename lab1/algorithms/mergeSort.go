package algorithms

func (a *Algorithm) merge(arr, left, right []int) {
	var i, j, k int

	for i < len(left) && j < len(right) {
		a.Comps++
		if left[i] <= right[j] {
			arr[k] = left[i]
			i++
		} else {
			arr[k] = right[j]
			j++
		}

		k++
	}

	for i < len(left) {
		arr[k] = left[i]
		i++
		k++
	}

	for j < len(right) {
		arr[k] = right[j]
		j++
		k++
	}
}

// MergeSort implementation
func (a *Algorithm) MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	lo, hi := bounds(arr)
	midd := (lo+hi)/2 + 1

	left := make([]int, len(arr[:midd]))
	right := make([]int, len(arr[midd:]))

	copy(left, arr[:midd])
	copy(right, arr[midd:])

	a.MergeSort(left)
	a.MergeSort(right)
	a.merge(arr, left, right)
}
