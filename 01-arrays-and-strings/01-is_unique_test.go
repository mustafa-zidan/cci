package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type IsUniqueTestingSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestIsUnique(t *testing.T) {
	suite.Run(t, new(IsUniqueTestingSuite))
}

func (suite *IsUniqueTestingSuite) TestIsUnique() {
	assert.True(suite.T(), isUnique(""))
	assert.True(suite.T(), isUnique("a"))
	assert.True(suite.T(), isUnique("abcd"))
	assert.False(suite.T(), isUnique("Hello"))
	assert.False(suite.T(), isUnique("xyzz"))
}

func (suite *IsUniqueTestingSuite) TestIsUniqueAlternativeImplementation() {
	assert.True(suite.T(), isUniqueAlternativeImplementation(""))
	assert.True(suite.T(), isUniqueAlternativeImplementation("a"))
	assert.True(suite.T(), isUniqueAlternativeImplementation("abcd"))
	assert.False(suite.T(), isUniqueAlternativeImplementation("Hello"))
}

func BenchmarkIsUnique(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isUnique("abcd")
	}
}

func BenchmarkIsUniqueAlternativeImplementation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isUniqueAlternativeImplementation("abcd")
	}
}
