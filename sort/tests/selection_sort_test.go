package tests

import (
	"reflect"
	"steelWinds/algorithms/sort"
	"testing"
)

type SortTest struct {
	Name   string
	Array  []int
	Sorted []int
}

func TestSelectionSort(t *testing.T) {
	testsInt := []SortTest{
		{"Sort variant 1", []int{6, 4, 10, 2, 3, 1}, []int{1, 2, 3, 4, 6, 10}},
		{"Sort variant 2", []int{4, 1, 9, 6, 7}, []int{1, 4, 6, 7, 9}},
		{"Sort variant 3", []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tt := range testsInt {
		t.Run(tt.Name, func(t *testing.T) {
			sorted_array := sort.SelectionSort(tt.Array)

			if !reflect.DeepEqual(sorted_array, tt.Sorted) {
				t.Errorf("Sort %v failed, result: %v, want %v", tt.Name, sorted_array, tt.Sorted)
			}
		})
	}
}
