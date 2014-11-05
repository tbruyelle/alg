package graphs

import (
	"github.com/tbruyelle/alg/datastruct"
)

// CC implements the Connected-component algorythm.
type CC struct {
	marked []bool
	Id     []int
	Count  int
}

// NewCC performs a cc algorythm on g.
func NewCC(g *datastruct.Graph) *CC {
	c := &CC{}
	c.marked = make([]bool, g.V())
	c.Id = make([]int, g.V())

	for v := 0; v < g.V(); v++ {
		if !c.marked[v] {
			c.bfs(g, v)
			c.Count++
		}
	}
	return c
}

func (c *CC) bfs(g *datastruct.Graph, s int) {
	q := datastruct.NewQueue()
	q.Push(s)
	c.marked[s] = true
	c.Id[s] = c.Count
	for {
		x := q.Pull()
		if x == nil {
			break
		}
		v := x.(int)
		for _, w := range g.Adj(v) {
			if !c.marked[w] {
				c.marked[w] = true
				c.Id[w] = c.Count
				q.Push(w)
			}
		}
	}
}

func (c *CC) Connected(v, w int) bool {
	return c.Id[v] == c.Id[w]
}
