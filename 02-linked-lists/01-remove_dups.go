package linkedlists

// Remove Dups: Write code to remove duplicates from an unsorted linked list.
//
// FOLLOW UP How would you solve this problem if a temporary buffer is not allowed?

// create an intermediate buffer and if the item already exists
// then remove it from the list
func removeDups(l *LinkedList) *LinkedList {
	hash := make([]bool, 128)
	var root, previous *LinkedList = l, nil
	for l != nil {
		if hash[l.Val.(rune)-'a'] {
			previous.Next = l.Next
		} else {
			hash[l.Val.(rune)-'a'] = true
			previous = l
		}
		l = l.Next
	}
	return root
}

// create a bit map where each character is presented by one bit
func removeDupsAlternativeSolution(l *LinkedList) *LinkedList {
	hash := uint64(0)
	var root, previous *LinkedList = l, nil
	for l != nil {
		// character position in the bit map
		c := uint64(1) << uint(l.Val.(rune)-'a')
		// this character exists before
		if hash&c != 0 {
			previous.Next = l.Next
		} else {
			hash |= c // add character to the hash
			previous = l
		}
		l = l.Next
	}
	return root
}
