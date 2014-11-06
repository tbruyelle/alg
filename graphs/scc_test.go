package graphs

import (
	"github.com/tbruyelle/alg/datastruct"
	"reflect"
	"testing"
)

func TestSCC(t *testing.T) {
	g := datastruct.NewDigraph(13)
	for _, v := range [][]int{{4, 2}, {2, 3}, {3, 2}, {6, 0}, {0, 1}, {2, 0}, {11, 12}, {12, 9}, {9, 10}, {9, 11}, {7, 9}, {10, 12}, {11, 4}, {4, 3}, {3, 5}, {6, 8}, {8, 6}, {5, 4}, {0, 5}, {6, 4}, {6, 9}, {7, 6}} {
		g.Edge(v[0], v[1])
	}

	s := NewSCC(g)

	want := []int{1, 0, 1, 1, 1, 1, 3, 4, 3, 2, 2, 2, 2}
	if !reflect.DeepEqual(s.id, want) {
		t.Errorf("SCC id=%v, want %v", s.id, want)
	}
}
