package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	fmt.Println("MergeSort")
	a := []int{98, 60, 20, 73, 58, 13, 64, 43, 46, 51, 18, 19}
	fmt.Println(a)

	MergeSort(a)

	fmt.Println(a)
}

func TestHSort(t *testing.T) {
	fmt.Println("HSort")
	a := []int{43, 71, 67, 57, 29, 42, 53, 44, 24, 78}
	fmt.Println(a)

	HSort(4, a)

	fmt.Println(a)
}
