package graphs

import (
	"container/heap"
	"github.com/tbruyelle/alg/datastruct"
	"math"
)

// SP implements the Djikstra Shortest Path algorithm.
type SP struct {
	distTo []int
	edgeTo []*datastruct.DirectedEdge
	pq     datastruct.PriorityQueue
}

func NewSP(g *datastruct.WeightedDigraph, s int) *SP {
	sp := &SP{}
	sp.distTo = make([]int, g.V())
	for i := 0; i < len(sp.distTo); i++ {
		sp.distTo[i] = math.MaxInt64
	}
	sp.distTo[s] = 0
	sp.edgeTo = make([]*datastruct.DirectedEdge, g.V())

	sp.pq = datastruct.NewMinPQ()
	heap.Push(&sp.pq, item{s, sp})

	for sp.pq.Len() > 0 {
		i := heap.Pop(&sp.pq).(item)

		for _, e := range g.Adj(i.v) {
			sp.relax(e)
		}
	}
	return sp
}

func (sp *SP) DistTo(v int) int {
	return sp.distTo[v]
}

func (sp *SP) PathTo(v int) []datastruct.DirectedEdge {
	var path []datastruct.DirectedEdge
	for e := sp.edgeTo[v]; e != nil; e = sp.edgeTo[e.From()] {
		path = append(path, *e)
	}
	// Reverse the path array
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func (sp *SP) relax(e datastruct.DirectedEdge) {
	v, w := e.From(), e.To()
	if sp.distTo[w] > sp.distTo[v]+e.Weight {
		sp.distTo[w] = sp.distTo[v] + e.Weight
		sp.edgeTo[w] = &e

		if i := sp.pqIndex(w); i >= 0 {
			heap.Fix(&sp.pq, i)
		} else {
			heap.Push(&sp.pq, item{w, sp})
		}
	}
}

func (sp *SP) pqIndex(v int) int {
	for i := range sp.pq {
		if sp.pq[i].(item).v == v {
			return i
		}
	}
	return -1

}

type item struct {
	v  int
	sp *SP
}

func (i item) Priority() int {
	return i.sp.distTo[i.v]
}
