package sorting

import (
	"log"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	arr := []int{5, 3, 4, 1, 2, 7, 6, 9, 8, 0}
	InsertionSort(arr)
	log.Printf("InsertionSort result= %v\n", arr)
}

func TestInsertionSort2(t *testing.T) {
	arr := []int{5, 3, 4, 1, 2, 7, 6, 9, 8, 0}
	InsertionSort2(arr)
	log.Printf("InsertionSort2 result= %v\n", arr)
}
