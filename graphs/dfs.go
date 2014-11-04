package graphs

import (
	"github.com/tbruyelle/alg/datastruct"
)

// Dfs implements the Deep-first search algorythm
type Dfs struct {
	g      *datastruct.Graph
	s      int
	marked []bool
	edgeTo []int
}

// New performs a dfs algorythm on g, starting from vertex s.
func NewDfs(g *datastruct.Graph, s int) *Dfs {
	d := &Dfs{g: g, s: s}
	d.marked = make([]bool, g.V)
	d.edgeTo = make([]int, g.V)

	d.dfs(s)

	return d
}

func (d *Dfs) dfs(v int) {
	d.marked[v] = true
	for _, w := range d.g.Adj(v) {
		if !d.marked[w] {
			d.dfs(w)
			d.edgeTo[w] = v
		}
	}
}

func (d *Dfs) HasPathTo(v int) bool {
	return d.marked[v]
}

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
