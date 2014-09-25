package sort

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	a := []int{76, 90, 57, 19, 11, 79, 43, 21, 83, 73, 67, 59}
	fmt.Println("Partition", a)

	Partition(a, 0, len(a)-1)

	fmt.Println("Partition result:", a)
}
func TestPartition2(t *testing.T) {
	a := []string{"A", "B", "B", "A", "A", "B", "A", "B", "A", "A", "B", "B"}
	fmt.Println("Partition", a)

	PartitionStr(a, 0, len(a)-1)

	fmt.Println("Partition result:", a)
}
