package sorting

import (
	"testing"
)

func TestShellSort(t *testing.T) {
	in := []int{5, 3, 4, 1, 2, 7, 6, 9, 8, 0}
	ShellSort(in)
}

func TestShellSort2(t *testing.T) {
	in := []int{5, 3, 4, 1, 2, 7, 6, 9, 8, 0}
	ShellSort2(in)
}
