package tests

import (
	sortChecks "sort"
	"steelWinds/algorithms/sort"
	"testing"
)

func TestMergeSort(t *testing.T) {
	for _, tt := range SortTestVariants {
		t.Run(tt.Name, func(t *testing.T) {
			sorted_array := sort.MergeSort(tt.Array)

			if !sortChecks.IntsAreSorted(sorted_array) {
				t.Errorf("Sort %v failed", tt.Name)
			}
		})
	}
}
