package graphs

import (
	"github.com/tbruyelle/alg/datastruct"
)

// SCC implements the search of Strong Connected Components
// via the Kosaraju Sharir algoryhtm.
type SCC struct {
	id     []int
	marked []bool
	count  int
}

// NewSCC returns a new SCC with strongly connected components
// search done.
func NewSCC(g *datastruct.Digraph) *SCC {
	s := &SCC{}
	s.id = make([]int, g.V())
	s.marked = make([]bool, g.V())

	// First run a dfo on the reverse graph.
	dfo := NewDfo(g.Reverse())
	// Then run a dfs in the reverse post order.
	for _, v := range dfo.ReversePost {
		if !s.marked[v] {
			s.dfs(g, v)
			s.count++
		}
	}
	return s
}

func (s *SCC) dfs(g *datastruct.Digraph, v int) {
	s.marked[v] = true
	s.id[v] = s.count
	for _, w := range g.Adj(v) {
		if !s.marked[w] {
			s.dfs(g, w)
		}
	}
}

// StronglyConnected returns true if v and w are strongly
// connected.
func (s *SCC) StronglyConnected(v, w int) bool {
	return s.id[w] == s.id[v]
}
