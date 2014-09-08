package weightedquickunion

type WeightedQuickUnion struct {
	id []int
	sz []int
}

func New(N int) WeightedQuickUnion {
	qu := WeightedQuickUnion{}
	qu.id = make([]int, N)
	qu.sz = make([]int, N)
	for i := 0; i < N; i++ {
		qu.id[i] = i
		qu.sz[i] = 1
	}
	return qu
}

func (qu WeightedQuickUnion) Connected(p, q int) bool {
	return qu.root(p) == qu.root(q)
}

func (qu WeightedQuickUnion) Union(p, q int) {
	i := qu.root(p)
	j := qu.root(q)
	if i == j {
		return
	}
	if qu.sz[i] < qu.sz[j] {
		qu.id[i] = j
		qu.sz[j] += qu.sz[i]
	} else {
		qu.id[j] = i
		qu.sz[i] += qu.sz[j]
	}
}

func (qu WeightedQuickUnion) root(p int) int {
	for p != qu.id[p] {
		p = qu.id[p]
	}
	return p
}
