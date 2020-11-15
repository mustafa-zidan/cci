package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SumListsSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSumLists(t *testing.T) {
	suite.Run(t, new(SumListsSuite))
}

func (suite *SumListsSuite) TestSumLists() {
	l := (&LinkedList{}).New([]interface{}{7, 1, 6})
	r := (&LinkedList{}).New([]interface{}{5, 9, 2})
	expected := (&LinkedList{}).New([]interface{}{2, 1, 9})
	result := sum(l, r)
	assert.Equal(suite.T(), expected, result)
}

func BenchmarkSumLists(b *testing.B) {
	l := (&LinkedList{}).New([]interface{}{7, 1, 6})
	r := (&LinkedList{}).New([]interface{}{5, 9, 2})
	for i := 0; i < b.N; i++ {
		sum(l, r)
	}
}
