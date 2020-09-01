package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StringCompressionTestingSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestStringCompression(t *testing.T) {
	suite.Run(t, new(StringCompressionTestingSuite))
}

func (suite *StringCompressionTestingSuite) TestStringCompression() {
	assert.Equal(suite.T(), "a2b1c5a3", compress("aabcccccaaa"))

}

func BenchmarkStringCompression(b *testing.B) {
	for i := 0; i < b.N; i++ {
		compress("aabcccccaaa")
	}
}
