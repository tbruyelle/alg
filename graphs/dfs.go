package graphs

import (
	"github.com/tbruyelle/alg/datastruct"
)

// Dfs implements the Deep-first search algorythm
type Dfs struct {
	s      int
	marked []bool
	edgeTo []int
}

// NewDfs performs a dfs algorythm on g, starting from vertex s.
func NewDfs(g *datastruct.Graph, s int) *Dfs {
	d := &Dfs{s: s}
	d.marked = make([]bool, g.V)
	d.edgeTo = make([]int, g.V)

	d.dfs(g, s)

	return d
}

func (d *Dfs) dfs(g *datastruct.Graph, v int) {
	d.marked[v] = true
	for _, w := range g.Adj(v) {
		if !d.marked[w] {
			d.dfs(g, w)
			d.edgeTo[w] = v
		}
	}
}

// HasPathTo returns true if v and s are connected.
func (d *Dfs) HasPathTo(v int) bool {
	return d.marked[v]
}

// PathTo returns the path from v to s.
func (d *Dfs) PathTo(v int) []int {
	path := make([]int, 0)
	if !d.HasPathTo(v) {
		return path
	}
	for i := v; i != d.s; i = d.edgeTo[i] {
		path = append(path, i)
	}
	path = append(path, d.s)
	return path
}
