package hsort

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	a := []int{43, 71, 67, 57, 29, 42, 53, 44, 24, 78}
	fmt.Println(a)

	Sort(4, a)

	fmt.Println(a)
}
