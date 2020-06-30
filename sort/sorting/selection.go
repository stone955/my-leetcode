package sorting

// SelectionSort 选择排序
// 每次遍历选出一个最大的元素与最后一个元素交换位置
func SelectionSort(arr []int) {
	for i := len(arr); i > 0; i-- {
		maxPos := 0
		for j := 0; j < i; j++ {
			if arr[j] > arr[maxPos] {
				maxPos = j
			}
		}
		swap(arr, maxPos, i-1)
	}
}

// SelectionSort2 选择排序优化
// 每一次遍历同时选出最大元素和最小元素
func SelectionSort2(arr []int) {
	var begin, end = 0, len(arr)
	for begin < end {
		minPos := begin
		maxPos := begin
		for j := begin; j < end; j++ {
			if arr[j] > arr[maxPos] {
				maxPos = j
			}
			if arr[j] < arr[minPos] {
				minPos = j
			}
		}
		if maxPos == begin {
			maxPos = minPos
		}
		swap(arr, maxPos, end-1)
		if minPos == end-1 {
			minPos = maxPos
		}
		swap(arr, minPos, begin)
		begin++
		end--
	}
}
