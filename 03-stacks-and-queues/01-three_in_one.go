package stacksandqueues

import (
	"sync"
)

// Describe how you could use a single array to implement three stacks.

// This is an extention to the solution of the initial problem that allow any
// number of stacks to be implemented in a single array
//
// MultiPartationStack is a struct that allow creating multiple stacks
// by setting pointer to where the current position in the partation and a cap
// for each partation this allow us to douple the size of the partation each
// time we hit the partation capacity so we don't need to do the allocation
// with each push to the stack
type MultiPartitionStack struct {
	partations []int
	cap        []int
	lock       sync.RWMutex
	items      []interface{}
}

type StackPartation int

func NewMultiPartitionStack(partations int) *MultiPartitionStack {
	mps := MultiPartitionStack{}
	mps.cap = make([]int, partations)
	// set the right index for each partation
	mps.partations = make([]int, partations)
	for i, _ := range mps.cap {
		mps.cap[i] = (i * 2) + 1
		mps.partations[i] = i * 2
	}
	mps.items = make([]interface{}, partations*2)
	return &mps
}

// Push is appending an item to the stack partation
func (mps *MultiPartitionStack) Push(item interface{}, partation StackPartation) {
	mps.lock.Lock()
	defer mps.lock.Unlock()
	partationIndex := int(partation - 1)
	cap := mps.cap[partationIndex]
	partationPointer := mps.partations[partationIndex]
	// if we hit the capacity of the partation we douple the size of the partation
	if cap == partationPointer {
		// calculate the current partation capacity
		previousCap := 0
		if partationIndex > 0 {
			previousCap = mps.cap[partationIndex-1]
		}
		// allocat a new array that has a length equal to currentcap - previous cap index
		allocated := make([]interface{}, (cap-previousCap)*2)
		// add the allocated array to the partataion
		mps.items = append(mps.items[:cap+1], append(allocated, mps.items[cap+1:]...)...)
		// adjust the cap aand partation following the resized partation
		for i := partationIndex; i < len(mps.partations); i++ {
			mps.partations[i] += len(allocated)
			mps.cap[i] += len(allocated)
		}
	}
	// add the item to the index and increment the partation pointer by one
	mps.items[partationPointer] = item
	mps.partations[partationIndex] = partationPointer + 1
}

// Pop is allowing us to pop the latest item in the stack
func (mps *MultiPartitionStack) Pop(partation StackPartation) interface{} {
	mps.lock.Lock()
	defer mps.lock.Unlock()
	partationIndex := int(partation - 1)
	partationPointer := mps.partations[partationIndex]
	previousCap := 0
	// check if we are at the begining of the partation then return nil
	// partation pointer == previous cap then this means we hit
	if partation > 0 {
		previousCap = mps.cap[partationIndex-1]
	}

	if partationPointer == previousCap {
		return nil
	}

	// the last element of the current partation
	item := mps.items[partationPointer]
	mps.partations[partationIndex]--
	return item
}
