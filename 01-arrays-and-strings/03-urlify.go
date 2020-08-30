// URLify: write a method to replace all spaces in a string with '%20'.
// You may assume that the string has sufficent space at the end to hold
// the additional charachtersm, and that you are given the true length
// of the string.

// Example
// Input "Mr John Smith    ", 13
// Output "Mr%20John%20Smith"
package arrays

func urlifyAlternativeImplementation(s string, trueLength int) string {
	count := 0
	for i := 0; i < trueLength; i++ {
		if s[i] == ' ' {
			count++
		}
	}
	r := []rune(s)
	// each space will be replaced with
	//%20 which have triple the count of spaces
	l := trueLength + (2 * count)
	for i := trueLength - 1; i > 0; i-- {
		if r[i] == ' ' {
			r[l-3], r[l-2], r[l-1] = '%', '2', '0'
			l -= 3
		} else {
			r[l-1] = r[i]
			l--
		}
	}
	return string(r)
}

func urlify(s string, trueLength int) string {
	end := len(s) - 1
	r := []rune(s)
	for i := trueLength - 1; i >= 0; i-- {
		if r[i] == ' ' {
			r[end-2], r[end-1], r[end] = '%', '2', '0'
			end -= 3
		} else {
			r[end] = r[i]
			end--
		}
	}
	return string(r)
}
