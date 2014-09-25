package sort

import "fmt"

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
