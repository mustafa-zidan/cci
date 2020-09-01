package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ZeroMatrixTestingSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestZeroMatrix(t *testing.T) {
	suite.Run(t, new(ZeroMatrixTestingSuite))
}

func (suite *ZeroMatrixTestingSuite) TestZeroMatrix() {
	input := [][]int{[]int{0, 1, 1}, []int{1, 1, 1}, []int{1, 1, 1}}
	expected := [][]int{[]int{0, 0, 0}, []int{0, 1, 1}, []int{0, 1, 1}}
	assert.Equal(suite.T(), expected, zeroMatrix(input))
}

func BenchmarkZeroMatrix(b *testing.B) {
	input := [][]int{[]int{0, 1, 1}, []int{1, 1, 1}, []int{1, 1, 1}}
	for i := 0; i < b.N; i++ {
		zeroMatrix(input)
	}
}
