package sort

import (
	"golang.org/x/exp/constraints"
)

func QuickSort[T constraints.Integer | constraints.Float](arr []T, start, end int) []T {
	firstIdx, lastIdx, pivotIdx := start, end, int((start+end)/2)

	middle := arr[pivotIdx]

	for firstIdx <= lastIdx {
		for arr[firstIdx] < middle {
			firstIdx++
		}
		for arr[lastIdx] > middle {
			lastIdx--
		}

		if firstIdx <= lastIdx {
			arr[firstIdx], arr[lastIdx] = arr[lastIdx], arr[firstIdx]

			firstIdx++
			lastIdx--
		}
	}

	if start < end {
		QuickSort(arr, start, lastIdx)
		QuickSort(arr, firstIdx, end)
	}

	return arr
}
