package linkedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type PalindromSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPalindrom(t *testing.T) {
	suite.Run(t, new(PalindromSuite))
}

func (suite *PalindromSuite) TestPalindrome() {
	input := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'b', 'a'})
	result := palindrome(input)
	assert.True(suite.T(), result)

	input = (&LinkedList{}).New([]interface{}{'a', 'b', 'b', 'a'})
	result = palindrome(input)
	assert.True(suite.T(), result)

	input = (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'a'})
	result = palindrome(input)
	assert.False(suite.T(), result)
}

func BenchmarkPalindrome(b *testing.B) {
	input := (&LinkedList{}).New([]interface{}{'a', 'b', 'c', 'b', 'a'})
	for i := 0; i < b.N; i++ {
		palindrome(input)
	}
}
