package graphs

import (
	"github.com/tbruyelle/alg/datastruct"
	"testing"
)

func TestSP(t *testing.T) {
	g := datastruct.NewWeightedDigraph(8)
	g.Edge(0, 1, 5)
	g.Edge(0, 4, 9)
	g.Edge(0, 7, 8)
	g.Edge(1, 2, 12)
	g.Edge(1, 3, 15)
	g.Edge(1, 7, 4)
	g.Edge(2, 3, 3)
	g.Edge(2, 6, 11)
	g.Edge(3, 6, 9)
	g.Edge(4, 5, 4)
	g.Edge(4, 6, 20)
	g.Edge(4, 7, 5)
	g.Edge(5, 2, 1)
	g.Edge(5, 6, 13)
	g.Edge(7, 5, 6)
	g.Edge(7, 2, 7)

	sp := NewSP(g, 0)

	wantDistTo := []int{0, 5, 14, 17, 9, 13, 25, 8}
	wantEdgeTo := []string{"", "0->1", "5->2", "2->3", "0->4", "4->5", "2->6", "0->7"}
	for i := 1; i < g.V(); i++ {
		if sp.distTo[i] != wantDistTo[i] {
			t.Errorf("distTo[%d]=%d, want %d", i, sp.distTo[i], wantDistTo[i])
		}
		var e string
		if sp.edgeTo[i] != nil {
			e = sp.edgeTo[i].String()
		}
		if e != wantEdgeTo[i] {
			t.Errorf("edgeTo[i]=%s, want %s", i, e, wantEdgeTo[i])
		}
	}
}
