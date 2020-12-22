package stacksandqueues

import (
	"errors"
	"sync"
)

// Imagine a (literal) stack of plates. If the stack gets too high,
// it might topple. Therefore, in real life, we would likely start
// a new stack when the previous stack exceeds some threshold.
// Implement a data structure SetOfStacks that mimics this.
// SetOfStacks should be composed of several stacks and should create
// a new stack once the previous one exceeds capacity.
// SetOfStacks.push() and SetOfStacks.pop() should behave identically
// to a single stack (that is, pop( ) should return the same values as
// it would if there were just a single stack).
//
// FOLLOW UP
// Implement a function popAt(int index)
// which performs a pop operation on a specificsub-stack.

type SetOfStacks struct {
	maxCap  StackCapacity
	current *StackWithCap
	stacks  []*StackWithCap
	lock    sync.RWMutex
}

type StackWithCap struct {
	cap   StackCapacity
	index int
	items []interface{}
	lock  sync.RWMutex
}

type StackCapacity int

func NewStackWithCap(cap StackCapacity) *StackWithCap {
	stack := &StackWithCap{}
	stack.cap = cap
	stack.index = -1
	stack.items = make([]interface{}, int(cap))
	return stack
}

func (swc *StackWithCap) Push(item interface{}) error {
	swc.lock.Lock()
	defer swc.lock.Unlock()
	if int(swc.cap)-1 == swc.index {
		return errors.New("max cap reached")
	}
	swc.index++
	swc.items[swc.index] = item
	return nil
}

func (swc *StackWithCap) Pop() interface{} {
	swc.lock.Lock()
	defer swc.lock.Unlock()
	if swc.index == -1 {
		return nil
	}
	item := swc.items[swc.index]
	swc.items[swc.index] = nil
	swc.index--
	return item
}

func (swc *StackWithCap) PopAt(index int) interface{} {
	// check if internal index pointer is larger
	// than the index else return nil pop and remove
	// item from internal array
	swc.lock.Lock()
	defer swc.lock.Unlock()
	item := swc.items[index]
	swc.items = append(swc.items[:index], append(swc.items[index+1:], nil)...)
	swc.index--
	return item
}

func (swc *StackWithCap) isEmpty() bool {
	return swc.index <= -1
}

func NewSetOfStacks(cap StackCapacity) *SetOfStacks {
	sos := &SetOfStacks{}
	sos.maxCap = cap
	sos.current = NewStackWithCap(cap)
	sos.stacks = []*StackWithCap{sos.current}
	return sos
}

func (sos *SetOfStacks) Push(item interface{}) {
	err := sos.current.Push(item)
	if err != nil {
		sos.current = NewStackWithCap(sos.maxCap)
		sos.stacks = append(sos.stacks, sos.current)
		sos.current.Push(item)
	}
}

func (sos *SetOfStacks) Pop() interface{} {
	// Pop from the latest stack
	// if the current is empty and we still have
	// remove the current stack from the list and
	// set current to the latest from the list
	if sos.current.isEmpty() && len(sos.stacks) > 1 {
		sos.stacks = sos.stacks[:len(sos.stacks)-1]
		sos.current = sos.stacks[len(sos.stacks)-1]
	}
	item := sos.current.Pop()
	return item
}

func (sos *SetOfStacks) PopAt(index int) interface{} {
	// identify which stack to pop from
	stackIndex := index / int(sos.maxCap)
	// use stack internal popAt
	item := sos.stacks[stackIndex].PopAt(index % int(sos.maxCap))
	// we need to shift all the preceeding stacks by one
	for s := stackIndex; s < len(sos.stacks)-1; s++ {
		sos.stacks[s].items[sos.maxCap-1] = sos.stacks[s+1].items[0]
		sos.stacks[s+1].items = append(sos.stacks[s+1].items[1:], nil)
	}
	sos.current.index--
	return item
}
