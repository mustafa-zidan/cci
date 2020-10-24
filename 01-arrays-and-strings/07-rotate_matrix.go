package arrays

//Given an image represented by NxN matrix, where each pixel in the image is 4
//bytes, write a method to rotate the image by 90 degree. can you do it in place?

// O(n^2)
func rotateMatrix(matrix [][]int) [][]int {
	n := len(matrix) - 1
	// create NxN result
	result := make([][]int, n+1)

	for i, _ := range result {
		result[i] = make([]int, n+1)
	}

	for i, row := range matrix {
		for j, pixel := range row {
			result[j][n-i] = pixel
		}
	}
	return result
}
