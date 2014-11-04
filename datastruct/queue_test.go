package datastruct

import (
	"testing"
)

func TestQueue(t *testing.T) {
	q := NewQueue()

	q.Push(1)
	q.Push(4)
	q.Push(2)

	for _, x := range []int{1, 4, 2} {
		if y := q.Pull().(int); y != x {
			t.Errorf("Pull returns %d, want %d", y, x)
		}
	}
}
