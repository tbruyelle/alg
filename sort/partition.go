package sort

import "fmt"

func Partition(a []int, lo, hi int) int {
	i, j := lo+1, hi
	for {
		for a[i] < a[lo] {
			i++
			if i == hi {
				break
			}
		}
		fmt.Println("i", i, a[i], a[lo])
		for a[lo] < a[j] {
			j--
			if j == lo {
				break
			}
		}
		fmt.Println("j", j, a[j], a[lo])
		if i >= j {
			fmt.Println("break", i, j)
			break
		}
		a[i], a[j] = a[j], a[i]
		fmt.Println("Swap", a)
	}
	a[lo], a[j] = a[j], a[lo]
	return j
}

func PartitionStr(a []string, lo, hi int) int {
	i, j := lo, hi+1
	for {
		for i++; a[i] < a[lo] && i < hi; i++ {
		}
		fmt.Println("i", i, a[i], a[lo])
		for j--; a[lo] < a[j] && j > lo; j-- {
		}
		fmt.Println("j", j, a[j], a[lo])
		if i >= j {
			fmt.Println("break", i, j)
			break
		}
		a[i], a[j] = a[j], a[i]
		fmt.Println("Swap", a)
	}
	a[lo], a[j] = a[j], a[lo]
	return j
}
