package sorting

// ShellSort 希尔排序-固定间隔
func ShellSort(arr []int) {
	// 固定间隔
	for gap := 4; gap > 0; gap /= 2 {
		for i := 0; i < len(arr)-gap; i += gap {
			for j := i + gap; j > 0; j -= gap {
				if arr[j-gap] > arr[j] {
					// 交换
					temp := arr[j]
					arr[j] = arr[j-gap]
					arr[j-gap] = temp
				}
			}
		}
	}
}

// ShellSort2 希尔排序-根据数组长度计算间隔
func ShellSort2(arr []int) {
	// 计算gap
	h := 1
	for h < len(arr) {
		h = 3*h + 1
	}

	for gap := h; gap > 0; gap = (gap - 1) / 3 {
		for i := 0; i < len(arr)-gap; i += gap {
			for j := i + gap; j > 0; j -= gap {
				if arr[j-gap] > arr[j] {
					// 交换
					temp := arr[j]
					arr[j] = arr[j-gap]
					arr[j-gap] = temp
				}
			}
		}
	}
}
