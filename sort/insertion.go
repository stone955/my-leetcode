package sort

import (
	"fmt"
	"time"
)

// InsertionSort 插入排序
func InsertionSort(arr []int) {
	begin := time.Now()

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j-1] > arr[j] {
				// 交换
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}
	fmt.Printf("InsertionSort cost %v, arr %v\n", time.Now().Sub(begin).Milliseconds(), arr)
}
