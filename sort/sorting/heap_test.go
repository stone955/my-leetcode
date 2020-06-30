package sorting

import (
	"log"
	"testing"
)

func TestHeapSort(t *testing.T) {
	arr := []int{5, 3, 4, 1, 2, 0, 6, 7, 9, 8}
	HeapSort(arr)
	log.Printf("HeapSort result= %v\n", arr)
}
