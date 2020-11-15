package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PartitionSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPartition(t *testing.T) {
	suite.Run(t, new(PartitionSuite))
}

func (suite *PartitionSuite) TestPartition() {
	input := (&LinkedList{}).New([]interface{}{3, 5, 8, 5, 10, 2, 1})
	expected := (&LinkedList{}).New([]interface{}{3, 2, 1, 5, 8, 5, 10})
	partition(input, 5)
	assert.Equal(suite.T(), expected, input)
}

func BenchmarkPartition(b *testing.B) {
	input := (&LinkedList{}).New([]interface{}{3, 5, 8, 5, 10, 2, 1})
	for i := 0; i < b.N; i++ {
		partition(input, 5)
	}
}
