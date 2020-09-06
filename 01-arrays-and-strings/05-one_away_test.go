package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type OneAwayTestingSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestOneAway(t *testing.T) {
	suite.Run(t, new(OneAwayTestingSuite))
}

func (suite *OneAwayTestingSuite) TestOneAway() {
	assert.True(suite.T(), oneAway("pale", "ple"))
	assert.True(suite.T(), oneAway("pale", "pales"))
	assert.True(suite.T(), oneAway("pales", "pale"))
	assert.True(suite.T(), oneAway("pale", "bale"))
	assert.False(suite.T(), oneAway("pale", "bake"))
	assert.False(suite.T(), oneAway("pale", "ba"))
}

func BenchmarkOneAway(b *testing.B) {
	for i := 0; i < b.N; i++ {
		oneAway("pale", "ple")
	}
}
