package tests

import (
	"math/rand"
	"time"
)

type SortTest struct {
	Name  string
	Array []int
}

var rn = rand.New(rand.NewSource(time.Now().Unix()))

var SortTestVariants = []SortTest{
	{"Sort 100 els", rn.Perm(100)},
	{"Sort 1000 els", rn.Perm(1000)},
	{"Sort 10000 els", rn.Perm(10000)},
}
