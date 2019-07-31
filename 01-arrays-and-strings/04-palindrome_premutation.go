package main

import(
	"fmt"
	"unicode"
)

// func palindromePremutation(s string) bool {
// 	if len(s) <= 0 {
// 		return true
// 	}
// 	odd := 0
// 	dict := make(map[rune]int)
// 	for _, c := range s {
// 		if ! unicode.IsLetter(c)  {
// 			continue
// 		}
// 		c = unicode.ToLower(c)
// 		dict[c]++;
// 		if dict[c] %2 == 0 {
// 			odd--
// 		} else {
// 			odd++
// 		}
// 	}
// 	return odd < 2
// }

func palindromePremutation(s string) bool {
	alphabit := 0
	for _, c := range s {
		if ! unicode.IsLetter(c)  {
			continue
		}
		c = unicode.ToLower(c)
		toggle(alphabit, c)
	}
	return alphabit == 0 || hasExactlyOneBit(alphabit)
}

func hasExactlyOneBit(n int) bool {
	return (n & (n - 1) == 0)
}

func toggle(n int , c rune) {
	mask := 1 << uint(c)
	if n & mask == 0 { // new char
		n |= mask
	} else {
		n ^= mask
	}
}

func main() {
	fmt.Println("Tact aCo", palindromePremutation("Tact aCo"))
}