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

	want := []int{4, 1, 2, 5, 0, 6, 3}
	if !reflect.DeepEqual(d.ReversePost, want) {
		t.Errorf("ReversePost %v, want %v", d.ReversePost, want)
	}
}
