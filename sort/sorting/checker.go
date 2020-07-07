package sorting

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	// Bubble 冒泡排序
	Bubble = "BubbleSort"
	// Bubble2 冒泡排序优化
	Bubble2 = "BubbleSort2"
	// Bubble3 冒泡排序优化2
	Bubble3 = "BubbleSort3"
	// Selection 选择排序
	Selection = "SelectionSort"
	// Heap 堆排序
	Heap = "HeapSort"
	// Insertion 插入排序
	Insertion = "InsertionSort"
	// Insertion2 插入排序优化，交换改成挪动
	Insertion2 = "InsertionSort2"
	// Shell 希尔排序
	Shell = "ShellSort"
	// Shell2 希尔排序优化
	Shell2 = "Shell2Sort"
	// Merge 归并排序
	Merge = "MergeSort"
	// SinglePivotQuick 单轴快排
	SinglePivotQuick = "SinglePivotQuickSort"
	// DualPivotQuick 双轴快排
	DualPivotQuick = "DualPivotQuickSort"
	// Counting 计数排序
	Counting = "CountingSort"
	// SumCounting 合并计数排序
	SumCounting = "SumCountingSort"
)

// Type 排序类型
type Type string

// CheckSort 验证排序结果
func CheckSort(typ Type, num, n int) {
	arr := make([]int, num)
	checkArr := make([]int, num)
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	for i := 0; i < num; i++ {
		r := r.Intn(n)
		arr[i] = r
		checkArr[i] = r
	}

	begin := time.Now()

	switch typ {
	case Bubble:
		BubbleSort(arr)
	case Bubble2:
		BubbleSort2(arr)
	case Bubble3:
		BubbleSort3(arr)
	case Selection:
		SelectionSort(arr)
	case Heap:
		HeapSort(arr)
	case Insertion:
		InsertionSort(arr)
	case Insertion2:
		InsertionSort2(arr)
	case Shell:
		ShellSort(arr)
	case Shell2:
		ShellSort2(arr)
	case Merge:
		MergeSort(arr)
	case SinglePivotQuick:
		SinglePivotQuickSort(arr)
	case DualPivotQuick:
		DualPivotQuickSort(arr)
	case Counting:
		CountingSort(arr)
	case SumCounting:
		SumCountingSort(arr)
	}

	costs := time.Now().Sub(begin).Milliseconds()

	sort.Ints(checkArr)

	if checked := check(arr, checkArr); checked {
		fmt.Printf("CheckSort success type= %v, costs= %v ms\n", typ, costs)
	} else {
		fmt.Printf("CheckSort error type= %v\n", typ)
	}
}

func check(arr []int, checkArr []int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] != checkArr[i] {
			return false
		}
	}
	return true
}
