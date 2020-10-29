package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DeleteMiddleNodeSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestDeleteMiddleNode(t *testing.T) {
	suite.Run(t, new(DeleteMiddleNodeSuite))
}

func (suite *DeleteMiddleNodeSuite) TestDeleteMiddleNode() {
	input := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'e'})
	expected := (&LinkedList{}).New([]interface{}{'a', 'b', 'd', 'e'})
	deleteMiddleNode(input.Get(2))
	assert.Equal(suite.T(), expected, input)
}

func BenchmarkDeleteMiddleNode(b *testing.B) {
	input := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'c', 'a', 'e'})
	for i := 0; i < b.N; i++ {
		deleteMiddleNode(input.Get(2))
	}
}
