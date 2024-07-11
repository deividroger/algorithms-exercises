package main

func TransposeMatrix(matrix [][]int) [][]int {

	result := make([][]int, len(matrix[0]))

	for col := 0; col < len(matrix[0]); col++ {
		result[col] = make([]int, len(matrix))

		for row := 0; row < len(matrix); row++ {
			result[col][row] = matrix[row][col]
		}
	}

	return result
}
