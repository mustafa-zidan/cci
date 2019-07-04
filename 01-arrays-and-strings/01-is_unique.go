package main

import "fmt"

//
// func isUnique(str string) bool {
// 	length := len(str)
// 	if length > 128 { return false } // ascii limit
// 	if length == 0 || length == 1 { return true }

// 	check := make(map[rune]bool, length)
// 	for _,char := range str {
// 		if v, ok := check[char]; ok && v {
// 			return false
// 		}
// 		check[char] = true
// 	}
// 	return true
// }

// checker is acting like a map of bits starts with
// 00000...000 then when it is anded with anything it
// result zero and then we push one in it's unique place
// in the bit map so if there abcfgh will translate to
// 111011100... when this value is anded with any value
// that has a unique letter e.g,
func isUnique(str string) bool {
	if len(str) > 128 { return false } // ascii limit
	var checker uint
	for _, char := range str {
		c := uint(char) - uint('a')
		if checker & (1 << c ) > 0 {
			return false
		}
		checker |= 1 << c
	}
	return true
}


func main() {
	fmt.Println("isUnique(\"\"):\t", isUnique(""))
	fmt.Println("isUnique(\"a\"):\t", isUnique("a"))
	fmt.Println("isUnique(\"abcd\"):\t", isUnique("abcd"))
	fmt.Println("isUnique(\"Hello\"):\t", isUnique("Hello"))
}