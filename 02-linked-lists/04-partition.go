package linkedlists

//Partition: Write code to partition a linked list around a value x,
//such that all nodes less than x come before all nodes greater than or
//equal to x. If x is contained within the list, the values of x only need
//to be after the elements less than x.
//
//The partition element x can appear anywhere in the "right partition";
// it does not need to appear between the left and right partitions.
//
// EXAMPLE Input:
// 		3 -> 5 -> 8 -> 5 -> 10 -> 2 -> 1 [partition = 5]
// Output:
// 		3 -> 1 -> 2 -> 5 -> 8 -> 5 -> 10

func partition(l *LinkedList, p int) {
	left, right := &LinkedList{}, &LinkedList{}
	rootL, rootR := left, right
	for l != nil {
		if l.Val.(int) >= p { // link item to the right partition
			right.Next = l
			right = right.Next
		} else { // link item to the left partition
			left.Next = l
			left = left.Next
		}
		l = l.Next
	}
	// unlink the last item to avoid circular linked list
	right.Next = nil
	// link left and right partition
	left.Next = rootR.Next
	l = rootL.Next
}
