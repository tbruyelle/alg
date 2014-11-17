package datastruct

// UnionFind implements the weighted quick union algorythm.
type UnionFind struct {
	id []int
	sz []int
}

func NewUnionFind(N int) UnionFind {
	u := UnionFind{}
	u.id = make([]int, N)
	u.sz = make([]int, N)
	for i := 0; i < N; i++ {
		u.id[i] = i
		u.sz[i] = 1
	}
	return u
}

func (u UnionFind) Connected(p, q int) bool {
	return u.root(p) == u.root(q)
}

func (u UnionFind) Union(p, q int) {
	i := u.root(p)
	j := u.root(q)
	if i == j {
		return
	}
	if u.sz[i] < u.sz[j] {
		u.id[i] = j
		u.sz[j] += u.sz[i]
	} else {
		u.id[j] = i
		u.sz[i] += u.sz[j]
	}
}

func (u UnionFind) root(p int) int {
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}
