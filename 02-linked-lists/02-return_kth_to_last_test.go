package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ReturnKthToLastSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestReturnKthToLast(t *testing.T) {
	suite.Run(t, new(ReturnKthToLastSuite))
}

func (suite *ReturnKthToLastSuite) TestReturnKthToLast() {
	input := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'c', 'a', 'e'})
	assert.Equal(suite.T(), 'e', returnKthToLast(input, 0))
	assert.Equal(suite.T(), 'a', returnKthToLast(input, 1))
	assert.Equal(suite.T(), 'c', returnKthToLast(input, 2))
	assert.Equal(suite.T(), 'd', returnKthToLast(input, 3))
	assert.Equal(suite.T(), 'c', returnKthToLast(input, 4))
}

func BenchmarkReturnKthToLast(b *testing.B) {
	input := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'd', 'c', 'a', 'e'})
	for i := 0; i < b.N; i++ {
		returnKthToLast(input, 3)
	}
}
