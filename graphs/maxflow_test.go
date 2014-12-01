package graphs

import (
	"github.com/tbruyelle/alg/datastruct"
	"testing"
)

func TestMaxFlow(t *testing.T) {
	g := datastruct.NewFlowGraph(8)
	g.Edge(0, 1, 0, 10)
	g.Edge(0, 2, 0, 5)
	g.Edge(0, 3, 0, 15)
	g.Edge(1, 6, 0, 6)
	g.Edge(1, 5, 0, 15)
	g.Edge(1, 2, 0, 4)
	g.Edge(2, 5, 0, 8)
	g.Edge(2, 3, 0, 4)
	g.Edge(3, 4, 0, 16)
	g.Edge(4, 2, 0, 6)
	g.Edge(4, 7, 0, 10)
	g.Edge(5, 4, 0, 15)
	g.Edge(5, 7, 0, 10)
	g.Edge(6, 5, 0, 15)
	g.Edge(6, 7, 0, 10)

	f := NewMaxFlow(g, 0, 7)

	if maxFlow := float64(26); f.Value != maxFlow {
		t.Errorf("MaxFlow = %d, want %d", f.Value, maxFlow)
	}
}
