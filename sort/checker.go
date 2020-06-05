package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const (
	Selection = iota
	Bubble
	Insertion
	Shell
	Shell2
	Merge
	SinglePivotQuick
	DualPivotQuick
	Counting
	SumCounting
)

type Type int

func (t Type) String() string {
	switch t {
	case Selection:
		return "SelectionSort"
	case Bubble:
		return "BubbleSort"
	case Insertion:
		return "InsertionSort"
	case Shell:
		return "ShellSort"
	case Shell2:
		return "ShellSort2"
	case Merge:
		return "MergeSort"
	case SinglePivotQuick:
		return "SinglePivotQuick"
	case DualPivotQuick:
		return "DualPivotQuick"
	case Counting:
		return "Counting"
	case SumCounting:
		return "SumCounting"
	}
	return "Unknown"
}

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

	switch typ {
	case Selection:
		SelectionSort(arr)
	case Bubble:
		BubbleSort(arr)
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

	sort.Ints(checkArr)

	if checked := check(arr, checkArr); checked {
		fmt.Printf("%v CheckSort success\n", typ.String())
	} else {
		fmt.Printf("%v CheckSort error\n", typ.String())
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
