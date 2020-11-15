package linkedlists

//Sum Lists: You have two numbers represented by a linked list,
//where each node contains a single digit. The digits are stored in
//reverse order, such that the 1 's digit is at the head of the list.
//
//Write a function that adds the two numbers and returns the sum as a linked list.
//
//EXAMPLE Input:
//	(7 -> 1 -> 6) + (5 -> 9 -> 2). That is, 617 + 295.
//Output:  2 -> 1 -> 9. That is, 912.
//
//FOLLOW UP Suppose the digits are stored in forward order.
//Repeat the above problem.
//
//EXAMPLE Input:
//	(6 -> 1 -> 7) + (2 -> 9 -> 5). That is, 617 + 295.
//Output: 9 -> 1 -> 2. That is, 912.

func sum(l, r *LinkedList) *LinkedList {
	s := &LinkedList{}
	current, carry := s, 0
	for l != nil || r != nil {
		val := 0
		if l == nil && r != nil {
			val = r.Val.(int) + carry
			r = r.Next
		} else if r == nil && l != nil {
			val = (l.Val.(int) + carry)
			l = l.Next
		} else {
			val = (l.Val.(int) + r.Val.(int) + carry)
			l, r = l.Next, r.Next
		}
		carry = val / 10
		current.Next = &LinkedList{val % 10, nil}
		current = current.Next
	}
	return s.Next
}
