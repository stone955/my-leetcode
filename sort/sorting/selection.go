package sorting

// SelectionSort 选择排序
func SelectionSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		// 第i次遍历, 选出最小的与第i个元素交换位置
		minPos := i
		temp := arr[i]
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < temp {
				minPos = j
				temp = arr[minPos]
			}
		}
		if arr[minPos] < arr[i] {
			temp = arr[i]
			arr[i] = arr[minPos]
			arr[minPos] = temp
		}
	}
}
