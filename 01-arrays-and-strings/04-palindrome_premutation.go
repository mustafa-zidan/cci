package main

import(
	"fmt"
	"unicode"
)

func palindromePremutation(s string) bool {
	if len(s) <= 0 {
		return true
	}
	odd := 0
	dict := make(map[rune]int)
	for _, c := range s {
		if ! unicode.IsLetter(c)  {
			continue
		}
		c = unicode.ToLower(c)
		dict[c]++;
		if dict[c] %2 == 0 {
			odd--
		} else {
			odd++
		}
	}
	return odd < 2
}

func main() {
	fmt.Println("Tact aCo", palindromePremutation("Tact aCo"))
}