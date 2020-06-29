package main

import "github.com/stone955/my-leetcode/sort/sorting"

func main() {
	num, n := 10000, 1000
	sorting.CheckSort(sorting.Bubble, num, n)
	sorting.CheckSort(sorting.Bubble2, num, n)
	sorting.CheckSort(sorting.Bubble3, num, n)
	sorting.CheckSort(sorting.Selection, num, n)
	sorting.CheckSort(sorting.Insertion, num, n)
	sorting.CheckSort(sorting.Shell, num, n)
	sorting.CheckSort(sorting.Shell2, num, n)
	sorting.CheckSort(sorting.Merge, num, n)
	sorting.CheckSort(sorting.SinglePivotQuick, num, n)
	sorting.CheckSort(sorting.DualPivotQuick, num, n)
	sorting.CheckSort(sorting.Counting, num, n)
	sorting.CheckSort(sorting.SumCounting, num, n)
}
