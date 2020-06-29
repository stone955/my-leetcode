package sorting

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	Bubble           = "BubbleSort"
	Bubble2          = "BubbleSort2"
	Bubble3          = "BubbleSort3"
	Selection        = "SelectionSort"
	Insertion        = "InsertionSort"
	Shell            = "ShellSort"
	Shell2           = "Shell2Sort"
	Merge            = "MergeSort"
	SinglePivotQuick = "SinglePivotQuickSort"
	DualPivotQuick   = "DualPivotQuickSort"
	Counting         = "CountingSort"
	SumCounting      = "SumCountingSort"
)

type Type string

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
	case Insertion:
		InsertionSort(arr)
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
