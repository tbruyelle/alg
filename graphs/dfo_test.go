package graphs

import (
	"github.com/tbruyelle/alg/datastruct"
	"reflect"
	"testing"
)

func TestDfo(t *testing.T) {
	g := datastruct.NewDigraph(7)
	g.Edge(0, 5)
	g.Edge(0, 2)
	g.Edge(0, 1)
	g.Edge(3, 6)
	g.Edge(3, 5)
	g.Edge(3, 4)
	g.Edge(5, 2)
	g.Edge(6, 4)
	g.Edge(6, 0)
	g.Edge(3, 2)
	g.Edge(1, 4)
	g.Edge(2, 1)

	d := NewDfo(g)

	want := []int{3, 6, 0, 5, 2, 1, 4}
	if !reflect.DeepEqual(d.ReversePost, want) {
		t.Errorf("ReversePost %v, want %v", d.ReversePost, want)
	}
}

func TestDfoNotDAG(t *testing.T) {
	g := datastruct.NewDigraph(3)
	g.Edge(0, 1)
	g.Edge(1, 2)
	g.Edge(2, 0)

	d := NewDfo(g)

	// Not implemented for now, should return no reversePost
	// because the graph isn't a DAG.
	_ = d
}
