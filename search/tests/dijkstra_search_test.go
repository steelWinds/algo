package tests

import (
	"fmt"
	"steelWinds/algo/search/dijkstra"
	"testing"
)

func TestDijkstraSearch(t *testing.T) {
	test_graph := dijkstra.NewGraph()
	test_graph.AddEdge("S", "B", 4)
	test_graph.AddEdge("S", "C", 2)
	test_graph.AddEdge("B", "C", 1)
	test_graph.AddEdge("B", "D", 5)
	test_graph.AddEdge("C", "D", 8)
	test_graph.AddEdge("C", "E", 10)
	test_graph.AddEdge("D", "E", 2)
	test_graph.AddEdge("D", "T", 6)
	test_graph.AddEdge("E", "T", 2)

	// TODO: write normal tests for this

	fmt.Println(test_graph.GetPath("S", "B"))
}
