package sorting

// getMinMax 计算数组的边界
func getMinMax(arr []int) (int, int) {
	var min, max int
	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}
	return min, max
}

// CountingSort 计数排序，不稳定的排序
func CountingSort(arr []int) {
	// 先计算边界
	min, max := getMinMax(arr)
	// 创建计数数组
	counting := make([]int, max-min+1)

	// 遍历待排序数据计数
	for i := 0; i < len(arr); i++ {
		counting[arr[i]-min]++
	}

	//fmt.Printf("Counting arr: %v\n", counting)

	// 遍历计数数组进行排序
	idx := 0
	for i := 0; i < len(counting); i++ {
		for j := 0; j < counting[i]; j++ {
			arr[idx] = i + min
			idx++
		}
	}
}

// SumCountingSort 累加计数排序，稳定的排序
func SumCountingSort(arr []int) {
	// 先计算边界
	min, max := getMinMax(arr)
	// 创建计数数组
	counting := make([]int, max-min+1)

	// 遍历待排序数据计数
	for i := 0; i < len(arr); i++ {
		counting[arr[i]-min]++
	}

	// 如果只按上面的方法排序，排序是不稳定的
	// 增加下面的处理能够使排序稳定，即将计数数组变成累加数组，好处是每个累加数组的值对应的就是这个元素的位置。
	for i := 1; i < len(counting); i++ {
		counting[i] = counting[i] + counting[i-1]
	}

	//fmt.Printf("Counting arr: %v\n", counting)

	// 计算完累计数组，对原数组从后向前遍历一遍
	out := make([]int, len(arr))
	for i := len(arr); i > 0; i-- {
		idx := counting[arr[i-1]] - 1
		out[idx] = arr[i-1]
		counting[arr[i-1]]--
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = out[i]
	}
}
