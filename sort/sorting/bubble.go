package sorting

// BubbleSort 冒泡排序
func BubbleSort(arr []int) {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			}
		}
	}
}

// BubbleSort2 冒泡排序优化
// 如果未发生过交换，则退出循环
func BubbleSort2(arr []int) {
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
}

// BubbleSort3 冒泡排序优化
// 记录上次交换位置，减少遍历元素个数
func BubbleSort3(arr []int) {
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
}
