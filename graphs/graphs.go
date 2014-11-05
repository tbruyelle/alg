package graphs

type Interface interface {
	Adj(v int) []int
	V() int
	E() int
}
