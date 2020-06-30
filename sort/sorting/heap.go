package sorting

// HeapSort 堆排序
func HeapSort(arr []int) {
	heapSort(arr, 0, len(arr)-1)
}

func heapSort(arr []int, begin int, end int) {
	// 如果 begin == end 排序结束
	if begin == end {
		return
	}
	// 对数组的 [begin, end] 元素进行堆排序
	// 1.建大顶堆 TODO

	// 2.堆顶元素与末尾元素交换位置
	swap(arr, begin, end)
	// 3.剩余元素再递归调用堆排序函数
	heapSort(arr, begin, end-1)
}
