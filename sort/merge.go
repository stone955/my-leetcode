package sort

import (
	"fmt"
	"time"
)

func MergeSort(arr []int) {
	begin := time.Now()
	sortedArr := mergeSort(arr)
	// 赋值
	for i := 0; i < len(arr); i++ {
		arr[i] = sortedArr[i]
	}
	fmt.Printf("MergeSort cost %v, arr %v\n", time.Now().Sub(begin).Milliseconds(), sortedArr)
}

func mergeSort(arr []int) []int {
	// 如果元素不能再分，则直接返回
	if len(arr) == 1 {
		return arr
	}
	// 分解
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {

	mergedArr := make([]int, len(left)+len(right))

	i := 0 // 前半段起始位置
	j := 0 // 后半段起始位置
	k := 0 // 新数组起始位置

	for i < len(left) && j < len(right) { // 前后两段数组都没遍历完，则进入循环
		if left[i] <= right[j] {
			mergedArr[k] = left[i]
			k++
			i++
		} else {
			mergedArr[k] = right[j]
			k++
			j++
		}
	}

	for i < len(left) {
		mergedArr[k] = left[i]
		k++
		i++
	}

	for j < len(right) {
		mergedArr[k] = right[j]
		k++
		j++
	}
	return mergedArr
}
