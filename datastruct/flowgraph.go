package datastruct

// FlowGraph represents a graph data structure flavored for the maxflow/mincut algorithms.
type FlowGraph struct {
	v   int
	adj [][]*FlowEdge
}

type FlowEdge struct {
	v, w     int
	Flow     float64
	Capacity float64
}

// NewFlowDigraph returns an adjacency-list graph representation.
func NewFlowGraph(v int) *FlowGraph {
	g := &FlowGraph{v: v}
	g.adj = make([][]*FlowEdge, v)
	return g
}

// Edge adds a new edge to the graph.
func (g *FlowGraph) Edge(v, w int, flow, capacity float64) {
	e := &FlowEdge{v, w, flow, capacity}
	g.adj[v] = append(g.adj[v], e)
	g.adj[w] = append(g.adj[w], e)
}

func (g *FlowGraph) Adj(v int) []*FlowEdge {
	return g.adj[v]
}

func (g *FlowGraph) V() int {
	return g.v
}

func (g *FlowGraph) E() int {
	// not implemented
	return 0
}

func (e *FlowEdge) From() int {
	return e.v
}

func (e *FlowEdge) To() int {
	return e.w
}

func (e *FlowEdge) Other(v int) int {
	if v == e.v {
		return e.w
	}
	if v == e.w {
		return e.v
	}
	panic("Bad arg")
}

func (e *FlowEdge) ResidualCapacity(v int) float64 {
	if v == e.v {
		// Backward edge
		return e.Flow
	} else if v == e.w {
		// Forward edge
		return e.Capacity - e.Flow
	}
	panic("Bad arg")
}

func (e *FlowEdge) AddResidualFlow(v int, delta float64) {
	if v == e.v {
		// Backward edge
		e.Flow -= delta
	} else if v == e.w {
		// Forward edge
		e.Flow += delta
	} else {
		panic("Bad arg")
	}
}
