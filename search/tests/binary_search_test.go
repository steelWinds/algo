package tests

import (
	"steelWinds/algorithms/search"
	"testing"
)

type SearchTest[T int16 | float64] struct {
	Name   string
	Array  []T
	Target T
	Want   bool
}

func TestBinarySearch(t *testing.T) {
	testsInt := []SearchTest[int16]{
		{"4 included in array", []int16{1, 2, 3, 4, 5}, 4, true},
		{"4 not included in array", []int16{1, 2, 3, 4, 5}, 9, false},
	}

	testsFloat := []SearchTest[float64]{
		{"8 included in array", []float64{4, 5, 8, 9, 11}, 8, true},
	}

	t.Run("Int array binary search", func(t *testing.T) {
		for _, tt := range testsInt {
			t.Run(tt.Name, func(t *testing.T) {
				target, ok := search.BinarySearch(tt.Array, tt.Target)

				if target != tt.Target || ok != tt.Want {
					t.Error("Test failed")
				}
			})
		}
	})

	t.Run("Float array binary search", func(t *testing.T) {
		for _, tt := range testsFloat {
			t.Run(tt.Name, func(t *testing.T) {
				target, ok := search.BinarySearch(tt.Array, tt.Target)

				if target != tt.Target || ok != tt.Want {
					t.Error("Test failed")
				}
			})
		}
	})
}
