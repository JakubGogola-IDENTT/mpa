package algorithms

import "fmt"

func merge(arr, left, right []int) {
	var i, j, k int

	for i < len(left) && j < len(right) {
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
func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	midd := len(arr) / 2

	fmt.Println(midd)

	left := make([]int, len(arr[:midd]))
	right := make([]int, len(arr[midd:]))

	copy(left, arr[:midd])
	copy(right, arr[midd:])

	MergeSort(left)
	MergeSort(right)
	merge(arr, left, right)
}
