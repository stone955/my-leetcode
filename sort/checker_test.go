package sort

import "testing"

func TestCheckSort(t *testing.T) {
	CheckSort(Selection, 10000)
	CheckSort(Bubble, 10000)
	CheckSort(Insertion, 10000)
}
