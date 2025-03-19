package util

import (
	"testing"
)

// TestSliceContains tests the SliceContains function
func TestSliceContains(t *testing.T) {
	tests := []struct {
		slice    []string
		str      string
		contains bool
	}{
		{[]string{"a", "b", "c"}, "a", true},
		{[]string{"a", "b", "c"}, "d", false},
		{[]string{}, "a", false},
	}

	for _, test := range tests {
		result := SliceContains(test.slice, test.str)
		if result != test.contains {
			t.Errorf("SliceContains(%v, %s) = %v; want %v", test.slice, test.str, result, test.contains)
		}
	}
}

// TestSliceIndex - tests the SliceIndex function
func TestSliceIndex(t *testing.T) {
	tests := []struct {
		slice []string
		str   string
		index int
	}{
		{[]string{"a", "b", "c"}, "a", 0},
		{[]string{"a", "b", "c"}, "d", -1},
		{[]string{}, "a", -1},
	}

	for _, test := range tests {
		result := SliceIndex(test.slice, test.str)
		if result != test.index {
			t.Errorf("SliceIndex(%v, %s) = %d; want %d", test.slice, test.str, result, test.index)
		}
	}
}

// TestSliceRemoveElement - tests the SliceRemoveElement function
func TestSliceRemoveElement(t *testing.T) {
	tests := []struct {
		slice    []string
		index    int
		expected []string
	}{
		{[]string{"a", "b", "c"}, 1, []string{"a", "c"}},
		{[]string{"a", "b", "c"}, 3, []string{"a", "b", "c"}},
		{[]string{"a", "b", "c"}, -1, []string{"a", "b", "c"}},
		{[]string{"a", "b", "c"}, 0, []string{"b", "c"}},
		{[]string{"a"}, 0, []string{}},
	}

	for _, test := range tests {
		result := SliceRemoveElement(test.slice, test.index)
		if !equalSlices(result, test.expected) {
			t.Errorf("SliceRemoveElement(%v, %d) = %v; want %v", test.slice, test.index, result, test.expected)
		}
	}
}

// TestSliceGetRandomIndex - tests the SliceGetRandomIndex function
func TestSliceGetRandomIndex(t *testing.T) {
	slice := []string{"a", "b", "c"}
	index := SliceGetRandomIndex(slice)
	if index < 0 || index >= len(slice) {
		t.Errorf("SliceGetRandomIndex(%v) = %d; want index in range [0, %d)", slice, index, len(slice))
	}

	// Testing empty slice edge case
	var emptySlice []string
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("SliceGetRandomIndex did not panic on empty slice")
		}
	}()
	SliceGetRandomIndex(emptySlice)
}

// Helper function to compare two slices for equality
func equalSlices(a, b []string) bool {
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
