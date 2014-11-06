package datastruct

import (
	"reflect"
	"testing"
)

func TestNewDigraph(t *testing.T) {
	V := 10

	g := NewDigraph(V)

	if g.V() != V {
		t.Errorf("Digraph V=%d, want %d", g.V(), V)
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

	want := []int{1, 2, 5}
	if !reflect.DeepEqual(g.Adj(0), want) {
		t.Errorf("Adj(0)=%v, want %v", g.Adj(0), want)
	}
}

func TestDigraphReverse(t *testing.T) {
	g := NewDigraph(4)
	g.Edge(0, 1)
	g.Edge(0, 2)
	g.Edge(0, 3)
	g.Edge(2, 1)
	g.Edge(1, 0)

	r := g.Reverse()

	want := [][]int{{1}, {0, 2}, {0}, {0}}
	for v := 0; v < r.V(); v++ {
		if !reflect.DeepEqual(r.Adj(v), want[v]) {
			t.Errorf("Reverse Adj(%d)=%v, want %v", v, r.Adj(v), want[v])
		}
	}
}
