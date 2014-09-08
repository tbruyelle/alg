package quickunion

type QuickUnion []int

func New(N int) QuickUnion {
	qu := make([]int, N)
	for i := 0; i < N; i++ {
		qu[i] = i
	}
	return qu
}

func (qu QuickUnion) Connected(p, q int) bool {
	return qu.root(p) == qu.root(q)
}

func (qu QuickUnion) Union(p, q int) {
	qu[qu.root(p)] = qu.root(q)
}

func (qu QuickUnion) root(p int) int {
	for p != qu[p] {
		p = qu[p]
	}
	return p
}
