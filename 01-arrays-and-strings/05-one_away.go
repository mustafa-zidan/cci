package arrays

import "math"

// One Away: there are three types of edits that can be preformed on strings,
// insert a character, remove a character or replace a character. Given two
// strings, write a function to check if they are one edit (or zero edits) away
// example:
// pale,  ple -> true
// pales, pale -> true
// pale,  bale -> true
// pale,  bake -> true

func oneAway(src, target string) bool {
	if math.Abs(float64(len(src)-len(target))) > 1 {
		return false
	} // sanity check before we start
	for i := 0; i < len(src) && i < len(target); i++ {
		if src[i] != target[i] {
			return isRemovedCharacter(src, target, i) ||
				isReplacedCharacter(src, target, i) ||
				isAddedCharacter(src, target, i)
		}
	}
	return true
}

func isRemovedCharacter(src, target string, i int) bool {
	return i+1 < len(src) && i < len(target) && src[i+1:] == target[i:]
}

func isReplacedCharacter(src, target string, i int) bool {
	return i+1 < len(src) && i+1 < len(target) && src[i+1:] == target[i+1:]
}

func isAddedCharacter(src, target string, i int) bool {
	return i < len(src) && i+1 < len(target) && src[i:] == target[i+1:]
}
