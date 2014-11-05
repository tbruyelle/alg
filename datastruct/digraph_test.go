package datastruct

import (
	"testing"
)

func TestNewDigraph(t *testing.T) {
	V := 10

	g := NewDigraph(V)

	if g.V != V {
		t.Errorf("Digraph V=%d, want %d", g.V, V)
	}
	adj := g.Adj(0)
	if l := len(adj); l > 0 {
		t.Errorf("Adj(0) len=%d, want 0", l)
	}
}

func TestDigraphEdge(t *testing.T) {
	g := NewDigraph(10)

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
