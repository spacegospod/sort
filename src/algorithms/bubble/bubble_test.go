package bubble

import (
	"sort/src/testutil"
	"testing"
)

func TestSortEmpty(t *testing.T) {
	testutil.TestSortEmpty(Sort, t)
}

func TestSortOneElement(t *testing.T) {
	testutil.TestSortOneElement(Sort, t)
}

func TestSortTwoEqualElements(t *testing.T) {
	testutil.TestSortTwoEqualElements(Sort, t)
}

func TestSortTwoElements(t *testing.T) {
	testutil.TestSortTwoElements(Sort, t)
}

func TestSortEqualElements(t *testing.T) {
	testutil.TestSortEqualElements(Sort, t)
}

func TestSortUniqueElements(t *testing.T) {
	testutil.TestSortUniqueElements(Sort, t)
}

func TestSortNotUniqueElements(t *testing.T) {
	testutil.TestSortNotUniqueElements(Sort, t)
}