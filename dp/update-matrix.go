package dp


func updateMatrix(matrix [][]int) [][]int {

	l := len(matrix)
	dp := make([][]int, l)
	n := len(matrix[0])

	for a := 0; a < l; a++ {
		dp[a] = make([]int, n)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < n; j++ {

			if matrix[i][j] == 0 {

				dp[i][j] = 0

			}else if j == 0 && i == 0 {

			}
		}
	}

	return dp

}
