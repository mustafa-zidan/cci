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

func (suite *AddWithoutPlusTestingSuite) TestAddWithoutPlus() {
	assert.Equal(suite.T(), 9, add(4, 5))
	assert.Equal(suite.T(), 8, add(4, 4))
	assert.Equal(suite.T(), 143, add(133, 10))
}

func BenchmarkAddWithoutPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(4, 5)
	}
}
