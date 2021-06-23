package comparator

// Comparator function
// Returns true if the values are in order or if they are equal
// Returns false if the values are not in order
type Comparator func(int, int) bool

func Asc(a, b int) bool {
	return a < b
}

func Desc(a, b int) bool {
	return a > b
}