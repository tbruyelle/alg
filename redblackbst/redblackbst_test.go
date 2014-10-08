package redblackbst

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	b := New()
	val := 1
	b.Put(1, &val)

	res := b.Get(1)
	if res == nil || *res != 1 {
		t.Error("BST missing a value for key 1")
	}
}

func TestString(t *testing.T) {
	b := New()
	val := 1
	// 71 83 15 14 67 28 40 30 98 95
	b.Put(71, &val)
	b.Put(83, &val)
	b.Put(15, &val)
	b.Put(14, &val)
	b.Put(67, &val)
	b.Put(28, &val)
	b.Put(40, &val)
	b.Put(30, &val)
	b.Put(98, &val)
	b.Put(95, &val)

	fmt.Println("put", b)
}
