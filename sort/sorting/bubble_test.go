package sorting

import (
	"log"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{5, 3, 4, 1, 2, 0, 6, 7, 9, 8}
	BubbleSort(arr)
	log.Printf("BubbleSort result= %v\n", arr)
}

func TestBubbleSort2(t *testing.T) {
	arr := []int{5, 3, 4, 1, 2, 0, 6, 7, 9, 8}
	BubbleSort2(arr)
	log.Printf("BubbleSort2 result= %v\n", arr)
}

func TestBubbleSort3(t *testing.T) {
	arr := []int{5, 3, 4, 1, 2, 0, 6, 7, 9, 8}
	BubbleSort3(arr)
	log.Printf("BubbleSort3 result= %v\n", arr)
}
