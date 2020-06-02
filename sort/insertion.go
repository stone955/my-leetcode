package sort

import "fmt"

// InsertionSort 插入排序
func InsertionSort(in []int) {
	for i := 0; i < len(in)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if in[j-1] > in[j] {
				// 交换
				temp := in[j]
				in[j] = in[j-1]
				in[j-1] = temp
			}
		}
	}
	fmt.Printf("InsertionSort() %v\n", in)
}
