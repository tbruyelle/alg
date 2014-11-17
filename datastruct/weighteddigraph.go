package datastruct

// WeightedDigraph represents the graph data structure.
type WeightedDigraph struct {
	v     int
	e     int
	adj   [][]DirectedEdge
	edges []DirectedEdge
}

type DirectedEdge struct {
	v, w   int
	Weight int
}

// NewWeightedDigraph returns an adjacency-list graph representation.
func NewWeightedDigraph(v int) *WeightedDigraph {
	g := &WeightedDigraph{v: v}
	g.adj = make([][]DirectedEdge, v)
	return g
}

// DirectedEdge adds a new edge to the graph.
func (g *WeightedDigraph) Edge(v, w, Weight int) {
	g.adj[v] = append(g.adj[v], DirectedEdge{v, w, Weight})
	g.e++
	g.edges = append(g.edges, DirectedEdge{v, w, Weight})
}

// DirectedEdges returns all the added edges.
func (g *WeightedDigraph) Edges() []DirectedEdge {
	return g.edges
}

// Adj returns the adjacent vertices of the vertex v.
func (g *WeightedDigraph) Adj(v int) []DirectedEdge {
	return g.adj[v]
}

// V returns the number of vertices.
func (g *WeightedDigraph) V() int {
	return g.v
}

// E returns the number of edges.
func (g *WeightedDigraph) E() int {
	return g.e
}

func (e DirectedEdge) From() int {
	return e.v
}

func (e DirectedEdge) To() int {
	return e.w
}

// Priority implements the Item interface from PriorityQueue.
func (e DirectedEdge) Priority() int {
	return e.Weight
}
