package linkedlists

// Return Kth to Last: Implement an algorithm to find the kth to last
// element of a singly linked list.

// First idea is that we can build a reversed linked list and once
// we hit the end we move forward in the reversed linked list k steps
// Space complexity is O(n) and time complexity is also O(n)
//
// We can also Build a hash table with index and item this will be
// same complexity in terms of time and space but more effiecent in
// retrieving the item.
//
// Second idea would be keep an array of only the last k values. This
// will turn the space complexity to O(K) but time complexity will be
// O(n+k) - need to revise this value -as we need to shift the array each
// time we discover a non-terminal node.
//
// a better idea is to keep the Kth Node only in Memory.
// and use kth.Next each time we pass the Kth element
// This is O(1) space complixy and O(n) time complexity
func returnKthToLast(l *LinkedList, k int) interface{} {
	var kth *LinkedList = l
	count := 0
	for l != nil {
		if count > k {
			kth = kth.Next
		}
		l = l.Next
		count++
	}

	if count < k {
		return nil
	} else {
		return kth.Val
	}
}
