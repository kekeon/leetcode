package array

func isToeplitzMatrix(matrix [][]int) bool {
	x := len(matrix)
	y := len(matrix[0])

	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}

	return true
}
