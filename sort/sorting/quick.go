package sorting

// SinglePivotQuickSort 单轴快排
func SinglePivotQuickSort(arr []int) {
	singlePivotQuickSort(arr, 0, len(arr)-1)
}

func singlePivotQuickSort(arr []int, leftBound, rightBound int) {
	if leftBound >= rightBound {
		return
	}
	// 轴的值
	pivot := arr[rightBound]
	// 定义指针
	left := leftBound
	right := rightBound - 1

	for left <= right {
		for left <= right && arr[left] <= pivot {
			left++
		}
		for left <= right && arr[right] > pivot {
			right--
		}

		// 交换
		if left < right {
			temp := arr[right]
			arr[right] = arr[left]
			arr[left] = temp
			left++
			right--
		}
	}

	// 将轴放在该放的位置
	temp := arr[rightBound]
	arr[rightBound] = arr[left]
	arr[left] = temp

	singlePivotQuickSort(arr, 0, left-1)
	singlePivotQuickSort(arr, left+1, rightBound)
}

// DualPivotQuickSort 双轴快排
func DualPivotQuickSort(arr []int) {
	dualPivotQuickSort(arr)
}

func dualPivotQuickSort(arr []int) {

}
