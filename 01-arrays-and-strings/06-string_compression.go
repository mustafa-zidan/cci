package arrays

import (
	"strconv"
	"strings"
)

// Implement a method to perform basic string compression using the counts of
// repeated characters, for example; the string aabcccccaaa would become a2b1c5a3.
// If the "compressed" string does not become smaller than the original string,
// your method should return the original string. You can assume the string has
// only upper case and lower case letters (a-z).

func compress(s string) string {
	var compressed strings.Builder
	runies := []rune(s)
	start, end := 0, 0
	for end < len(s) {
		if runies[start] == runies[end] {
			end++
		} else {
			compressed.WriteRune(runies[start])
			compressed.WriteString(strconv.Itoa(end - start))
			start = end
		}
		if compressed.Len() > len(s) {
			return s
		}
	}
	compressed.WriteRune(runies[start])
	compressed.WriteString(strconv.Itoa(end - start))
	return compressed.String()
}
