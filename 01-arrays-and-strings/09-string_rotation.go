package arrays

//Assume you have a method isSubstring which checks if one word is a substring
//of another. Given two strings, s1 and s2, write a code to check if s2 is
//a rotation of s1 using only one call to is SubString (e.g, "waterbottle" is
//a rotation of "erbottlewat")

func stringRotation(first, second string) bool {
	// no need to run the matcher
	if len(first) != len(second) {
		return false
	}
	// means no rotation in the two strings
	if first == second {
		return true
	}
	firstRunes, secondRunes := []rune(first), []rune(second)
	i, j := 0, 0

	for i < len(first) && j < len(second) {
		if firstRunes[i] == secondRunes[j] {
			i++
			j++
		} else {
			j++
		}
	}

	rotated := first[i:] + first[:i]
	return rotated == second
}
