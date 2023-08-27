package sort

import (
	"golang.org/x/exp/constraints"
)

func find_smallest_idx[T constraints.Integer | constraints.Float](array []T, start int) int {
	smallest_idx := start

	for i := start; i <= len(array)-1; i++ {
		if array[i] < array[smallest_idx] {
			smallest_idx = i
		}
	}

	return smallest_idx
}

func SelectionSort[T constraints.Integer | constraints.Float](array []T) []T {
	copy_array := make([]T, len(array))

	copy(copy_array, array)

	i := 0

	for i < len(copy_array)-1 {
		smallest_idx := find_smallest_idx(copy_array, i)

		copy_array[i], copy_array[smallest_idx] = copy_array[smallest_idx], copy_array[i]

		i++
	}

	return copy_array
}
