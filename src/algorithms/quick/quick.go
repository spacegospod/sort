package quick

import (
	"sort/src/comparator"
)

func Sort(inputSlice []int, c comparator.Comparator) []int {
	if len(inputSlice) > 1 {
		i := partition(inputSlice, c)
		if i > 0 {
			Sort(inputSlice[:i], c)
		}
		if i < len(inputSlice) {
			Sort(inputSlice[i:], c)
		}
	}

	return inputSlice
}

func partition(inputSlice []int, c comparator.Comparator) int {
	pivot := inputSlice[0]
	i, j := 1, len(inputSlice) - 1

	for i < j {
		if !c(pivot, inputSlice[i]) {
			i++
		} else {
			inputSlice[i], inputSlice[j] = inputSlice[j], inputSlice[i]
			j--
		}
	}

	if !c(pivot, inputSlice[i]) {
		inputSlice[0], inputSlice[i] = inputSlice[i], inputSlice[0]
	}
	return i
}