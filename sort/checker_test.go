package sort

import "testing"

func TestCheckSort(t *testing.T) {
	CheckSort(Selection, 100)
	CheckSort(Bubble, 100)
	CheckSort(Insertion, 100)
	CheckSort(Shell, 100)
	CheckSort(Shell2, 100)
	CheckSort(Merge, 100)
}
