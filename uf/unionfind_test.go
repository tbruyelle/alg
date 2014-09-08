package unionfind

import (
	"fmt"
	"testing"

	"github.com/tbruyelle/alg/uf/qf"
	"github.com/tbruyelle/alg/uf/qu"
	"github.com/tbruyelle/alg/uf/wqu"
)

func TestQuickFind(t *testing.T) {
	qf := quickfind.New(10)
	qf.Union(9, 3)
	qf.Union(6, 0)
	qf.Union(2, 1)
	qf.Union(9, 0)
	qf.Union(8, 9)
	qf.Union(8, 1)

	fmt.Println("QuickFind", qf)
}

func TestQuickUnion(t *testing.T) {
	qu := quickunion.New(10)
	qu.Union(8, 5)
	qu.Union(9, 0)
	qu.Union(0, 7)
	qu.Union(6, 1)
	qu.Union(1, 8)
	qu.Union(9, 3)
	qu.Union(7, 2)
	qu.Union(5, 9)
	qu.Union(8, 4)

	fmt.Println("QuickUnion", qu)
}

func TestWeightedQuickUnion(t *testing.T) {
	qu := weightedquickunion.New(10)
	qu.Union(8, 5)
	qu.Union(9, 0)
	qu.Union(0, 7)
	qu.Union(6, 1)
	qu.Union(1, 8)
	qu.Union(9, 3)
	qu.Union(7, 2)
	qu.Union(5, 9)
	qu.Union(8, 4)

	fmt.Printf("WeightedQuickUnion %+v", qu)
}
