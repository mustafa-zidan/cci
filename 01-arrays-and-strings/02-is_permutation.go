// check permutation: Given two strings, write a method to decide if one is a
// permutation of the other.
package arrays

import (
	"sort"
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

// sort and compare
func isPermutationAlternativeImplementation(src, target string) bool {
	if len(src) != len(target) {
		return false
	}
	return Sort(src) == Sort(target)
}

// Counter
func isPermutation(src, target string) bool {
	if len(src) != len(target) {
		return false
	}
	m := make([]rune, 128)
	for _, c := range src {
		m[c]++
	}

	for _, c := range target {
		if m[c] == 0 {
			return false
		}
		m[c]--
	}
	for _, count := range m {
		if count != 0 {
			return false
		}
	}
	return true
}
