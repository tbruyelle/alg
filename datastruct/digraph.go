package datastruct

// Digraph represents the graph data structure.
type Digraph struct {
	// V represents the number of vertices.
	V int
	// E represents the number of edges.
	E   int
	adj [][]int
}

// NewDigraph returns an adjacency-list graph representation.
func NewDigraph(V int) *Digraph {
	g := &Digraph{V: V}
	g.adj = make([][]int, V)
	return g
}

// Edge adds a new edge to the graph.
func (g *Digraph) Edge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.E++
}

// Adj returns the adjacent vertices of the vertex v.
func (g *Digraph) Adj(v int) []int {
	return g.adj[v]
}
