package datastruct

import (
	"fmt"
	"io"
	"log"

	"github.com/goccy/go-graphviz"
	"github.com/goccy/go-graphviz/cgraph"
)

func RenderDiGraph(g *Digraph, format graphviz.Format, w io.Writer) error {
	gv := graphviz.New()
	grv, err := gv.Graph(graphviz.Directed)
	if err != nil {
		return err
	}
	defer func() {
		err := grv.Close()
		if err != nil {
			log.Fatal(err)
		}

		gv.Close()
	}()
	nodes := make(map[int]*cgraph.Node)
	for v := 0; v < g.V(); v++ {
		n, err := grv.CreateNode(fmt.Sprint(v))
		if err != nil {
			return err
		}
		nodes[v] = n
	}
	for v := 0; v < g.V(); v++ {
		for _, w := range g.Adj(v) {
			_, err := grv.CreateEdge("", nodes[v], nodes[w])
			if err != nil {
				return err
			}

		}
	}
	return gv.Render(grv, format, w)
}
