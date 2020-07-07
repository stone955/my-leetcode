package sorting

// InsertionSort 插入排序
func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				swap(arr, j-1, j)
			}
		}
	}
}

// InsertionSort2 插入排序优化
// 交换改成挪动
func InsertionSort2(arr []int) {
	for i := 1; i < len(arr); i++ {
		index := i
		for j := i; j > 0; j-- {
			if arr[i] < arr[j-1] {
				index = j - 1
			}
		}
		if index != i {
			tmp := arr[i] // 待插入的元素
			// 挪动index ~ i元素
			for k := i - 1; k >= index; k-- {
				arr[k+1] = arr[k]
			}
			arr[index] = tmp
		}
	}
}
