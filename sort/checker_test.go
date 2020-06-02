package sort

import "testing"

func TestCheckSort(t *testing.T) {
	CheckSort(Selection, 100000)
	CheckSort(Bubble, 100000)
	CheckSort(Insertion, 100000)
	CheckSort(Shell, 100000)
	CheckSort(Shell2, 100000)
}
