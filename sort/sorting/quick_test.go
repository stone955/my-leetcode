package sorting

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := []int{8, 3, 4, 1, 2, 5, 7, 9, 0, 6}
	SinglePivotQuickSort(arr)
}
