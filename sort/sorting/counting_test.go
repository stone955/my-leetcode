package sorting

import "testing"

func TestCountingSort(t *testing.T) {
	arr := []int{5, 3, 4, 1, 2, 7, 6, 9, 6, 4, 8, 9, 8, 0}
	CountingSort(arr)
}

func TestSumCountingSort(t *testing.T) {
	arr := []int{5, 3, 4, 1, 2, 7, 6, 9, 6, 4, 8, 9, 8, 0}
	SumCountingSort(arr)
}
