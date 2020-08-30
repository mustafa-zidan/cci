// Palindrome Permutation: Given a string, write a function to check if it is a
// premutation of a palindrome. A palindrome is a word of a phrase that is the
// same forwards and backwards. A premutation is a rearrangement of the letters.
// the palindrome does not need to be limited to just dictionary words.

// Example:
// Input: "Tact Coa"
// Output: True (permutations: "taco cat", "atco cta", etc.)

package arrays

import (
	"unicode"
)

func palindromePermutationAlternativeImplementation(s string) bool {
	if len(s) == 0 {
		return true
	}
	odd := 0
	dict := make(map[rune]int)
	for _, c := range s {
		if !unicode.IsLetter(c) {
			continue
		}
		c = unicode.ToLower(c)
		dict[c]++
		if dict[c]%2 == 0 {
			odd--
		} else {
			odd++
		}
	}
	return odd < 2
}

// this one is very smart and I have no idea how I came up with this :|
func palindromePermutation(s string) bool {
	alphabit := 0
	for _, c := range s {
		if !unicode.IsLetter(c) {
			continue
		}
		c = unicode.ToLower(c)
		toggle(alphabit, c)
	}
	return alphabit == 0 || hasExactlyOneBit(alphabit)
}

func hasExactlyOneBit(n int) bool {
	return (n&(n-1) == 0)
}

func toggle(n int, c rune) {
	mask := 1 << uint(c)
	if n&mask == 0 { // new char
		n |= mask
	} else {
		n ^= mask
	}
}
