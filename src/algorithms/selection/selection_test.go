package selection

import (
	"sort/src/comparator"
	"sort/src/datasets"
	"sort/src/testutil"
	"testing"
)

func TestSortEmpty(t *testing.T) {
	testutil.TestEquality(Sort(datasets.Empty(), comparator.Asc), datasets.Empty(), t)
	testutil.TestEquality(Sort(datasets.Empty(), comparator.Desc), datasets.Empty(), t)
}

func TestSortOneElement(t *testing.T) {
	testutil.TestEquality(Sort(datasets.OneElement(), comparator.Asc), datasets.OneElement(), t)
	testutil.TestEquality(Sort(datasets.OneElement(), comparator.Desc), datasets.OneElement(), t)
}

func TestSortTwoEqualElements(t *testing.T) {
	testutil.TestEquality(Sort(datasets.TwoEqualElements(), comparator.Asc), datasets.TwoEqualElements(), t)
	testutil.TestEquality(Sort(datasets.TwoEqualElements(), comparator.Desc), datasets.TwoEqualElements(), t)
}

func TestSortTwoElements(t *testing.T) {
	testutil.TestEquality(Sort(datasets.TwoElementsAsc(), comparator.Asc), datasets.TwoElementsAsc(), t)
	testutil.TestEquality(Sort(datasets.TwoElementsDesc(), comparator.Asc), datasets.TwoElementsAsc(), t)
	testutil.TestEquality(Sort(datasets.TwoElementsDesc(), comparator.Desc), datasets.TwoElementsDesc(), t)
	testutil.TestEquality(Sort(datasets.TwoElementsAsc(), comparator.Desc), datasets.TwoElementsDesc(), t)
}

func TestSortEqualElements(t *testing.T) {
	testutil.TestEquality(Sort(datasets.EqualElements(), comparator.Asc), datasets.EqualElements(), t)
	testutil.TestEquality(Sort(datasets.EqualElements(), comparator.Asc), datasets.EqualElements(), t)
}

func TestSortUniqueElements(t *testing.T) {
	testutil.TestEquality(Sort(datasets.UnsortedUnique(), comparator.Asc), datasets.SortedUniqueAsc(), t)
	testutil.TestEquality(Sort(datasets.UnsortedUnique(), comparator.Desc), datasets.SortedUniqueDesc(), t)
}

func TestSortNotUniqueElements(t *testing.T) {
	testutil.TestEquality(Sort(datasets.Unsorted(), comparator.Asc), datasets.SortedAsc(), t)
	testutil.TestEquality(Sort(datasets.Unsorted(), comparator.Desc), datasets.SortedDesc(), t)
}