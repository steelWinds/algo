package tests

import (
	sortChecks "sort"
	"steelWinds/algo/sort"
	"testing"
)

func TestQuickSort(t *testing.T) {
	for _, tt := range SortTestVariants {
		t.Run(tt.Name, func(t *testing.T) {
			sorted_array := sort.QuickSort(tt.Array, 0, len(tt.Array)-1)

			if !sortChecks.IntsAreSorted(sorted_array) {
				t.Errorf("Sort %v failed", tt.Name)
			}
		})
	}
}
