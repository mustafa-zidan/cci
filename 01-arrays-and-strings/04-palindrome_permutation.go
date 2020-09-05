package arrays

// Palindrome Permutation: Given a string, write a function to check if it is a
// premutation of a palindrome. A palindrome is a word of a phrase that is the
// same forwards and backwards. A premutation is a rearrangement of the letters.
// the palindrome does not need to be limited to just dictionary words.
//
// Example:
// Input: "Tact Coa"
// Output: True (permutations: "taco cat", "atco cta", etc.)

import (
	"unicode"
)

// Palindrome by defination has two conditions either all runes exist twice or
// all runes exist twice except 1
// We solve this by creating a map of rune and the number of it's occurence
// through the string, we also keep maintaining the number of odd runes appered
// in the string, if the total number of appearance is even then we decrement
// the odd runes count, else we increment int.
// then we simply check if the odd number of runes is either 0 or 1
// Space Complexity O(128), Runtime Complexity O(n)
func palindromePermutationAlternativeImplementation(s string) bool {
	if len(s) == 0 {
		return true
	}
	odd := 0
	m := make(map[rune]int)
	for _, c := range s {
		if !unicode.IsLetter(c) { // skip spaces and non-alphabitical characters
			continue
		}
		c = unicode.ToLower(c)
		m[c]++
		if m[c]%2 == 0 {
			odd--
		} else {
			odd++
		}
	}
	return odd == 0 || odd == 1
}

// this one is a smart solution and I have no idea how I came up with this :|
// here, an integer is used to keep toggling the bit corresponding to
// the character position in the alphabits by the following formula:
// if alphabit & 1 << rune == 0 it means the corresponding bit is zero and
// we need to switch it to one using alphabit = alphabit | 1 << rune
// else it means the character already exists and it needs to be switched to
// zero using alphabit = alphabit ^ 1 << rune
// and then check if it is zeroed out or has exactly one bit
// Space Complexity O(1), Runtime Complexity O(n)
func palindromePermutation(s string) bool {
	var alphabit uint64
	for _, c := range s {
		if !unicode.IsLetter(c) {
			continue
		}
		c = unicode.ToLower(c)
		toggle(alphabit, c)
	}
	return alphabit == 0 || hasExactlyOneBit(alphabit)
}

func hasExactlyOneBit(n uint64) bool {
	return (n&(n-1) == 0)
}

func toggle(n uint64, c rune) {
	mask := uint64(1) << int(c)
	if n&mask == 0 { // new char
		n |= mask
	} else {
		n ^= mask
	}
}
