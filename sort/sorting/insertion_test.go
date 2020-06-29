package sorting

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	in := []int{5, 3, 4, 1, 2, 7, 6, 9, 8, 0}
	InsertionSort(in)
}
