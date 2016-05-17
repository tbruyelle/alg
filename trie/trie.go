package trie

import (
	"fmt"
)

// Trie implements a R-way trie
type Trie struct {
	root  *node
	Radix int
}

type node struct {
	next []*node
	val  interface{}
}

func New(radix int) *Trie {
	t := &Trie{
		Radix: radix,
	}
	return t
}

func (t *Trie) Contain(key string) bool {
	return t.Get(key) != nil
}

func (t *Trie) Get(key string) interface{} {
	x := get(t.root, key, 0)
	if x == nil {
		return nil
	}
	return x.val
}

func get(x *node, key string, d int) *node {
	if x == nil {
		return nil
	}
	if len(key) == d {
		// Found the word
		return x
	}
	c := key[d]
	return get(x.next[c], key, d+1)
}

func (t *Trie) Put(key string, val interface{}) {
	t.root = put(t, t.root, key, val, 0)
}

func put(t *Trie, x *node, key string, val interface{}, d int) *node {
	if x == nil {
		x = t.newNode()
	}
	if len(key) == d {
		x.val = val
		return x
	}
	c := key[d]
	x.next[c] = put(t, x.next[c], key, val, d+1)
	return x
}

func (t *Trie) newNode() *node {
	n := &node{}
	n.next = make([]*node, t.Radix)
	return n
}

func (t *Trie) String() string {
	return t.root.String('@')
}

func (x *node) String(key byte) string {
	if x == nil {
		return ""
	}
	s := fmt.Sprintf("(%q)\n", key)
	for i := 0; i < len(x.next); i++ {
		s += x.next[i].String(byte(i))
	}
	return s
}
