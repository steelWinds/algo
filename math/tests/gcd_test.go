package tests

import (
	"steelWinds/algorithms/math"
	"testing"
)

type SearchTest struct {
	Name    string
	Numbers []int
	Want    float64
}

func TestBinarySearch(t *testing.T) {
	testsCases := []SearchTest{
		{"Gcd test 1 with 24 with 9", []int{24, 9}, 3},
		{"Gcd test 2 with 45 with 11", []int{45, 11}, 1},
		{"Gcd test 3 with 102 with 54", []int{102, 54}, 6},
	}

	for _, tt := range testsCases {
		t.Run(tt.Name, func(t *testing.T) {
			gcd := math.Gcd(tt.Numbers[0], tt.Numbers[1])

			if gcd != tt.Want {
				t.Error("Test failed")
			}
		})
	}
}
