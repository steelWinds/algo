package search

import (
	"math"

	"golang.org/x/exp/constraints"
)

func BinarySearch[T constraints.Integer | constraints.Float](array []T, target T) (n T, ok bool) {
	low := 0.0
	high := float64(len(array)) - 1

	targetFloat := float64(target)

	for low <= high {
		mid := int(math.Round(float64((low + high) / 2)))

		testTarget := float64(array[mid])

		if targetFloat == testTarget {
			return target, true
		} else if targetFloat > testTarget {
			low = float64(mid + 1)
		} else if targetFloat < testTarget {
			high = float64(mid - 1)
		}
	}

	return target, false
}
