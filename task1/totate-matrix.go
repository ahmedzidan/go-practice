package main

func main() {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotateMatrix(matrix)
}
func rotateMatrix(matrix [][]int) {
	n := len(matrix) // [[1,2,3],[4,5,6],[7,8,9]]
	// [[7,4,1],[8,5,2],[9,6,3]]
	/*
		[1,2,3]
		[4,5,6]
		[7,8,9]
	*/
	/*
		[1,4,7]
		[2,5,8]
		[3,6,9]
	*/
	for i := 0; i < n; i++ { //[1,2,3]
		for j := i + 1; j < n; j++ { //[4,5,6]
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	// [1,4,7]  left = right, right= left

	for i := 0; i < len(matrix); i++ {
		left, right := 0, n-1

		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
}
