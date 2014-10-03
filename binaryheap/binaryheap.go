package binaryheap

type BinaryHeap struct {
	a []int
	n int
}

func New(capacity int) *BinaryHeap {
	b := &BinaryHeap{}
	b.a = make([]int, capacity+1)
	return b
}

func (b *BinaryHeap) Empty() bool {
	return b.n == 0
}

func (b *BinaryHeap) Insert(key int) {
	b.n++
	b.a[b.n] = key
	b.swim(b.n)
}

func (b *BinaryHeap) DelMax() int {
	max := b.a[1]
	b.a[1], b.a[b.n] = b.a[b.n], b.a[1]
	b.n--
	b.a[b.n+1] = 0
	b.sink(1)
	return max
}

func (b *BinaryHeap) swim(k int) {
	for k > 1 && b.less(k/2, k) {
		b.a[k], b.a[k/2] = b.a[k/2], b.a[k]
		k = k / 2
	}
}

func (b *BinaryHeap) sink(k int) {
	for 2*k <= b.n {
		j := 2 * k
		if j < b.n && b.less(j, j+1) {
			j++
		}
		if !b.less(k, j) {
			break
		}
		b.a[k], b.a[j] = b.a[j], b.a[k]
		k = j
	}
}

func (b *BinaryHeap) less(i, j int) bool {
	return b.a[i] < b.a[j]
}
