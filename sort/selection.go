package sort

import "fmt"

// SelectionSort 选择排序
func SelectionSort(in []int) {
	for i := 0; i < len(in)-1; i++ {
		// 第i次遍历, 选出最小的与第i个元素交换位置
		minPos := i
		temp := in[i]
		for j := i + 1; j < len(in); j++ {
			if in[j] < temp {
				minPos = j
				temp = in[minPos]
			}
		}
		if in[minPos] < in[i] {
			temp = in[i]
			in[i] = in[minPos]
			in[minPos] = temp
		}
	}
	fmt.Printf("SelectionSort() %v\n", in)
}
