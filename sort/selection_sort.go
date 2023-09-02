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
	i := 0

	for i < len(array)-1 {
		smallest_idx := find_smallest_idx(array, i)

		array[i], array[smallest_idx] = array[smallest_idx], array[i]

		i++
	}

	return array
}
