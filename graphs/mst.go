package graphs

import (
	"container/heap"
	"github.com/tbruyelle/alg/datastruct"
)

// MST implements the Kruskal's Minimum Spanning Tree algorithm.
type MST struct {
	Edges []datastruct.Edge
}

func NewMST(g *datastruct.WeightedGraph) *MST {
	m := &MST{}

	pq := datastruct.NewMinPQ()
	for _, e := range g.Edges() {
		heap.Push(&pq, e)
	}
	uf := datastruct.NewUnionFind(g.V())
	for pq.Len() > 0 && len(m.Edges) < g.V()-1 {
		e := heap.Pop(&pq).(datastruct.Edge)
		v := e.Either()
		w := e.Other(v)
		if !uf.Connected(v, w) {
			uf.Union(v, w)
			m.Edges = append(m.Edges, e)
		}
	}
	return m
}
