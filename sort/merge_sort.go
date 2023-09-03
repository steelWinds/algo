package sort

import (
	"golang.org/x/exp/constraints"
)

func merge_arrs[T constraints.Integer | constraints.Float](leftArr, rightArr []T) []T {
	lenLeftArr := len(leftArr)
	lenRightArr := len(rightArr)

	resultArr := make([]T, 0, lenLeftArr+lenRightArr)

	left, right := 0, 0

	for left < lenLeftArr && right < lenRightArr {
		if leftArr[left] < rightArr[right] {
			resultArr = append(resultArr, leftArr[left])

			left++
		} else {
			resultArr = append(resultArr, rightArr[right])

			right++
		}
	}

	resultArr = append(resultArr, leftArr[left:]...)
	resultArr = append(resultArr, rightArr[right:]...)

	return resultArr
}

func MergeSort[T constraints.Integer | constraints.Float](arr []T) []T {
	middle := int(len(arr) / 2)

	leftArr, rightArr := arr[:middle], arr[middle:]

	if len(leftArr) > 1 {
		leftArr = MergeSort(leftArr)
	}

	if len(rightArr) > 1 {
		rightArr = MergeSort(rightArr)
	}

	return merge_arrs(leftArr, rightArr)
}
