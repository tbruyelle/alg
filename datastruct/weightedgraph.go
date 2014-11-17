package datastruct

// WeightedGraph represents the graph data structure.
type WeightedGraph struct {
	v     int
	e     int
	adj   [][]Edge
	edges []Edge
}

type Edge struct {
	v, w   int
	Weight int
}

// NewWeightedGraph returns an adjacency-list graph representation.
func NewWeightedGraph(v int) *WeightedGraph {
	g := &WeightedGraph{v: v}
	g.adj = make([][]Edge, v)
	return g
}

// Edge adds a new edge to the graph.
func (g *WeightedGraph) Edge(v, w, Weight int) {
	g.adj[v] = append(g.adj[v], Edge{v, w, Weight})
	g.adj[w] = append(g.adj[w], Edge{w, v, Weight})
	g.e++
	g.edges = append(g.edges, Edge{v, w, Weight})
}

// Edges returns all the added edges.
func (g *WeightedGraph) Edges() []Edge {
	return g.edges
}

// Adj returns the adjacent vertices of the vertex v.
func (g *WeightedGraph) Adj(v int) []Edge {
	return g.adj[v]
}

// V returns the number of vertices.
func (g *WeightedGraph) V() int {
	return g.v
}

// E returns the number of edges.
func (g *WeightedGraph) E() int {
	return g.e
}

func (e Edge) Either() int {
	return e.v
}

func (e Edge) Other(v int) int {
	if v == e.v {
		return e.w
	}
	return e.v
}

// Priority implements the Item interface from PriorityQueue.
func (e Edge) Priority() int {
	return e.Weight
}
