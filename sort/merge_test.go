package sort

import "testing"

func TestMergeSort(t *testing.T) {
	in := []int{1, 2, 3, 5, 7, 8, 0, 5, 6, 9}
	MergeSort(in)
}
