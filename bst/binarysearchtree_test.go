package binarysearchtree

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

func TestDelete(t *testing.T) {
	b := New()
	val := 1
	//  32 26 60 11 40 89 13 35 48 74 73 84
	b.Put(32, &val)
	b.Put(26, &val)
	b.Put(60, &val)
	b.Put(11, &val)
	b.Put(40, &val)
	b.Put(89, &val)
	b.Put(13, &val)
	b.Put(35, &val)
	b.Put(48, &val)
	b.Put(74, &val)
	b.Put(73, &val)
	b.Put(84, &val)

	b.Delete(13)
	b.Delete(26)
	b.Delete(60)

	fmt.Println("delete", b)
}
