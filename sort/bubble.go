package sort

import "fmt"

// BubbleSort 冒泡排序
func BubbleSort(in []int) {
	for i := len(in) - 1; i > 0; i-- {
		var swapped bool
		for j := 0; j < i; j++ {
			if in[j] > in[j+1] {
				temp := in[j+1]
				in[j+1] = in[j]
				in[j] = temp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	fmt.Printf("BubbleSort() %v\n", in)
}
