package hard

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AddWithoutPlusTestingSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAddWithoutPlus(t *testing.T) {
	suite.Run(t, new(AddWithoutPlusTestingSuite))
}

func (suite *AddWithoutPlusTestingSuite) TestAddWithoutPlusFaulty() {
	assert.Equal(suite.T(), 9, addWithoutPlusFaulty(4, 5))
	assert.Equal(suite.T(), 8, addWithoutPlusFaulty(4, 4))
	assert.Equal(suite.T(), 143, addWithoutPlusFaulty(133, 10))
	// this is where the algorithm breaks
	assert.Equal(suite.T(), 117, addWithoutPlusFaulty(123, 10))
}

func (suite *AddWithoutPlusTestingSuite) TestAddWithoutPlus() {
	assert.Equal(suite.T(), 9, addWithoutPlus(4, 5))
	assert.Equal(suite.T(), 8, addWithoutPlus(4, 4))
	assert.Equal(suite.T(), 143, addWithoutPlus(133, 10))
	assert.Equal(suite.T(), 133, addWithoutPlus(123, 10))
}

func BenchmarkAddWithoutPlusFaulty(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addWithoutPlusFaulty(133, 10)
	}
}
func BenchmarkAddWithoutPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addWithoutPlus(133, 10)
	}
}
