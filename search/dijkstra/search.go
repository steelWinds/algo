package dijkstra

func (g *Graph) GetPath(origin, destiny string) (int, []string) {
	h := newHeap()

	h.push(path{value: 0, nodes: []string{origin}})

	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		// Find the nearest yet to visit node
		p := h.pop()

		node := p.nodes[len(p.nodes)-1]

		if visited[node] {
			continue
		}

		if node == destiny {
			return p.value, p.nodes
		}

		for _, e := range g.GetEdges(node) {
			if !visited[e.node] {
				// We calculate the total spent so far plus the cost and the path of getting here
				h.push(path{value: p.value + e.weight, nodes: append([]string{}, append(p.nodes, e.node)...)})
			}
		}

		visited[node] = true
	}

	return 0, nil
}
