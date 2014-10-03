package binarysearchtree

import (
	"fmt"
)

type BinarySearchTree struct {
	root *node
}

type node struct {
	key         int
	value       *int
	left, right *node
	count       int
}

func New() *BinarySearchTree {
	b := &BinarySearchTree{}
	return b
}

// Level-order traversal representation of the BST
func (b *BinarySearchTree) String() string {
	var res string
	for i := 1; i <= height(b.root); i++ {
		res += str(b.root, i)
	}
	return res
}

func str(x *node, lvl int) string {
	if x == nil {
		return ""
	}
	if lvl == 1 {
		return fmt.Sprintf("%d ", x.key)
	} else if lvl > 1 {
		return str(x.left, lvl-1) + str(x.right, lvl-1)
	}
	return ""
}

func height(x *node) int {
	if x == nil {
		return 0
	}
	lh := height(x.left)
	rh := height(x.right)
	if lh > rh {
		return lh + 1
	} else {
		return rh + 1
	}
}

func (b *BinarySearchTree) Size() int {
	return size(b.root)
}

func size(x *node) int {
	if x == nil {
		return 0
	}
	return x.count
}

func (b *BinarySearchTree) Get(key int) *int {
	x := b.root
	for x != nil {
		cmp := compare(key, x.key)
		switch {
		case cmp < 0:
			x = x.left
		case cmp > 0:
			x = x.right
		default:
			return x.value
		}
	}
	return nil
}

func (b *BinarySearchTree) Put(key int, value *int) {
	b.root = put(b.root, key, value)
}

func put(x *node, key int, value *int) *node {
	if x == nil {
		return &node{key: key, value: value, count: 1}
	}
	cmp := compare(key, x.key)
	switch {
	case cmp < 0:
		x.left = put(x.left, key, value)
	case cmp > 0:
		x.right = put(x.right, key, value)
	default:
		x.value = value
	}
	x.count = 1 + size(x.left) + size(x.right)
	return x
}

// Hibbard deletion
func (b *BinarySearchTree) Delete(key int) {
	b.root = del(b.root, key)
}

func del(x *node, key int) *node {
	if x == nil {
		return nil
	}
	cmp := compare(key, x.key)
	switch {
	case cmp < 0:
		x.left = del(x.left, key)
	case cmp > 0:
		x.right = del(x.right, key)
	default:
		if x.left == nil {
			return x.right
		}
		if x.right == nil {
			return x.left
		}
		t := x
		x = min(t.right)
		x.right = delMin(t.right)
		x.left = t.left
	}
	x.count = 1 + size(x.left) + size(x.right)
	return x
}

func (b *BinarySearchTree) DeleteMin() {
	b.root = delMin(b.root)
}

func delMin(x *node) *node {
	if x == nil {
		return nil
	}
	if x.left == nil {
		return x.right
	}
	x.left = delMin(x.left)
	x.count = 1 + size(x.left) + size(x.right)
	return x
}

func (b *BinarySearchTree) Rank(key int) int {
	return rank(b.root, key)
}

func rank(x *node, key int) int {
	if x == nil {
		return 0
	}
	cmp := compare(key, x.key)
	switch {
	case cmp < 0:
		return rank(x.left, key)
	case cmp > 0:
		return 1 + size(x.left) + rank(x.right, key)
	default:
		return size(x.left)
	}
}

func (b *BinarySearchTree) Min() *int {
	x := min(b.root)
	if x == nil {
		return nil
	}
	return x.value
}

func min(x *node) *node {
	if x == nil {
		return nil
	}
	if x.left == nil {
		return x
	}
	return min(x.left)
}

func (b *BinarySearchTree) Max() *int {
	x := max(b.root)
	if x == nil {
		return nil
	}
	return x.value
}

func max(x *node) *node {
	if x == nil {
		return nil
	}
	if x.right == nil {
		return x
	}
	return max(x.right)
}

func compare(key1, key2 int) int {
	return key1 - key2
}
