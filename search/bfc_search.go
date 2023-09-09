package search

import (
	"github.com/gammazero/deque"
	"golang.org/x/exp/constraints"
)

type IGenericGraph interface {
	constraints.Integer | constraints.Float | ~string
}

func push_items[T IGenericGraph](deque *deque.Deque[T], items []T) {
	for _, item := range items {
		deque.PushBack(item)
	}
}

func BFCSearch[T IGenericGraph](
	start T,
	graph map[T][]T,
	condition func(v T) bool,
) bool {
	search_stack := deque.New[T]()
	searched_els := make(map[T]bool, 0)

	push_items(search_stack, graph[start])

	for search_stack.Len() > 0 {
		current_item := search_stack.PopFront()

		if _, ok := searched_els[current_item]; ok {
			continue
		}

		if condition(current_item) {
			return true
		}

		push_items(search_stack, graph[current_item])

		searched_els[current_item] = true
	}

	return false
}
