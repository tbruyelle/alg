package datastruct

import (
	"testing"
)

func TestUnionFind(t *testing.T) {
	u := NewUnionFind(4)
	u.Union(0, 1)
	u.Union(1, 2)

	want := []bool{true, true, false, false}
	for i := 0; i < 3; i++ {
		if c := u.Connected(i, i+1); c != want[i] {
			t.Errorf("Connected(%d,%d)=%t, want %t", i, i+1, c, want[i])
		}
	}

}
