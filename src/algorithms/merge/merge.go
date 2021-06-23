package merge

import (
	"sort/src/comparator"
)

func Sort(inputSlice []int, c comparator.Comparator) []int {
	if len(inputSlice) > 1 {
		a := Sort(inputSlice[:len(inputSlice)/2], c)
		b := Sort(inputSlice[len(inputSlice)/2:], c)
		return merge(a, b, c)
	}

	return inputSlice
}

func merge(a, b []int, c comparator.Comparator) []int {
	i, j := 0, 0
	merged := make([]int, len(a) + len(b))
	for k := 0; k < len(merged); k++ {
		if i == len(a) {
			merged[k] = b[j]
			j++
		} else if j == len(b) {
			merged[k] = a[i]
			i++
		} else if c(a[i], b[j]) >= 0 {
			merged[k] = a[i]
			i++
		} else {
			merged[k] = b[j]
			j++
		}
	}

	return merged
}