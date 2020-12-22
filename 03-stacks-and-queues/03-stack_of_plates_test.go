package stacksandqueues

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SetOfStacksSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestSetOfStacks(t *testing.T) {
	suite.Run(t, new(SetOfStacksSuite))
}

func (suite *SetOfStacksSuite) TestSetOfStacks() {
	sos := NewSetOfStacks(StackCapacity(2))

	assert.Equal(suite.T(), len(sos.stacks), 1)
	assert.Equal(suite.T(), sos.current.cap, StackCapacity(2))
	assert.Equal(suite.T(), sos.current.index, -1)
	assert.Equal(suite.T(), cap(sos.current.items), 2)
	// creating stacks if we hit the single stack cap
	sos.Push(1)
	sos.Push(2)
	sos.Push(3)

	assert.Equal(suite.T(), len(sos.stacks), 2)
	assert.Equal(suite.T(), sos.current.cap, StackCapacity(2))
	assert.Equal(suite.T(), sos.current.index, 0)
	assert.Equal(suite.T(), cap(sos.current.items), 2)

	// test the pop
	assert.Equal(suite.T(), sos.Pop(), 3)

	// test pop works acorss stacks
	assert.Equal(suite.T(), sos.Pop(), 2)
	assert.Equal(suite.T(), len(sos.stacks), 1)
	assert.Equal(suite.T(), sos.current.index, 0)
	assert.Equal(suite.T(), cap(sos.current.items), 2)

	// test popAt
	for i := 4; i < 10; i++ {
		sos.Push(i)
	}
	assert.Equal(suite.T(), sos.PopAt(3), 6)

}
