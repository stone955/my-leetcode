package sorting

import "testing"

func TestCheckSort(t *testing.T) {
	num, n := 1000, 1000
	CheckSort(Bubble, num, n)
	CheckSort(Bubble2, num, n)
	CheckSort(Bubble3, num, n)
	CheckSort(Selection, num, n)
	CheckSort(Heap, num, n)
	CheckSort(Insertion, num, n)
	CheckSort(Shell, num, n)
	CheckSort(Shell2, num, n)
	CheckSort(Merge, num, n)
	CheckSort(SinglePivotQuick, num, n)
	CheckSort(DualPivotQuick, num, n)
	CheckSort(Counting, num, n)
	CheckSort(SumCounting, num, n)
}
