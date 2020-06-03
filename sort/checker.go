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
	}
	return "Unknown"
}

func CheckSort(typ Type, num int) {
	in := make([]int, num)
	checkIn := make([]int, num)
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)
	for i := 0; i < num; i++ {
		r := r.Intn(num)
		in[i] = r
		checkIn[i] = r
	}

	switch typ {
	case Selection:
		SelectionSort(in)
	case Bubble:
		BubbleSort(in)
	case Insertion:
		InsertionSort(in)
	case Shell:
		ShellSort(in)
	case Shell2:
		ShellSort2(in)
	case Merge:
		MergeSort(in)
	case SinglePivotQuick:
		SinglePivotQuickSort(in)
	case DualPivotQuick:
		DualPivotQuickSort(in)
	}

	sort.Ints(checkIn)

	if checked := check(in, checkIn); checked {
		fmt.Printf("%v CheckSort success\n", typ.String())
	} else {
		fmt.Printf("%v CheckSort error\n", typ.String())
	}
}

func check(in []int, checkIn []int) bool {
	for i := 0; i < len(in); i++ {
		if in[i] != checkIn[i] {
			return false
		}
	}
	return true
}
