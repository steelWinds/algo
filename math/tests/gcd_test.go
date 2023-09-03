package tests

import (
	"steelWinds/algo/math"
	"testing"

	testMath "github.com/petr-korobeinikov/golang-example/math"
)

type GcdTest struct {
	Name     string
	Operands [2]int
	Want     float64
}

func TestBinarySearch(t *testing.T) {
	testsCases := []GcdTest{
		{"Gcd test 1 with 24 with 9", [2]int{24, 9}, 3},
		{"Gcd test 2 with 45 with 11", [2]int{45, 11}, 1},
		{"Gcd test 3 with 102 with 54", [2]int{102, 54}, 6},
	}

	for _, tt := range testsCases {
		t.Run(tt.Name, func(t *testing.T) {
			gcd := math.Gcd(tt.Operands[0], tt.Operands[1])

			if gcd != tt.Want {
				t.Errorf("Gcd test failed, got %v, want %v", gcd, testMath.GCDRemainder(tt.Operands[0], tt.Operands[1]))
			}
		})
	}
}
