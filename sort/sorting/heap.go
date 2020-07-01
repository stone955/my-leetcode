package sorting

// HeapSort 堆排序
func HeapSort(arr []int) {
	// 建大顶堆
	heapSort(arr, 0, len(arr)-1)
}

func heapSort(arr []int, begin int, end int) {
	// 如果 begin == end 建堆结束
	if begin == end {
		return
	}
	// 1.创建大顶堆
	buildMaxHeap(arr, begin, end)
	// 2.对剩余元素继续进行堆排序
	heapSort(arr, begin, end-1)
}

func buildMaxHeap(arr []int, begin int, end int) {
	// TODO
}
