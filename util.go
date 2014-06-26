package cram

import (
	"strconv"
)

func SlicesEqual(a, b []string) bool {
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

// Use linear search to return the index of target in slice of runes a
// Return -1 if not found.
func SearchRunes(a []rune, target rune) int {
	for i, r := range a {
		if r == target {
			return i
		}
	}
	return -1
}

// Use binary search to return the index of target in sorted slice of strings a
// Return -1 if not found.
func BSearchStrings(a []string, target string) int {
	s := 0
	e := len(a) - 1
	for s <= e {
		mid := (s + e) / 2
		if a[mid] == target {
			return mid
		}
		if a[mid] < target {
			s = mid + 1
		} else {
			e = mid - 1
		}
	}
	return -1
}

// Return whether a string is all digits
func IsDigits(a string) bool {
	if _, err := strconv.Atoi(a); err != nil {
		return false
	}
	return true
}
