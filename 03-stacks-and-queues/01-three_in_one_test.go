package stacksandqueues

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ThreeInOneSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestThreeInOne(t *testing.T) {
	suite.Run(t, new(ThreeInOneSuite))
}

func (suite *ThreeInOneSuite) TestThreeInOne() {
	multiStack := NewMultiPartitionStack(3)
	first, second, third := StackPartation(1), StackPartation(2), StackPartation(3)

	multiStack.Push(1, first)
	multiStack.Push(2, first)

	multiStack.Push(3, second)
	multiStack.Push(4, second)
	multiStack.Push(5, second)

	multiStack.Push(6, third)
	multiStack.Push(7, third)
	multiStack.Push(8, third)
	multiStack.Push(9, third)

	assert.Equal(suite.T(), multiStack.Pop(first), 2)
	assert.Equal(suite.T(), multiStack.Pop(second), 5)
	assert.Equal(suite.T(), multiStack.Pop(third), 9)
	assert.Equal(suite.T(), multiStack.Pop(first), 1)
	assert.Nil(suite.T(), multiStack.Pop(first))

	assert.Equal(suite.T(), multiStack.Pop(second), 4)
	assert.Equal(suite.T(), multiStack.Pop(third), 8)
	assert.Equal(suite.T(), multiStack.Pop(second), 3)
	assert.Nil(suite.T(), multiStack.Pop(second))

	assert.Equal(suite.T(), multiStack.Pop(third), 7)
	assert.Equal(suite.T(), multiStack.Pop(third), 6)
	assert.Nil(suite.T(), multiStack.Pop(third))

	multiStack.Push(10, first)
	multiStack.Push(11, second)
	multiStack.Push(12, third)

	assert.Equal(suite.T(), multiStack.Pop(first), 10)
	assert.Equal(suite.T(), multiStack.Pop(second), 11)
	assert.Equal(suite.T(), multiStack.Pop(third), 12)

	assert.Nil(suite.T(), multiStack.Pop(first))
	assert.Nil(suite.T(), multiStack.Pop(second))
	assert.Nil(suite.T(), multiStack.Pop(third))
}

// TODO test concurrency against this design.
