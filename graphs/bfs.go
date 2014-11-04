package graphs

import (
	"fmt"
	"github.com/tbruyelle/alg/datastruct"
)

// Bfs implements the Bread-first search algorythm.
type Bfs struct {
	s      int
	marked []bool
	edgeTo []int
}

// NewBfs performs a bfs algorythm on g, starting from vertex s.
func NewBfs(g *datastruct.Graph, s int) *Bfs {
	b := &Bfs{s: s}
	b.marked = make([]bool, g.V)
	b.edgeTo = make([]int, g.V)

	q := datastruct.NewQueue()
	q.Push(s)
	for {
		x := q.Pull()
		if x == nil {
			break
		}
		v := x.(int)
		for _, w := range g.Adj(v) {
			if !b.marked[w] {
				b.marked[w] = true
				b.edgeTo[w] = v
				q.Push(w)
			}
		}
	}

	return b
}

// HasPathTo returns true if v and s are connected.
func (b *Bfs) HasPathTo(v int) bool {
	return b.marked[v]
}

// PathTo returns the path from v to s.
func (b *Bfs) PathTo(v int) []int {
	path := make([]int, 0)
	if !b.HasPathTo(v) {
		return path
	}
	for x := v; x != b.s; x = b.edgeTo[x] {
		path = append(path, x)
	}
	path = append(path, b.s)
	return path
}
