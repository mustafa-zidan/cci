package linkedlists

import (
	"fmt"
	"reflect"
	"strings"
)

type LinkedList struct {
	Val  interface{}
	Next *LinkedList
}

func (l *LinkedList) New(items []interface{}) *LinkedList {
	if len(items) == 0 {
		return l
	}
	l = &LinkedList{items[0], nil}
	root := l
	for i := 1; i < len(items); i++ {
		l.Next = &LinkedList{items[i], nil}
		l = l.Next
	}
	return root
}

func (l *LinkedList) String() string {
	var builder strings.Builder
	current := l
	for current != nil {
		value := reflect.ValueOf(current.Val)
		switch current.Val.(type) {
		case int:
			builder.WriteString(fmt.Sprint(value.Int()))
		case float64:
			builder.WriteString(fmt.Sprint(value.Float()))
		case string:
			builder.WriteString(value.String())
		case rune:
			builder.WriteRune(current.Val.(rune))
		default:
			builder.WriteString(value.String())
		}

		current = current.Next
		if current != nil {
			builder.WriteString(" -> ")
		}
	}
	return builder.String()
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

	var previous *DoubleLinkedList
	l = &DoubleLinkedList{nil, items[0], nil}
	root := l

	for i := 1; i < len(items); i++ {
		l.Next = &DoubleLinkedList{previous, items[i], nil}
		previous = l
		l = l.Next
	}

	return root
}

func (l *DoubleLinkedList) String() string {
	var builder strings.Builder
	current := l
	for current != nil {
		value := reflect.ValueOf(current.Val)
		switch current.Val.(type) {
		case int:
			builder.WriteString(fmt.Sprint(value.Int()))
		case float64:
			builder.WriteString(fmt.Sprint(value.Float()))
		case string:
			builder.WriteString(value.String())
		case rune:
			builder.WriteRune(current.Val.(rune))
		default:
			builder.WriteString(value.String())
		}

		current = current.Next
		if current != nil {
			builder.WriteString(" <-> ")
		}
	}
	return builder.String()
}
