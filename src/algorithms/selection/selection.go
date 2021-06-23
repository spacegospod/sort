package selection

import (
	"sort/src/comparator"
)

func Sort(inputSlice []int, c comparator.Comparator) []int {
	for i := 0; i < len(inputSlice); i++ {
		for j := i; j < len(inputSlice); j++ {
			if !c(inputSlice[i], inputSlice[j]) {
				inputSlice[j], inputSlice[i] = inputSlice[i], inputSlice[j]
			}
		}
	}
	return inputSlice
}