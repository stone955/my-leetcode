package sorting

import (
	"log"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	arr := []int{4, 7, 6, 3, 8, 1, 2, 5, 9, 0}
	SelectionSort(arr)
	log.Printf("SelectionSort result= %v\n", arr)
}

func TestSelectionSort2(t *testing.T) {
	arr := []int{4, 7, 6, 3, 8, 1, 2, 5, 9, 0}
	SelectionSort2(arr)
	log.Printf("SelectionSort2 result= %v\n", arr)
}
