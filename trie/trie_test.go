package trie

import (
	"testing"
)

func TestTrie(t *testing.T) {
	tr := New(256)

	tr.Put("HELLO", "WORLD")
	tr.Put("WORLD", "HELLO")
	tr.Put("HEY", "GIRL")
	tr.Put("WAIT", "WAT")

	data := []struct {
		key  string
		want bool
	}{
		{"HELLO", true},
		{"WORLD", true},
		{"HELL", false},
		{"yapikai", false},
		{"HEY", true},
		{"WAIT", true},
	}
	for i := 0; i < len(data); i++ {
		val := tr.Get(data[i].key)
		if !data[i].want && val != nil {
			t.Errorf("%d) Get(%s) returned %s, want nil", i, data[i].key, val)
		}
		if data[i].want && val == nil {
			t.Errorf("%d) Get(%s) returned nil", i, data[i].key)
		}
	}

}
