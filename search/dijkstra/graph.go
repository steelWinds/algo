package dijkstra

type Edge struct {
	node   string
	weight int
}

type Graph struct {
	nodes map[string][]Edge
}

func NewGraph() *Graph {
	return &Graph{nodes: make(map[string][]Edge)}
}

func (g *Graph) AddEdge(origin, destiny string, weight int) {
	g.nodes[origin] = append(g.nodes[origin], Edge{node: destiny, weight: weight})
	g.nodes[destiny] = append(g.nodes[destiny], Edge{node: origin, weight: weight})
}

func (g *Graph) GetEdges(node string) []Edge {
	return g.nodes[node]
}
