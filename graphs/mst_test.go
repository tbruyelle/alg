package graphs

import (
	"fmt"
	"github.com/tbruyelle/alg/datastruct"
	"testing"
)

func TestMST(t *testing.T) {
	g := datastruct.NewWeightedGraph(8)
	g.Edge(0, 7, 16)
	g.Edge(2, 3, 17)
	g.Edge(1, 7, 19)
	g.Edge(0, 2, 26)
	g.Edge(5, 7, 28)
	g.Edge(1, 3, 29)
	g.Edge(1, 5, 32)
	g.Edge(2, 7, 34)
	g.Edge(4, 5, 35)
	g.Edge(1, 2, 36)
	g.Edge(4, 7, 37)
	g.Edge(0, 4, 38)
	g.Edge(6, 2, 40)
	g.Edge(3, 6, 52)
	g.Edge(6, 0, 58)
	g.Edge(6, 4, 93)

	mst := NewMST(g)

	fmt.Println(mst.Edges)
}
