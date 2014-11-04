package dfs

import (
	"github.com/tbruyelle/alg/graphs/graph"
	"testing"
)

func TestHasPathTo(t *testing.T) {
	g := graph.New(10)
	g.Edge(0, 1)
	g.Edge(1, 2)
	g.Edge(2, 5)
	g.Edge(3, 8)
	g.Edge(3, 7)

	d := New(g, 0)

	if !d.HasPathTo(1) {
		t.Errorf("0 should have path to 1")
	}
	if !d.HasPathTo(2) {
		t.Errorf("0 should have path to 2")
	}
	if !d.HasPathTo(5) {
		t.Errorf("0 should have path to 5")
	}
	if d.HasPathTo(3) {
		t.Errorf("0 should not have path to 3")
	}
	if d.HasPathTo(8) {
		t.Errorf("0 should not have path to 8")
	}
	if d.HasPathTo(7) {
		t.Errorf("0 should not have path to 7")
	}
}

func TestPathTo(t *testing.T) {
	g := graph.New(10)
	g.Edge(0, 1)
	g.Edge(1, 2)
	g.Edge(2, 5)
	g.Edge(3, 8)
	g.Edge(3, 7)

	d := New(g, 0)
	path := d.PathTo(5)

	if l := len(path); l != 4 {
		t.Fatalf("path len=%d, want 4", l)
	}
	want := []int{5, 2, 1, 0}
	for i := range path {
		if path[i] != want[i] {
			t.Errorf("Path[%d]=%d, want %d", i, path[i], want[i])
		}
	}
}
