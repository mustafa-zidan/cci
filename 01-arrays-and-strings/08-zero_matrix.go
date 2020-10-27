package arrays

// Write an algorithm such that if an element in an MxN matrix is 0,
// it's entire row and column are set to 0

// The idea here is to create two arrays of boolean one for each position
// in a row each for the each position in a column although this is not
// either time or space efficient
// time complexity O(M*N) space complexity O(M+N)
func zeroMatrix(matrix [][]int) [][]int {
	if len(matrix) == 0 {
		return matrix
	}
	rows, columns := make([]bool, len(matrix)), make([]bool, len(matrix[0]))
	// build mapping
	for i, row := range matrix {
		for j, column := range row {
			if column == 0 {
				rows[i], columns[j] = true, true
			}
		}
	}

	// zero rows
	for i, row := range rows {
		if row {
			for j, _ := range matrix[i] {
				matrix[i][j] = 0
			}
		}
	}

	// zero columns
	for j, column := range columns {
		if column {
			for i, _ := range matrix {
				matrix[i][j] = 0
			}
		}
	}

	return matrix
}

// TODO use two integers that map these rows and columns to bits to get the space
// comlexity to 1
