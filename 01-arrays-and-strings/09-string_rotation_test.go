package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StringRotationTestingSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestStringRotation(t *testing.T) {
	suite.Run(t, new(StringRotationTestingSuite))
}

func (suite *StringCompressionTestingSuite) TestStringRotation() {
	assert.True(suite.T(), stringRotation("waterbottle", "etbottlewat"))

}

func BenchmarkStringRotation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringRotation("waterbottle", "etbottlewat")
	}
}