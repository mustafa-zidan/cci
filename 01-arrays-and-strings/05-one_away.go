package arrays

// One Away: there are three types of edits that can be preformed on strings,
// insert a character, remove a character or replace a character. Given two
// strings, write a function to check if they are one edit (or zero edits) away
// example:
// pale,  ple  -> true
// pales, pale -> true
// pale,  bale -> true
// pale,  bake -> true

func oneAway(src, target string) bool {
	if len(src) > len(target) {
		// swap if necessary now we are sure src is always smaller than target
		src, target = target, src
	}
	if len(target)-len(src) > 1 {
		return false
	} // perform sanity check before we start
	for i := 0; i < len(src) && i < len(target); i++ {
		if src[i] != target[i] {
			return i+1 < len(src) && (isRemovedCharacter(src, target, i) || isReplacedCharacter(src, target, i))
		}
	}
	return true
}

func isRemovedCharacter(src, target string, i int) bool {
	return src[i:] == target[i+1:]
}

func isReplacedCharacter(src, target string, i int) bool {
	return src[i+1:] == target[i+1:]
}
