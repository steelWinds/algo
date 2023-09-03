package sort

import (
	"golang.org/x/exp/constraints"
)

func BubbleSort[T constraints.Integer | constraints.Float](arr []T) []T {
	maxStep := len(arr)

	for range arr {
		for currentNIdx, currentN := range arr {
			nextIdx := currentNIdx + 1

			if nextIdx >= maxStep {
				break
			}

			if currentN > arr[nextIdx] {
				arr[currentNIdx], arr[nextIdx] = arr[nextIdx], arr[currentNIdx]
			}
		}

		maxStep--
	}

	return arr
}
