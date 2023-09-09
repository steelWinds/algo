package tests

import (
	"steelWinds/algo/search"
	"testing"
)

func TestBFCSearch(t *testing.T) {
	test_graph := make(map[string][]string)

	test_graph["you"] = []string{"alice", "bob", "claire"}
	test_graph["bob"] = []string{"anuj", "peggy"}
	test_graph["alice"] = []string{"peggy"}
	test_graph["claire"] = []string{"thom", "jonny"}
	test_graph["anuj"] = []string{}
	test_graph["peggy"] = []string{}
	test_graph["thom"] = []string{}
	test_graph["jonny"] = []string{}

	test_variants_true := make(map[string]bool)
	test_variants_false := make(map[string]bool)

	// True variants
	test_variants_true["bob"] = true
	test_variants_true["peggy"] = true
	test_variants_true["jonny"] = true
	test_variants_true["thom"] = true

	// False variants
	test_variants_false["kirill"] = false
	test_variants_false["jonh"] = false
	test_variants_false["kira"] = false
	test_variants_false["bobby"] = false

	t.Run("True values", func(t *testing.T) {
		for key, val := range test_variants_true {
			ok := search.BFCSearch("you", test_graph, func(item string) bool {
				return item == key
			})

			if ok != val {
				t.Errorf("Search failed, element %v not found", key)
			}
		}
	})

	t.Run("False values", func(t *testing.T) {
		for key, val := range test_variants_false {
			ok := search.BFCSearch("you", test_graph, func(item string) bool {
				return item == key
			})

			if ok != val {
				t.Errorf("Search failed, element %v not found", key)
			}
		}
	})
}
