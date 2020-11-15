package linkedlists

import (
	"unicode"
)

//Palindrome: Implement a function to check if a linked list is a palindrome.

func palindrome(l *LinkedList) bool {
	var alphabit uint64
	for l != nil {
		c := unicode.ToLower(l.Val.(rune))
		if unicode.IsLetter(c) {
			alphabit = toggle(alphabit, c)
		}
		l = l.Next
	}
	return alphabit == 0 || hasExactlyOneBit(alphabit)
}

func toggle(n uint64, c rune) uint64 {
	mask := uint64(1) << uint(c-'A')
	if n&mask == 0 { // bit does not exists
		n |= mask // turn the corresponding bit to 1
	} else {
		n ^= mask // turn the corresponding bit to 0
	}
	return n
}

func hasExactlyOneBit(n uint64) bool {
	return (n&(n-1) == 0)
}
