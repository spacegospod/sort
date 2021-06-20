package testutil

import (
	"sort/src/comparator"
	"sort/src/datasets"
	"testing"
)

type Sort func([]int, comparator.Comparator) []int

func areSlicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func testEquality(actual, expected []int, t *testing.T) {
	if !areSlicesEqual(actual, expected) {
		t.Errorf("Slices not equal! Expected %v but got %v", expected, actual)
	}
}

func TestSortEmpty(s Sort, t *testing.T) {
	testEquality(s(datasets.Empty(), comparator.Asc), datasets.Empty(), t)
	testEquality(s(datasets.Empty(), comparator.Desc), datasets.Empty(), t)
}

func TestSortOneElement(s Sort, t *testing.T) {
	testEquality(s(datasets.OneElement(), comparator.Asc), datasets.OneElement(), t)
	testEquality(s(datasets.OneElement(), comparator.Desc), datasets.OneElement(), t)
}

func TestSortTwoEqualElements(s Sort, t *testing.T) {
	testEquality(s(datasets.TwoEqualElements(), comparator.Asc), datasets.TwoEqualElements(), t)
	testEquality(s(datasets.TwoEqualElements(), comparator.Desc), datasets.TwoEqualElements(), t)
}

func TestSortTwoElements(s Sort, t *testing.T) {
	testEquality(s(datasets.TwoElementsAsc(), comparator.Asc), datasets.TwoElementsAsc(), t)
	testEquality(s(datasets.TwoElementsDesc(), comparator.Asc), datasets.TwoElementsAsc(), t)
	testEquality(s(datasets.TwoElementsDesc(), comparator.Desc), datasets.TwoElementsDesc(), t)
	testEquality(s(datasets.TwoElementsAsc(), comparator.Desc), datasets.TwoElementsDesc(), t)
}

func TestSortEqualElements(s Sort, t *testing.T) {
	testEquality(s(datasets.EqualElements(), comparator.Asc), datasets.EqualElements(), t)
	testEquality(s(datasets.EqualElements(), comparator.Asc), datasets.EqualElements(), t)
}

func TestSortUniqueElements(s Sort, t *testing.T) {
	testEquality(s(datasets.UnsortedUnique(), comparator.Asc), datasets.SortedUniqueAsc(), t)
	testEquality(s(datasets.UnsortedUnique(), comparator.Desc), datasets.SortedUniqueDesc(), t)
}

func TestSortNotUniqueElements(s Sort, t *testing.T) {
	testEquality(s(datasets.Unsorted(), comparator.Asc), datasets.SortedAsc(), t)
	testEquality(s(datasets.Unsorted(), comparator.Desc), datasets.SortedDesc(), t)
}