package sort

import "fmt"

// MergeSort sorts the array with the merge sort algorithm.
func MergeSort(a []int) {
	aux := make([]int, len(a))
	mergeSort(a, aux, 0, len(a)-1)
}

func BottomUpMergeSort(a []int) {
	//N := len(a)
	//aux := make([]int, N)

}

func mergeSort(a, aux []int, lo, hi int) {
	if lo >= hi {
		return
	}
	mid := lo + (hi-lo)/2
	mergeSort(a, aux, lo, mid)
	mergeSort(a, aux, mid+1, hi)
	merge(a, aux, lo, mid, hi)
	fmt.Println("merging", lo, mid, hi, a)
}

func merge(a, aux []int, lo, mid, hi int) {
	// Copy
	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}
	// Merge
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		switch {
		case i > mid:
			a[k] = aux[j]
			j++
		case j > hi:
			a[k] = aux[i]
			i++
		case aux[j] < aux[i]:
			a[k] = aux[j]
			j++
		default:
			a[k] = aux[i]
			i++
		}
	}
}

// HSort h-sorts the array given the h parameter.
func HSort(h int, a []int) {
	for i := h; i < len(a); i++ {
		for j := i; j >= h; j-- {
			if a[j-h] > a[j] {
				a[j], a[j-h] = a[j-h], a[j]
				fmt.Println(a)
			}
		}
	}
}
