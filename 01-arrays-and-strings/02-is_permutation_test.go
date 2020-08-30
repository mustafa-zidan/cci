package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type IsPermutationTestingSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestIsPermutation(t *testing.T) {
	suite.Run(t, new(IsPermutationTestingSuite))
}

func (suite *IsPermutationTestingSuite) TestIsPermutation() {
	assert.False(suite.T(), isPermutation("hello", "hello my friend"))
	assert.True(suite.T(), isPermutation("elloh", "hello"))
}

func (suite *IsPermutationTestingSuite) TestIsPermutationAlternativeImplementation() {
	assert.False(suite.T(), isPermutationAlternativeImplementation("hello", "hello my friend"))
	assert.True(suite.T(), isPermutationAlternativeImplementation("elloh", "hello"))
}

func BenchmarkIsPermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPermutation("elloh", "hello")
	}
}

func BenchmarkIsPermutationAlternativeImplementation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPermutationAlternativeImplementation("elloh", "hello")
	}
}
