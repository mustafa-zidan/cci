package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PalindromePermutationTestingSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPalindromePermutation(t *testing.T) {
	suite.Run(t, new(PalindromePermutationTestingSuite))
}

func (suite *PalindromePermutationTestingSuite) TestPalindromePermutation() {
	assert.True(suite.T(), palindromePermutation("Tact aCo"))
	assert.True(suite.T(), palindromePermutation("Tact aCo!"))
	assert.True(suite.T(), palindromePermutation("atco cta"))
	assert.True(suite.T(), palindromePermutation("ZYX ZYX!?@"))
	assert.False(suite.T(), palindromePermutation("abcd"))
}

func (suite *PalindromePermutationTestingSuite) TestPalindromePermutationAlternativeImplementation() {
	assert.True(suite.T(), palindromePermutationAlternativeImplementation("Tact aCo"))
	assert.True(suite.T(), palindromePermutationAlternativeImplementation("Tact aCo"))
	assert.True(suite.T(), palindromePermutationAlternativeImplementation("Tact aCo!"))
	assert.True(suite.T(), palindromePermutationAlternativeImplementation("atco cta"))
	assert.True(suite.T(), palindromePermutationAlternativeImplementation("ZYX ZYX!?@"))
	assert.False(suite.T(), palindromePermutationAlternativeImplementation("abcd"))
}

func BenchmarkPalindromePermutation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		palindromePermutation("Tact aCo")
	}
}

func BenchmarkPalindromePermutationAlternativeImplementation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		palindromePermutationAlternativeImplementation("Tact aCo")
	}
}
