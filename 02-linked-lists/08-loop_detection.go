package linkedlists

// Loop Detection: Given a circular linked list, implement an algorithm that
// returns the node at the beginning of the loop.

// DEFINITION Circular linked list: A (corrupt) linked list in which a node's
// next pointer points to an earlier node, so as to make a loop in the
// linked list.
// EXAMPLE Input: A -> B -> C -> D -> E -> C [the same C as earlier]
// Output: C

func findLoopAlternative(l *LinkedList) *LinkedList {
	if l == nil {
		return nil
	}

	m := make(map[*LinkedList]bool, 0)
	for l != nil {
		if _, ok := m[l]; ok {
			return l
		}
		m[l] = true
		l = l.Next
	}
	return nil
}

func findLoop(l *LinkedList) *LinkedList {
	if l == nil {
		return nil
	}
	slow, fast := l, l
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if fast == slow {
			break
		}
	}

	// error check - no meeting point
	if fast == nil || fast.Next == nil {
		return nil
	}
	// move slow to start and check the end of the loop
	slow = l
	for slow != fast {
		slow, fast = slow.Next, fast.Next

	}

	return fast
}
