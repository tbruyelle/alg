package sort

import (
	"fmt"
	"testing"
)

func TestHSort(t *testing.T) {
	a := []int{43, 71, 67, 57, 29, 42, 53, 44, 24, 78}
	fmt.Println(a)

	HSort(4, a)

	fmt.Println(a)
}
