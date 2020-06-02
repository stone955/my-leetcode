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
	}

	sort.Ints(checkIn)

	for i := 0; i < len(in); i++ {
		if in[i] != checkIn[i] {
			fmt.Printf("%v CheckSort error\n", typ.String())
			break
		}
	}
	fmt.Printf("%v CheckSort success\n", typ.String())
}
