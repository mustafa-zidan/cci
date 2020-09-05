package arrays

// Implement an algorithm to determine if a string has all unique characters.
// what if you cannot use additional data structure?

// build an array of 128 bool where each element will hold the info if
// a rune appeared in the string or not. we check if the element with the index
// or the rune is true then it means it returns false otherwise it continue
// walking through the string
// Space Complexity O(128), Runtime Complexity O(128) - ASCII LIMIT
func isUniqueAlternativeImplementation(str string) bool {
	length := 128
	if len(str) > length {
		return false
	} // ascii limit
	check := make([]bool, length)
	for _, char := range str {
		c := uint(char)
		if check[c] {
			return false
		}
		check[c] = true
	}
	return true
}

// checker is acting like a map of bits starts with
// 00000...000 then when it is anded with anything it
// result zero and then we push one in it's unique place
// in the bit map so if there "abcfgh" will translate to
// 111011100... when this value is anded with any value
// that has a unique letter e.g, "d" it results 0 which
// means this number is unique
// Space Complexity O(1), Runtime Complexity O(128) - ASCII limit
func isUnique(str string) bool {
	if len(str) > 128 {
		return false
	} // ascii limit
	var hash uint // int is 64 and uint holds 128 bit depending on the cpu architecture
	for _, char := range str {
		c := uint(char) - uint('a')
		if hash&(1<<c) > 0 {
			return false
		}
		hash |= 1 << c
	}
	return true
}
