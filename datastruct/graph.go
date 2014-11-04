package datastruct

// Graph represents the graph data structure.
type Graph struct {
	// V represents the number of vertices.
	V int
	// E represents the number of edges.
	E   int
	adj [][]int
}

// NewGraph returns an adjacency-list graph representation.
func NewGraph(V int) *Graph {
	g := &Graph{V: V}
	g.adj = make([][]int, V)
	return g
}

// Edge adds a new edge to the graph.
func (g *Graph) Edge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.E++
}

// Adj returns the adjacent vertices of the vertex v.
func (g *Graph) Adj(v int) []int {
	return g.adj[v]
}
