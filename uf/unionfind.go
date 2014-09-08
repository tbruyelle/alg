package unionfind

type UnionFinder interface {
	Connected(p, q int) bool
	Union(p, q int)
}
