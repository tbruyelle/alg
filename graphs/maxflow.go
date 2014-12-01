package graphs

import (
	"github.com/tbruyelle/alg/datastruct"
	"math"
)

// MaxFlow implements the Ford-Fulkerson algorithm to compute the max flow of a graph.
type MaxFlow struct {
	marked []bool
	edgeTo []*datastruct.FlowEdge
	Value  float64
}

func NewMaxFlow(g *datastruct.FlowGraph, s, t int) *MaxFlow {
	f := &MaxFlow{}

	for f.hasAugmentingPath(g, s, t) {
		bottle := math.MaxFloat64
		for v := t; v != s; v = f.edgeTo[v].Other(v) {
			bottle = math.Min(bottle, f.edgeTo[v].ResidualCapacity(v))
		}
		for v := t; v != s; v = f.edgeTo[v].Other(v) {
			f.edgeTo[v].AddResidualFlow(v, bottle)
		}
		f.Value += bottle
	}

	return f
}

func (f *MaxFlow) hasAugmentingPath(g *datastruct.FlowGraph, s, t int) bool {
	f.marked = make([]bool, g.V())
	f.edgeTo = make([]*datastruct.FlowEdge, g.V())

	q := datastruct.NewQueue()
	q.Push(s)
	f.marked[s] = true
	for !q.Empty() {
		v := q.Pull().(int)
		for _, e := range g.Adj(v) {
			w := e.Other(v)
			if e.ResidualCapacity(w) > 0 && !f.marked[w] {
				f.edgeTo[w] = e
				f.marked[w] = true
				q.Push(w)
			}
		}
	}
	return f.marked[t]
}

func (f *MaxFlow) InCut(v int) bool {
	return f.marked[v]
}
