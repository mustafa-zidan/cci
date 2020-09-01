package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RotateMatrixTestingSuite struct {
	suite.Suite
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestRotateMatrix(t *testing.T) {
	suite.Run(t, new(RotateMatrixTestingSuite))
}

func (suite *RotateMatrixTestingSuite) TestRotateMatrix() {
	input := [][]int{[]int{0, 0, 1}, []int{0, 0, 1}, []int{0, 0, 1}}
	expected := [][]int{[]int{0, 0, 0}, []int{0, 0, 0}, []int{1, 1, 1}}
	assert.Equal(suite.T(), expected, rotateMatrix(input))
}

func BenchmarkRotateMatrix(b *testing.B) {
	input := [][]int{[]int{0, 0, 1}, []int{0, 0, 1}, []int{0, 0, 1}}
	for i := 0; i < b.N; i++ {
		rotateMatrix(input)
	}
}
