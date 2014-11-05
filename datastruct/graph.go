package datastruct

// Graph represents the graph data structure.
type Graph struct {
	v   int
	e   int
	adj [][]int
}

// NewGraph returns an adjacency-list graph representation.
func NewGraph(v int) *Graph {
	g := &Graph{v: v}
	g.adj = make([][]int, v)
	return g
}

// Edge adds a new edge to the graph.
func (g *Graph) Edge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.e++
}

// Adj returns the adjacent vertices of the vertex v.
func (g *Graph) Adj(v int) []int {
	return g.adj[v]
}

// V returns the number of vertices.
func (g *Graph) V() int {
	return g.v
}

// E returns the number of edges.
func (g *Graph) E() int {
	return g.e
}
