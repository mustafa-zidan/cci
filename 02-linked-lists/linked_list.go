package linkedlists

type LinkedList struct {
	Val  interface{}
	Next *LinkedList
}

func (l *LinkedList) New(items []interface{}) *LinkedList {
	if len(items) == 0 {
		return l
	}
	l.Val = items[0]
	current := l.Next
	for i := 1; i < len(items); i++ {
		current = &LinkedList{items[i], nil}
		current = current.Next
	}

	return l
}

type DoubleLinkedList struct {
	Previous *DoubleLinkedList
	Val      interface{}
	Next     *DoubleLinkedList
}

func (l *DoubleLinkedList) New(items []interface{}) *DoubleLinkedList {
	if len(items) == 0 {
		return l
	}
	l.Val = items[0]
	previous, current := l, l.Next
	for i := 1; i < len(items); i++ {
		current = &DoubleLinkedList{previous, items[i], nil}
		previous = current
		current = current.Next
	}
	return l
}
