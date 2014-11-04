package graphs

import (
	"github.com/tbruyelle/alg/datastruct"
	"testing"
)

func TestBfsHasPathTo(t *testing.T) {
	g := datastruct.NewGraph(10)
	g.Edge(0, 1)
	g.Edge(1, 2)
	g.Edge(2, 5)
	g.Edge(3, 8)
	g.Edge(3, 7)

	b := NewBfs(g, 0)

	if !b.HasPathTo(1) {
		t.Errorf("0 should have path to 1")
	}
	if !b.HasPathTo(2) {
		t.Errorf("0 should have path to 2")
	}
	if !b.HasPathTo(5) {
		t.Errorf("0 should have path to 5")
	}
	if b.HasPathTo(3) {
		t.Errorf("0 should not have path to 3")
	}
	if b.HasPathTo(8) {
		t.Errorf("0 should not have path to 8")
	}
	if b.HasPathTo(7) {
		t.Errorf("0 should not have path to 7")
	}
}

func TestBfsPathTo(t *testing.T) {
	g := datastruct.NewGraph(10)
	g.Edge(0, 1)
	g.Edge(1, 2)
	g.Edge(2, 5)
	g.Edge(3, 8)
	g.Edge(3, 7)

	b := NewBfs(g, 0)
	path := b.PathTo(5)

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

func TestBfsDistanceTo(t *testing.T) {
	g := datastruct.NewGraph(11)
	g.Edge(0, 1)
	g.Edge(1, 2)
	g.Edge(0, 2)
	g.Edge(2, 5)
	g.Edge(5, 6)
	g.Edge(5, 9)
	g.Edge(9, 10)

	b := NewBfs(g, 0)

	data := []struct{ v, want int }{{1, 1}, {2, 1}, {5, 2}, {6, 3}, {9, 3}, {10, 4}, {3, 0}}
	for _, d := range data {
		if dist := b.DistanceTo(d.v); dist != d.want {
			t.Errorf("DistanceTo(%d)=%d, want %d", d.v, dist, d.want)
		}
	}
}
