package sort

import (
	"fmt"
	"time"
)

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	begin := time.Now()

	for i := len(arr) - 1; i > 0; i-- {
		var swapped bool
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	fmt.Printf("BubbleSort cost %v, arr %v\n", time.Now().Sub(begin).Milliseconds(), arr)
}
