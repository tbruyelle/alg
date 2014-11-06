package graphs

import (
	"github.com/tbruyelle/alg/datastruct"
)

// Dfo implements the Deep-first order algorythm, aka topological sort.
type Dfo struct {
	marked []bool
	// All vertices of g in "Reverse dfs postorder".
	ReversePost []int
}

// NewDfo computes the "Reverse dfs postorder" on g.
func NewDfo(g *datastruct.Digraph) *Dfo {
	d := &Dfo{}
	d.marked = make([]bool, g.V())

	for v := 0; v < g.V(); v++ {
		if !d.marked[v] {
			d.dfs(g, v)
		}
	}
	// Reverse the ReversePost array
	for i, j := 0, len(d.ReversePost)-1; i < j; i, j = i+1, j-1 {
		d.ReversePost[i], d.ReversePost[j] = d.ReversePost[j], d.ReversePost[i]
	}

	return d
}

func (d *Dfo) dfs(g *datastruct.Digraph, v int) {
	d.marked[v] = true
	for _, w := range g.Adj(v) {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}
	d.ReversePost = append(d.ReversePost, v)
}
