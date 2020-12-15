package linkedlists

// Intersection: Given two (singly) linked lists, determine if the
// two lists intersect.
// Return the inter-secting node.
// Note that the intersection is defined based on reference, not value.
// That is, if the kth node of the first linked list is the exact same node
// (by reference) as the jth node of the second linked list,
// then they are intersecting.
//
func intersection(first, second *LinkedList) *LinkedList {
	if first == nil || second == nil {
		return nil
	}

	// get Tail info
	firstTail, firstSize := getTail(first)
	secondTail, secondSize := getTail(second)

	// this means there is no intersection
	if firstTail != secondTail {
		return nil
	}

	// move the longest linked list forward
	// so both match in length
	if firstSize > secondSize {
		for i := 0; i < firstSize-secondSize; i++ {
			first = first.Next
		}
	} else if secondSize > firstSize {
		for i := 0; i < secondSize-firstSize; i++ {
			second = second.Next
		}
	}

	// find the first node that matches
	for first != second {
		first = first.Next
		second = second.Next
	}

	if first == second {
		return first
	} else {
		return nil
	}
}

func getTail(l *LinkedList) (*LinkedList, int) {
	current, count := l, 1
	for current.Next != nil {
		count++
		current = current.Next
	}
	return current, count
}
