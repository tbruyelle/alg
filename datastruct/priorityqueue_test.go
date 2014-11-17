package datastruct

import (
	"container/heap"
	"testing"
)

type item int

func (i item) Priority() int {
	return int(i)
}

func TestMinPQ(t *testing.T) {
	pq := NewMinPQ()
	heap.Push(&pq, item(2))
	heap.Push(&pq, item(0))
	heap.Push(&pq, item(3))
	heap.Push(&pq, item(1))

	want := []int{0, 1, 2, 3}
	for i := 0; i < 4; i++ {
		if d := int(heap.Pop(&pq).(item)); d != want[i] {
			t.Errorf("Pop() #%d returns %d, want %d", i, d, want[i])
		}
	}
}
