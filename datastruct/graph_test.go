package datastruct

import (
	"os"
	"testing"

	"github.com/goccy/go-graphviz"
)

func TestNewGraph(t *testing.T) {
	V := 10

	g := NewGraph(V)

	if g.V() != V {
		t.Errorf("Graph V=%d, want %d", g.V(), V)
	}
	adj := g.Adj(0)
	if l := len(adj); l > 0 {
		t.Errorf("Adj(0) len=%d, want 0", l)
	}
}

func TestGraphEdge(t *testing.T) {
	g := NewGraph(10)

	g.Edge(0, 1)
	g.Edge(0, 2)
	g.Edge(0, 5)

	adj := g.Adj(0)
	if l := len(adj); l != 3 {
		t.Fatalf("Ajd(0) len=%d, want 3", l)
	}
	if adj[0] != 1 {
		t.Errorf("adj[0]=%d, want 1", adj)
	}
	if adj[1] != 2 {
		t.Errorf("adj[1]=%d, want 2", adj)
	}
	if adj[2] != 5 {
		t.Errorf("adj[2]=%d, want 5", adj)
	}
}

func TestDupV(t *testing.T) {
	g := NewDigraph(4)
	g.Edge(0, 1)
	g.Edge(1, 2)
	g.Edge(2, 3)

	dupVert(g, 0)

	f, err := os.Create("/tmp/graph.png")
	if err != nil {
		t.Fatal(err)
	}
	RenderDiGraph(g, graphviz.PNG, f)
}

var nbDup = []int{0, 1, 1, 0}

func dupVert(g *Digraph, v int) {
	for i := 0; i < nbDup[v]; i++ {
		g.v++
		g.adj = append(g.adj, []int{})
		newv := g.v - 1
		// add edge from parents of v to newv
		for _, w := range g.Reverse().adj[v] {
			// add condition here to prevent add edges to wrong parents
			g.adj[w] = append(g.adj[w], newv)
		}
		// add edge from child of v to newv
		for _, w := range g.adj[v] {
			g.adj[newv] = append(g.adj[newv], w)
		}
	}

	for _, w := range g.adj[v] {
		dupVert(g, w)
	}

}
