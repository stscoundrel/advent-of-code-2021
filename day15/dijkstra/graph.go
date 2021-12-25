package dijkstra

type edge struct {
	node   string
	weight int
}

type graph struct {
	nodes map[string][]edge
}

func NewGraph() *graph {
	return &graph{nodes: make(map[string][]edge)}
}

func (g *graph) AddEdge(origin, destiny string, weight int) {
	g.nodes[origin] = append(g.nodes[origin], edge{node: destiny, weight: weight})
}

func (g *graph) getEdges(node string) []edge {
	return g.nodes[node]
}

func (g *graph) GetPath(origin, destiny string) (int, []string) {
	h := newHeap()
	h.push(path{value: 0, nodes: []string{origin}})
	visited := make(map[string]bool)

	for len(*h.values) > 0 {
		p := h.pop()
		node := p.nodes[len(p.nodes)-1]

		if visited[node] {
			continue
		}

		if node == destiny {
			return p.value, p.nodes
		}

		for _, e := range g.getEdges(node) {

			if !visited[e.node] {
				h.push(path{value: p.value + e.weight, nodes: append([]string{}, append(p.nodes, e.node)...)})
			}
		}

		visited[node] = true
	}

	return 0, nil
}
