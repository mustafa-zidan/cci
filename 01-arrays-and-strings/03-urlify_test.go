package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UrlifyTestingSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestUrlify(t *testing.T) {
	suite.Run(t, new(UrlifyTestingSuite))
}

func (suite *UrlifyTestingSuite) TestUrlify() {
	assert.Equal(suite.T(), "Mr%20John%20Smith", urlify("Mr John Smith    ", 13))
}

func (suite *UrlifyTestingSuite) TestUrlifyAlternativeImplementation() {
	assert.Equal(suite.T(), "Mr%20John%20Smith", urlifyAlternativeImplementation("Mr John Smith    ", 13))
}

func BenchmarkUrlify(b *testing.B) {
	for i := 0; i < b.N; i++ {
		urlify("Mr John Smith    ", 13)
	}
}

func BenchmarkUrlifyAlternativeImplementation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		urlifyAlternativeImplementation("Mr John Smith    ", 13)
	}
}
