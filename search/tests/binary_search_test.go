package tests

import (
	"sort"
	"steelWinds/algo/search"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	for _, tt := range SearchTestVariants {
		t.Run(tt.Name, func(t *testing.T) {
			sort.Ints(tt.Array)

			_, ok := search.BinarySearch(tt.Array, 1)

			if !ok {
				t.Errorf("Search failed, element %v not found", tt.Target)
			}
		})
	}
}
