package graphs

import (
	"github.com/tbruyelle/alg/datastruct"
	"testing"
)

func TestCC(t *testing.T) {
	g := datastruct.NewGraph(7)
	g.Edge(0, 1)
	g.Edge(1, 2)
	g.Edge(0, 2)
	g.Edge(2, 5)
	g.Edge(3, 4)
	g.Edge(3, 6)

	c := NewCC(g)

	if count := 2; count != c.Count {
		t.Errorf("CC count=%d, want %d", c.Count, count)
	}
	for _, v := range []int{0, 1, 2, 5} {
		if !c.Connected(0, v) {
			t.Errorf("Connected(0,%d) should returns true", v)
		}
		if c.Id[v] != 0 {
			t.Errorf("Id of %d = %d, want 0", v, c.Id[v])
		}
	}
	for _, v := range []int{3, 6, 4} {
		if !c.Connected(3, v) {
			t.Errorf("Connected(3,%d) should returns true", v)
		}
		if c.Id[v] != 1 {
			t.Errorf("Id of %d = %d, want 1", v, c.Id[v])
		}
	}
	for _, v := range []int{3, 6, 4} {
		if c.Connected(0, v) {
			t.Errorf("Connected(0,%d) should returns false", v)
		}
	}
}
