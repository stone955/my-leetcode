package sort

import "testing"

func TestMergeSort(t *testing.T) {
	in := []int{4, 1, 5, 2, 6, 3, 7, 9, 8, 0}
	MergeSort(in)
}
