package datasets

func Empty() []int {
	return []int {}
}

func OneElement() []int {
	return []int {1}
}
func TwoElementsAsc() []int {
	return []int{1, 2}
}
func TwoElementsDesc() []int {
	return []int{2, 1}
}
func TwoEqualElements() []int {
	return []int{1, 1}
}
func EqualElements() []int {
	return []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
}
func Unsorted() []int {
	return []int{5, 8, 4, 13, 2, 26, 27, 25, 8, 2}
}
func UnsortedUnique() []int {
	return []int{5, 8, 4, 13, 1, 26, 27, 25, 7, 2}
}
func SortedAsc() []int {
	return []int{2, 2, 4, 5, 8, 8, 13, 25, 26, 27}
}
func SortedDesc() []int {
	return []int{27, 26, 25, 13, 8, 8, 5, 4, 2, 2}
}
func SortedUniqueAsc() []int {
	return []int{1, 2, 4, 5, 7, 8, 13, 25, 26, 27}
}
func SortedUniqueDesc() []int {
	return []int{27, 26, 25, 13, 8, 7, 5, 4, 2, 1}
}