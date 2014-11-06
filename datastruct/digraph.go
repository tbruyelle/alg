package datastruct

// Digraph represents the graph data structure.
type Digraph struct {
	v   int
	e   int
	adj [][]int
}

// NewDigraph returns an adjacency-list graph representation.
func NewDigraph(v int) *Digraph {
	g := &Digraph{v: v}
	g.adj = make([][]int, v)
	return g
}

// Edge adds a new edge to the graph.
func (g *Digraph) Edge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.e++
}

// Adj returns the adjacent vertices of the vertex v.
func (g *Digraph) Adj(v int) []int {
	return g.adj[v]
}

// V returns the number of vertices.
func (g *Digraph) V() int {
	return g.v
}

// E returns the number of edges.
func (g *Digraph) E() int {
	return g.e
}

// Reverse returns the reverse digraph.
func (g *Digraph) Reverse() *Digraph {
	r := NewDigraph(g.V())
	for v := 0; v < g.V(); v++ {
		adj := g.Adj(v)
		for _, w := range adj {
			r.Edge(w, v)
		}
	}
	return r
}
