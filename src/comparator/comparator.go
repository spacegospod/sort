package comparator

// Comparator function
// Returns a value greater than 0 if the parameters are in sorted order.
// Returns a value less than 0 if the parameters are not in sorted order.
// Returns 0 if the two parameters are equal
type Comparator func(int, int) int

func Asc(a, b int) int {
	return b - a
}

func Desc(a, b int) int {
	return a - b
}