package sort

import (
	"fmt"
	"time"
)

func MergeSort(arr []int) {
	// TODO
	mergeSort(arr)
}

func mergeSort(arr []int) {
	merge(arr)
}

func merge(arr []int) {
	begin := time.Now()

	mid := len(arr) / 2
	tempArr := make([]int, len(arr))

	i := 0       // 前半段起始位置
	j := mid + 1 // 后半段起始位置
	k := 0       // 新数组起始位置

	for i <= mid && j < len(arr) { // 前后两段数组都没遍历完，则进入循环
		if arr[i] <= arr[j] {
			tempArr[k] = arr[i]
			k++
			i++
		} else {
			tempArr[k] = arr[j]
			k++
			j++
		}
	}

	for i <= mid {
		tempArr[k] = arr[i]
		k++
		i++
	}

	for j < len(arr) {
		tempArr[k] = arr[j]
		k++
		j++
	}

	fmt.Printf("MergeSort cost %v, arr %v\n", time.Now().Sub(begin).Milliseconds(), tempArr)
}
