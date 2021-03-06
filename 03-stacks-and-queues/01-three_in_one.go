package stacksandqueues

import (
	"sync"
)

// Describe how you could use a single array to implement three stacks.

// This is an extention to the solution of the initial problem that allow any
// number of stacks to be implemented in a single array
//
// MultiPartitionStack is a struct that allow creating multiple stacks
// by setting pointer to where the current position in the partition and a cap
// for each partition this allow us to douple the size of the partition each
// time we hit the partition capacity so we don't need to do the allocation
// with each push to the stack
type MultiPartitionStack struct {
	partitions []StackPartition
	cap        []StackPartition
	lock       sync.RWMutex
	items      []interface{}
}

type StackPartition int

func NewMultiPartitionStack(partitions int) *MultiPartitionStack {
	mps := MultiPartitionStack{}
	mps.cap = make([]StackPartition, partitions)
	// set the right index for each partition
	mps.partitions = make([]StackPartition, partitions)
	for i, _ := range mps.cap {
		mps.cap[i] = StackPartition((i * 2) + 1)
		mps.partitions[i] = StackPartition(i*2 - 1)
	}
	mps.items = make([]interface{}, partitions*2)
	return &mps
}

// Push is appending an item to the stack partition
func (mps *MultiPartitionStack) Push(item interface{}, partition StackPartition) {
	mps.lock.Lock()
	defer mps.lock.Unlock()
	// internal implementation is zero based
	partition--
	cap := mps.cap[partition]
	// if we hit the capacity of the partition we douple the size of the partition
	if cap == mps.partitions[partition] {
		// calculate the current partition capacity
		previousCap := StackPartition(0)
		if partition > 0 {
			previousCap = mps.cap[partition-1]
		}
		// allocate a new array that has a length equal to currentcap - previous cap index
		allocationSize := (cap - previousCap)
		allocated := make([]interface{}, int(allocationSize))
		// add the allocated array to the partition
		mps.items = append(mps.items[:cap+1], append(allocated, mps.items[cap+1:]...)...)
		// adjust the cap aand partition following the resized partition
		mps.cap[partition] += allocationSize
		for i := partition + 1; i < StackPartition(len(mps.partitions)); i++ {
			mps.partitions[i] += allocationSize
			mps.cap[i] += allocationSize
		}
	}
	// add the item to the index and increment the partition pointer by one
	mps.partitions[partition]++
	mps.items[mps.partitions[partition]] = item
}

// Pop is allowing us to pop the latest item in the stack
func (mps *MultiPartitionStack) Pop(partition StackPartition) interface{} {
	mps.lock.Lock()
	defer mps.lock.Unlock()
	partition--
	partitionPointer := mps.partitions[partition]
	previousCap := StackPartition(-1)
	// check if we are at the begining of the partition then return nil
	// partition pointer == previous cap then this means we hit
	if partition > 1 {
		previousCap = mps.cap[partition-1]
	}

	if partitionPointer == previousCap || partitionPointer < 0 {
		return nil
	}

	// the last element of the current partition
	item := mps.items[partitionPointer]
	mps.items[partitionPointer] = nil
	mps.partitions[partition]--
	// TODO resize down if the pointer is hitting 1/4 partation cap
	return item
}
