package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LoopDetectionSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestLoopDetection(t *testing.T) {
	suite.Run(t, new(LoopDetectionSuite))
}

func (suite *LoopDetectionSuite) TestLoopDetection() {
	//case 1 nil
	assert.Nil(suite.T(), findLoop(nil))
	// case 2 no loops
	input := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'e'})
	assert.Nil(suite.T(), findLoop(input))
	// case 3 loop detected
	input = (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'e'})
	// link the two lists
	input.Get(4).Next = input.Get(2)
	expected := input.Get(2)

	result := findLoop(input)
	assert.Equal(suite.T(), expected, result)
}

func (suite *LoopDetectionSuite) TestLoopDetectionAlternative() {
	//case 1 nil
	assert.Nil(suite.T(), findLoopAlternative(nil))
	// case 2 no loops
	input := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'e'})
	assert.Nil(suite.T(), findLoopAlternative(input))
	// case 3 loop detected
	input = (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'e'})
	// link the two lists
	input.Get(4).Next = input.Get(2)
	expected := input.Get(2)

	result := findLoopAlternative(input)
	assert.Equal(suite.T(), expected, result)
}

func BenchmarkLoopDetection(b *testing.B) {
	input := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'e'})
	for i := 0; i < b.N; i++ {
		findLoop(input)
	}
}

func BenchmarkLoopDetectionAlternative(b *testing.B) {
	input := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'e'})
	for i := 0; i < b.N; i++ {
		findLoopAlternative(input)
	}
}
