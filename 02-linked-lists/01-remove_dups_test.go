package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RemoveDupsTestingSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestRemoveDups(t *testing.T) {
	suite.Run(t, new(RemoveDupsTestingSuite))
}

func (suite *RemoveDupsTestingSuite) TestRemoveDups() {
	input := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'c', 'a', 'e'})
	expected := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'e'})
	output := removeDups(input)
	assert.Equal(suite.T(), expected, output)
}

func (suite *RemoveDupsTestingSuite) TestRemoveDupsAlternativeSolution() {
	input := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'c', 'a', 'e'})
	expected := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'e'})
	output := removeDupsAlternativeSolution(input)
	assert.Equal(suite.T(), expected, output)
}

func BenchmarkRemoveDups(b *testing.B) {
	input := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'c', 'a', 'e'})
	for i := 0; i < b.N; i++ {
		removeDups(input)
	}
}

func BenchmarkRemoveDupsAlternativeSolution(b *testing.B) {
	input := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'c', 'a', 'e'})
	for i := 0; i < b.N; i++ {
		removeDupsAlternativeSolution(input)
	}
}
