// Adjacency-list graph representation
package graph

type Graph struct {
	V, E int
	adj  [][]int
}

func New(V int) *Graph {
	g := &Graph{V: V}
	g.adj = make([][]int, V)
	return g
}

func (g *Graph) Edge(v, w int) {
	g.adj[v] = append(g.adj[v], w)
	g.adj[w] = append(g.adj[w], v)
	g.E++
}

func (g *Graph) Adj(v int) []int {
	return g.adj[v]
}
