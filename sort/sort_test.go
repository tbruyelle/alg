package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	fmt.Println("MergeSort")
	a := []int{29, 74, 47, 13, 61, 21, 56, 80, 42, 88, 67, 84}
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
