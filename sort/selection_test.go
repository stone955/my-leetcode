package sort

import "testing"

func TestSelectionSort(t *testing.T) {
	in := []int{4, 7, 6, 3, 8, 1, 2, 5, 9, 0}
	SelectionSort(in)
}
