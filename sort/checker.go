package sort

import (
	"fmt"
	"math/rand"
	"sort"
)

type Type int

const (
	Selection = iota
	Bubble
)

func CheckSort(typ Type, num int) {
	in := make([]int, num)
	checkIn := make([]int, num)
	for i := 0; i < num; i++ {
		r := rand.Intn(num)
		in[i] = r
		checkIn[i] = r
	}

	switch typ {
	case Selection:
		SelectionSort(in)
	case Bubble:
		BubbleSort(in)
	}

	sort.Ints(checkIn)

	for i := 0; i < len(in); i++ {
		if in[i] != checkIn[i] {
			fmt.Println("CheckSort error")
			break
		}
	}
	fmt.Println("CheckSort success")
}
