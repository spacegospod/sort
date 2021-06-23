package bubble

import (
	"sort/src/comparator"
)

func Sort(inputSlice []int, c comparator.Comparator) []int {
	for i := 0; i < len(inputSlice) - 1; i++ {
		sorted := true
		for j := 0; j < len(inputSlice) - 1 - i; j++ {
			if !c(inputSlice[j], inputSlice[j + 1]) {
				inputSlice[j + 1], inputSlice[j] = inputSlice[j], inputSlice[j + 1]
				sorted = false
			}
		}
		if sorted {
			break
		}
	}

	return inputSlice
}