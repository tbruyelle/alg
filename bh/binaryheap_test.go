package binaryheap

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	b := New(13)
	//79 72 71 63 54 18 38 41 47 53
	b.Insert(79)
	b.Insert(38)
	b.Insert(72)
	b.Insert(63)
	b.Insert(71)
	b.Insert(54)
	b.Insert(53)
	b.Insert(18)
	b.Insert(41)
	b.Insert(47)
	fmt.Println("OK", b.a)

	max := 100000
	for !b.Empty() {
		key := b.DelMax()
		fmt.Println(b.a)
		if key > max {
			t.Errorf("Error %d should be > %d", key, max)
		}
		max = key
	}
}

func TestOrder(t *testing.T) {
	b := New(13)
	for i, v := range []int{79, 72, 71, 63, 54, 18, 38, 41, 47, 53} {
		b.a[i+1] = v
		b.n++
	}
	fmt.Println(b.a)

	b.Insert(99)
	fmt.Println(b.a)
	b.Insert(92)
	fmt.Println(b.a)
	b.Insert(82)
	fmt.Println(b.a)
}

func TestOrder2(t *testing.T) {
	b := New(10)
	for i, v := range []int{99, 71, 86, 67, 44, 12, 62, 41, 30, 19} {
		b.a[i+1] = v
		b.n++
	}
	fmt.Println("order2", b.a)

	b.DelMax()
	fmt.Println(b.a)
	b.DelMax()
	fmt.Println(b.a)
	b.DelMax()
	fmt.Println(b.a)
}
