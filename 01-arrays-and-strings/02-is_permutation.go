package arrays

// check permutation: Given two strings, write a method to decide if one is a
// permutation of the other.

import (
	"sort"
	"unicode"
)

type SortedRunes []rune

func (s SortedRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s SortedRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortedRunes) Len() int {
	return len(s)
}

func Sort(s string) string {
	r := []rune(s)
	sort.Sort(SortedRunes(r))
	return string(r)
}

// sort both strings then compare and compare
// Space Complexity O(1), Runtime Complexity O(n log(n))
// this approach is faulty since it will mismatch if the cases are
// different
func isPermutationAlternativeImplementation(src, target string) bool {
	if len(src) != len(target) {
		return false
	}
	return Sort(src) == Sort(target)
}

// Create an array that will hold the characters and the count of which
// each rune appears in both strings
// Space Complexity O(128), Runtime Complexity O(n)
func isPermutation(src, target string) bool {
	if len(src) != len(target) {
		return false
	}
	s, t := []rune(src), []rune(target)
	m := make([]rune, 128)

	for i := 0; i < len(s); i++ {
		m[unicode.ToLower(s[i])]++
		m[unicode.ToLower(t[i])]--
	}
	for _, count := range m {
		if count != 0 {
			return false
		}
	}
	return true
}
