package testutil

import (
	"testing"
)

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

func TestEquality(actual, expected []int, t *testing.T) {
	if !areSlicesEqual(actual, expected) {
		t.Errorf("Slices not equal! Expected %v but got %v", expected, actual)
	}
}