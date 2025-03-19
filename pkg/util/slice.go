package util

import "math/rand"

// SliceIndex - returns the position of the first occurrence of v in slice s,
// or -1 if v is not present in s.
// Suitable for any type that supports the == operator.
func SliceIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if v == s[i] {
			return i
		}
	}
	return -1
}

// SliceContains - returns true if slice contains the givens string
// Suitable for any type that supports the == operator.
func SliceContains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

// SliceRemoveElement -safely removes an element from a slice, if it's in bounds
func SliceRemoveElement(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

// SliceGetRandomIndex returns a random index for a slice
func SliceGetRandomIndex(slice []string) int {
	return rand.Intn(len(slice))
}
