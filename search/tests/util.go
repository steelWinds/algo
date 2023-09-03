package tests

import (
	"math/rand"
	"time"
)

type SearchTest struct {
	Name   string
	Array  []int
	Target int
}

var rn = rand.New(rand.NewSource(time.Now().Unix()))

var array = rn.Perm(10000)
var lenArray = len(array)

var SearchTestVariants = []SearchTest{
	{"Sort 100 els", array, array[rn.Intn(lenArray)]},
	{"Sort 1000 els", array, array[rn.Intn(lenArray)]},
	{"Sort 10000 els", array, array[rn.Intn(lenArray)]},
}
