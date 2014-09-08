package unionfind

import (
	"fmt"
	"testing"

	"github.com/tbruyelle/alg/uf/qf"
)

func TestQuickFind(t *testing.T) {
	qf := quickfind.New(10)
	qf.Union(9, 3)
	qf.Union(6, 0)
	qf.Union(2, 1)
	qf.Union(9, 0)
	qf.Union(8, 9)
	qf.Union(8, 1)

	fmt.Println(qf)
}
