package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type IntersectionSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestIntersection(t *testing.T) {
	suite.Run(t, new(IntersectionSuite))
}

func (suite *IntersectionSuite) TestIntersection() {
	//case 1 nils
	assert.Nil(suite.T(), intersection(nil, nil))
	// case 2 second is shorter than first
	first := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'e'})
	second := (&LinkedList{}).New([]interface{}{'f', 'g', 'h'})
	// link the two lists
	second.Get(2).Next = first.Get(2)
	expected := first.Get(2)

	result := intersection(first, second)
	assert.Equal(suite.T(), expected, result)

	// case 2 second is shorter than first
	first = (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'e', 'f'})
	second = (&LinkedList{}).New([]interface{}{'g', 'h'})
	// link the two lists
	second.Get(1).Next = first.Get(3)
	expected = first.Get(3)

	result = intersection(first, second)
	assert.Equal(suite.T(), expected, result)

	// case 3 same length
	first = (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'e', 'f'})
	second = (&LinkedList{}).New([]interface{}{'g', 'h'})
	// link the two lists
	second.Get(1).Next = first.Get(2)
	expected = first.Get(2)

	result = intersection(first, second)
	assert.Equal(suite.T(), expected, result)

	// case 4 no intersection
	first = (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'e', 'f'})
	second = (&LinkedList{}).New([]interface{}{'g', 'h', 'i', 'j', 'k', 'h'})

	assert.Nil(suite.T(), intersection(first, second))
}

func BenchmarkIntersection(b *testing.B) {
	for i := 0; i < b.N; i++ {
		first := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'b', 'a'})
		second := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'b', 'a'})
		intersection(first, second)
	}
}
