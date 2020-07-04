package sorting

// HeapSort 堆排序
func HeapSort(arr []int) {
	// 建大顶堆
	for end := len(arr); end > 0; end-- {
		heapSort(arr, end)
	}
}

func heapSort(arr []int, end int) {
	// 1.创建大顶堆[0, end]
	buildMaxHeap(arr, end)
	// 2.交换位置
	swap(arr, 0, end-1)
}

// 自上而下建堆
func buildMaxHeap(arr []int, end int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		adjustHeap(arr, i, end)
	}
}

func adjustHeap(arr []int, curr int, end int) {
	left := 2*curr + 1  // 当前节点的左子节点
	right := 2*curr + 2 // 当前节点的右子节点
	largest := curr     // 当前最大节点
	// 选出最大节点与当前节点交换位置
	if left < end && arr[largest] < arr[left] {
		largest = left
	}
	if right < end && arr[largest] < arr[right] {
		largest = right
	}
	// 如果发生交换，则继续向下调整
	if largest != curr {
		swap(arr, curr, largest)
		adjustHeap(arr, largest, end)
	}
}
